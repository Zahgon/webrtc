//go:build !js

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/ivfreader"
)

const (
	videoFileName = "output.ivf"
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

	mediaEngine := &webrtc.MediaEngine{}
	if err := mediaEngine.RegisterDefaultCodecs(); err != nil {
		panic(err)
	}

	api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))

	gatherer, err := api.NewICEGatherer(iceOptions)
	if err != nil {
		panic(err)
	}

	ice := api.NewICETransport(gatherer)

	dtls, err := api.NewDTLSTransport(ice, nil)
	if err != nil {
		panic(err)
	}

	var (
		rtpReceiver       *webrtc.RTPReceiver
		rtpSendParameters webrtc.RTPSendParameters
	)

	if *isOffer { //nolint:nestif

		file, fileErr := os.Open(videoFileName)
		if fileErr != nil {
			panic(fileErr)
		}

		ivf, header, fileErr := ivfreader.NewWith(file)
		if fileErr != nil {
			panic(fileErr)
		}

		trackLocal := fourCCToTrack(header.FourCC)

		rtpSender, fileErr := api.NewRTPSender(trackLocal, dtls)
		if fileErr != nil {
			panic(fileErr)
		}

		rtpSendParameters = rtpSender.GetParameters()

		if fileErr = rtpSender.Send(rtpSendParameters); fileErr != nil {
			panic(fileErr)
		}

		go writeFileToTrack(ivf, header, trackLocal)
	} else {
		if rtpReceiver, err = api.NewRTPReceiver(webrtc.RTPCodecTypeVideo, dtls); err != nil {
			panic(err)
		}
	}

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

	signal := Signal{
		ICECandidates:     iceCandidates,
		ICEParameters:     iceParams,
		DTLSParameters:    dtlsParams,
		RTPSendParameters: rtpSendParameters,
	}

	iceRole := webrtc.ICERoleControlled

	fmt.Println(encode(&signal))
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

	if err = ice.Start(nil, remoteSignal.ICEParameters, &iceRole); err != nil {
		panic(err)
	}

	if err = dtls.Start(remoteSignal.DTLSParameters); err != nil {
		panic(err)
	}

	if !*isOffer {
		if err = rtpReceiver.Receive(webrtc.RTPReceiveParameters{
			Encodings: []webrtc.RTPDecodingParameters{
				{
					RTPCodingParameters: remoteSignal.RTPSendParameters.Encodings[0].RTPCodingParameters,
				},
			},
		}); err != nil {
			panic(err)
		}

		remoteTrack := rtpReceiver.Track()
		pkt, _, err := remoteTrack.ReadRTP()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Got RTP Packet with SSRC %d \n", pkt.SSRC)
	}

	select {}
}

func fourCCToTrack(fourCC string) *webrtc.TrackLocalStaticSample {
	_ = "STUB: not implemented"
	return nil
}

func writeFileToTrack(ivf *ivfreader.IVFReader, header *ivfreader.IVFFileHeader, track *webrtc.TrackLocalStaticSample) {
	_ = "STUB: not implemented"
	return
}

//nolint: gocritic

type Signal struct {
	ICECandidates     []webrtc.ICECandidate    `json:"iceCandidates"`
	ICEParameters     webrtc.ICEParameters     `json:"iceParameters"`
	DTLSParameters    webrtc.DTLSParameters    `json:"dtlsParameters"`
	RTPSendParameters webrtc.RTPSendParameters `json:"rtpSendParameters"`
}

func readUntilNewline() (in string) { _ = "STUB: not implemented"; return "" }

func encode(obj *Signal) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *Signal) { _ = "STUB: not implemented"; return }

func httpSDPServer(port int) chan string { _ = "STUB: not implemented"; return nil }

//nolint: errcheck
