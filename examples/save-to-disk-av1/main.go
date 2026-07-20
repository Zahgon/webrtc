//go:build !js

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/intervalpli"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
	"github.com/pion/webrtc/v4/pkg/media/ivfwriter"
)

func saveToDisk(writer media.Writer, track *webrtc.TrackRemote) { _ = "STUB: not implemented"; return }

func main() {

	mediaEngine := &webrtc.MediaEngine{}

	if err := mediaEngine.RegisterCodec(webrtc.RTPCodecParameters{
		RTPCodecCapability: webrtc.RTPCodecCapability{
			MimeType:     webrtc.MimeTypeAV1,
			ClockRate:    90000,
			Channels:     0,
			SDPFmtpLine:  "",
			RTCPFeedback: nil,
		},
		PayloadType: 96,
	}, webrtc.RTPCodecTypeVideo); err != nil {
		panic(err)
	}

	interceptorRegistry := &interceptor.Registry{}

	intervalPliFactory, err := intervalpli.NewReceiverInterceptor()
	if err != nil {
		panic(err)
	}
	interceptorRegistry.Add(intervalPliFactory)

	if err = webrtc.RegisterDefaultInterceptors(mediaEngine, interceptorRegistry); err != nil {
		panic(err)
	}

	api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine), webrtc.WithInterceptorRegistry(interceptorRegistry))

	config := webrtc.Configuration{}

	peerConnection, err := api.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	if _, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo); err != nil {
		panic(err)
	}

	ivfFile, err := ivfwriter.New("output.ivf", ivfwriter.WithCodec(webrtc.MimeTypeAV1))
	if err != nil {
		panic(err)
	}

	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) { //nolint: revive
		if strings.EqualFold(track.Codec().MimeType, webrtc.MimeTypeAV1) {
			fmt.Println("Got AV1 track, saving to disk as output.ivf")
			saveToDisk(ivfFile, track)
		}
	})

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())

		if connectionState == webrtc.ICEConnectionStateConnected {
			fmt.Println("Ctrl+C the remote client to stop the demo")
		} else if connectionState == webrtc.ICEConnectionStateFailed || connectionState == webrtc.ICEConnectionStateClosed {
			if closeErr := ivfFile.Close(); closeErr != nil {
				panic(closeErr)
			}

			fmt.Println("Done writing media files")

			if closeErr := peerConnection.Close(); closeErr != nil {
				panic(closeErr)
			}

			os.Exit(0)
		}
	})

	offer := webrtc.SessionDescription{}
	decode(readUntilNewline(), &offer)

	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		panic(err)
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	<-gatherComplete

	fmt.Println(encode(peerConnection.LocalDescription()))

	select {}
}

func readUntilNewline() (in string) { _ = "STUB: not implemented"; return "" }

func encode(obj *webrtc.SessionDescription) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *webrtc.SessionDescription) { _ = "STUB: not implemented"; return }
