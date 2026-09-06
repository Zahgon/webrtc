//go:build !js

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync/atomic"
	"time"

	"github.com/pion/logging"
	"github.com/pion/transport/v4/vnet"
	"github.com/pion/webrtc/v4"
)

func main() {
	var inboundBytes int32
	var outboundBytes int32

	wan, err := vnet.NewRouter(&vnet.RouterConfig{
		CIDR:          "1.2.3.0/24",
		LoggerFactory: logging.NewDefaultLoggerFactory(),
	})
	panicIfError(err)

	wan.AddChunkFilter(func(chunk vnet.Chunk) bool {
		netType := chunk.SourceAddr().Network()
		if netType == "udp" {
			dstAddr := chunk.DestinationAddr().String()
			host, _, err2 := net.SplitHostPort(dstAddr)
			panicIfError(err2)
			if host == "1.2.3.4" {

				atomic.AddInt32(&inboundBytes, int32(len(chunk.UserData()))) //nolint:gosec // G115
			}
			srcAddr := chunk.SourceAddr().String()
			host, _, err2 = net.SplitHostPort(srcAddr)
			panicIfError(err2)
			if host == "1.2.3.4" {

				atomic.AddInt32(&outboundBytes, int32(len(chunk.UserData()))) //nolint:gosec // G115
			}
		}

		return true
	})

	go func() {
		duration := 2 * time.Second
		for {
			time.Sleep(duration)

			inBytes := atomic.SwapInt32(&inboundBytes, 0)
			outBytes := atomic.SwapInt32(&outboundBytes, 0)
			inboundThroughput := float64(inBytes) / duration.Seconds()
			outboundThroughput := float64(outBytes) / duration.Seconds()
			log.Printf("inbound throughput : %.01f [Byte/s]\n", inboundThroughput)
			log.Printf("outbound throughput: %.01f [Byte/s]\n", outboundThroughput)
		}
	}()

	offerVNet, err := vnet.NewNet(&vnet.NetConfig{
		StaticIPs: []string{"1.2.3.4"},
	})
	panicIfError(err)

	panicIfError(wan.AddNet(offerVNet))

	offerSettingEngine := webrtc.SettingEngine{}
	offerSettingEngine.SetNet(offerVNet)
	offerAPI := webrtc.NewAPI(webrtc.WithSettingEngine(offerSettingEngine))

	answerVNet, err := vnet.NewNet(&vnet.NetConfig{
		StaticIPs: []string{"1.2.3.5"},
	})
	panicIfError(err)

	panicIfError(wan.AddNet(answerVNet))

	answerSettingEngine := webrtc.SettingEngine{}
	answerSettingEngine.SetNet(answerVNet)
	answerAPI := webrtc.NewAPI(webrtc.WithSettingEngine(answerSettingEngine))

	panicIfError(wan.Start())

	offerPeerConnection, err := offerAPI.NewPeerConnection(webrtc.Configuration{})
	panicIfError(err)
	defer func() {
		if cErr := offerPeerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close offerPeerConnection: %v\n", cErr)
		}
	}()

	answerPeerConnection, err := answerAPI.NewPeerConnection(webrtc.Configuration{})
	panicIfError(err)
	defer func() {
		if cErr := answerPeerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close answerPeerConnection: %v\n", cErr)
		}
	}()

	offerPeerConnection.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s (offerer)\n", state.String())

		if state == webrtc.PeerConnectionStateFailed {

			fmt.Println("Peer Connection has gone to failed exiting")
			os.Exit(0)
		}

		if state == webrtc.PeerConnectionStateClosed {

			fmt.Println("Peer Connection has gone to closed exiting")
			os.Exit(0)
		}
	})

	answerPeerConnection.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s (answerer)\n", state.String())

		if state == webrtc.PeerConnectionStateFailed {

			fmt.Println("Peer Connection has gone to failed exiting")
			os.Exit(0)
		}

		if state == webrtc.PeerConnectionStateClosed {

			fmt.Println("Peer Connection has gone to closed exiting")
			os.Exit(0)
		}
	})

	answerPeerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			panicIfError(offerPeerConnection.AddICECandidate(candidate.ToJSON()))
		}
	})

	offerPeerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			panicIfError(answerPeerConnection.AddICECandidate(candidate.ToJSON()))
		}
	})

	offerDataChannel, err := offerPeerConnection.CreateDataChannel("label", nil)
	panicIfError(err)

	msgSendLoop := func(dc *webrtc.DataChannel, interval time.Duration) {
		for {
			time.Sleep(interval)
			panicIfError(dc.SendText("My DataChannel Message"))
		}
	}

	offerDataChannel.OnOpen(func() {

		msgSendLoop(offerDataChannel, 100*time.Millisecond)
	})

	answerPeerConnection.OnDataChannel(func(answerDataChannel *webrtc.DataChannel) {
		answerDataChannel.OnOpen(func() {

			msgSendLoop(answerDataChannel, 200*time.Millisecond)
		})
	})

	offer, err := offerPeerConnection.CreateOffer(nil)
	panicIfError(err)
	panicIfError(offerPeerConnection.SetLocalDescription(offer))
	panicIfError(answerPeerConnection.SetRemoteDescription(offer))

	answer, err := answerPeerConnection.CreateAnswer(nil)
	panicIfError(err)
	panicIfError(answerPeerConnection.SetLocalDescription(answer))
	panicIfError(offerPeerConnection.SetRemoteDescription(answer))

	select {}
}

func panicIfError(err error) { _ = "STUB: not implemented"; return }
