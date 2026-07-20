package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pion/webrtc/v4"
)

const messageSize = 15

func main() {
	sdpChan := httpSDPServer(8080)

	s := webrtc.SettingEngine{}
	s.DetachDataChannels()

	api := webrtc.NewAPI(webrtc.WithSettingEngine(s))

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	peerConnection, err := api.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", cErr)
		}
	}()

	peerConnection.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", state.String())

		if state == webrtc.PeerConnectionStateFailed {

			fmt.Println("Peer Connection has gone to failed exiting")
			os.Exit(0)
		}

		if state == webrtc.PeerConnectionStateClosed {

			fmt.Println("Peer Connection has gone to closed exiting")
			os.Exit(0)
		}
	})

	dataChannel, err := peerConnection.CreateDataChannel("", nil)
	if err != nil {
		panic(err)
	}

	dataChannel.OnOpen(func() {
		fmt.Printf("Data channel '%s'-'%d' open.\n", dataChannel.Label(), dataChannel.ID())

		raw, dErr := dataChannel.Detach()
		if dErr != nil {
			panic(dErr)
		}

		go ReadLoop(raw)

		go WriteLoop(raw)
	})

	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	if err = peerConnection.SetLocalDescription(offer); err != nil {
		panic(err)
	}

	<-gatherComplete

	fmt.Println(encode(peerConnection.LocalDescription()))

	answer := webrtc.SessionDescription{}
	decode(<-sdpChan, &answer)

	err = peerConnection.SetRemoteDescription(answer)
	if err != nil {
		panic(err)
	}

	select {}
}

func ReadLoop(d io.Reader) { _ = "STUB: not implemented"; return }

func WriteLoop(d io.Writer) { _ = "STUB: not implemented"; return }

func httpSDPServer(port int) chan string { _ = "STUB: not implemented"; return nil }

//nolint: errcheck

func encode(obj *webrtc.SessionDescription) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *webrtc.SessionDescription) { _ = "STUB: not implemented"; return }
