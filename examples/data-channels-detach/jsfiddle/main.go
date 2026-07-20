//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"io"
	"syscall/js"

	"github.com/pion/webrtc/v4"
)

const messageSize = 15

func main() {

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
		handleError(err)
	}

	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		handleError(err)
	}

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		log(fmt.Sprintf("ICE Connection State has changed: %s\n", connectionState.String()))
	})

	dataChannel.OnOpen(func() {
		log(fmt.Sprintf("Data channel '%s'-'%d' open.\n", dataChannel.Label(), dataChannel.ID()))

		raw, dErr := dataChannel.Detach()
		if dErr != nil {
			handleError(dErr)
		}

		go ReadLoop(raw)

		go WriteLoop(raw)
	})

	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		handleError(err)
	}

	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		handleError(err)
	}

	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		log(fmt.Sprint(state))
	})
	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			encodedDescr := encode(peerConnection.LocalDescription())
			el := getElementByID("localSessionDescription")
			el.Set("value", encodedDescr)
		}
	})

	js.Global().Set("startSession", js.FuncOf(func(_ js.Value, _ []js.Value) any {
		go func() {
			el := getElementByID("remoteSessionDescription")
			sd := el.Get("value").String()
			if sd == "" {
				js.Global().Call("alert", "Session Description must not be empty")
				return
			}

			descr := webrtc.SessionDescription{}
			decode(sd, &descr)
			if err := peerConnection.SetRemoteDescription(descr); err != nil {
				handleError(err)
			}
		}()
		return js.Undefined()
	}))

	select {}
}

func ReadLoop(d io.Reader) { _ = "STUB: not implemented"; return }

func WriteLoop(d io.Writer) { _ = "STUB: not implemented"; return }

func log(msg string) { _ = "STUB: not implemented"; return }

func handleError(err error) { _ = "STUB: not implemented"; return }

func getElementByID(id string) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func readUntilNewline() (in string) { _ = "STUB: not implemented"; return "" }

func encode(obj *webrtc.SessionDescription) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *webrtc.SessionDescription) { _ = "STUB: not implemented"; return }
