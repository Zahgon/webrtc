//go:build !js

package main

import (
	"flag"
	"fmt"

	"github.com/pion/webrtc/v4"
)

func main() {
	isOffer := flag.Bool("offer", false, "Act as the offerer if set")
	port := flag.Int("port", 8080, "http server port")
	flag.Parse()

	iceOptions := webrtc.ICEGatherOptions{
		ICEServers: []webrtc.ICEServer{
			{URLs: []string{"stun:stun.l.google.com:19302"}},
		},
	}

	api := webrtc.NewAPI()

	gatherer, err := api.NewICEGatherer(iceOptions)
	if err != nil {
		panic(err)
	}

	ice := api.NewICETransport(gatherer)

	dtls, err := api.NewDTLSTransport(ice, nil)
	if err != nil {
		panic(err)
	}

	sctp := api.NewSCTPTransport(dtls)

	sctp.OnDataChannel(func(channel *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", channel.Label(), channel.ID())

		channel.OnOpen(handleOnOpen(channel))
		channel.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", channel.Label(), string(msg.Data))
		})
	})

	gatherFinished := make(chan struct{})
	gatherer.OnLocalCandidate(func(candidate *webrtc.ICECandidate) {
		if candidate == nil {
			close(gatherFinished)
		}
	})

	if err = gatherer.Gather(); err != nil {
		panic(err)
	}

	<-gatherFinished

	iceCandidates, err := gatherer.GetLocalCandidates()
	if err != nil {
		panic(err)
	}

	iceParams, err := gatherer.GetLocalParameters()
	if err != nil {
		panic(err)
	}

	dtlsParams, err := dtls.GetLocalParameters()
	if err != nil {
		panic(err)
	}

	sctpCapabilities := sctp.GetCapabilities()

	s := Signal{
		ICECandidates:    iceCandidates,
		ICEParameters:    iceParams,
		DTLSParameters:   dtlsParams,
		SCTPCapabilities: sctpCapabilities,
	}

	iceRole := webrtc.ICERoleControlled

	fmt.Println(encode(s))
	remoteSignal := Signal{}

	if *isOffer {
		signalingChan := httpSDPServer(*port)
		decode(<-signalingChan, &remoteSignal)

		iceRole = webrtc.ICERoleControlling
	} else {
		decode(readUntilNewline(), &remoteSignal)
	}

	if err = ice.SetRemoteCandidates(remoteSignal.ICECandidates); err != nil {
		panic(err)
	}

	err = ice.Start(nil, remoteSignal.ICEParameters, &iceRole)
	if err != nil {
		panic(err)
	}

	if err = dtls.Start(remoteSignal.DTLSParameters); err != nil {
		panic(err)
	}

	if err = sctp.Start(remoteSignal.SCTPCapabilities); err != nil {
		panic(err)
	}

	if *isOffer {
		var id uint16 = 1

		dcParams := &webrtc.DataChannelParameters{
			Label: "Foo",
			ID:    &id,
		}
		var channel *webrtc.DataChannel
		channel, err = api.NewDataChannel(sctp, dcParams)
		if err != nil {
			panic(err)
		}

		go handleOnOpen(channel)()
		channel.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", channel.Label(), string(msg.Data))
		})
	}

	select {}
}

type Signal struct {
	ICECandidates    []webrtc.ICECandidate   `json:"iceCandidates"`
	ICEParameters    webrtc.ICEParameters    `json:"iceParameters"`
	DTLSParameters   webrtc.DTLSParameters   `json:"dtlsParameters"`
	SCTPCapabilities webrtc.SCTPCapabilities `json:"sctpCapabilities"`
}

func handleOnOpen(channel *webrtc.DataChannel) func() { _ = "STUB: not implemented"; return nil }

func readUntilNewline() (in string) { _ = "STUB: not implemented"; return "" }

func encode(obj Signal) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *Signal) { _ = "STUB: not implemented"; return }

func httpSDPServer(port int) chan string { _ = "STUB: not implemented"; return nil }

//nolint: errcheck
