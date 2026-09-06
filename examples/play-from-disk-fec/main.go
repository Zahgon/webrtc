//go:build !js

package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pion/interceptor"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
	"github.com/pion/webrtc/v4/pkg/media/ivfreader"
)

const (
	videoFileName  = "output.ivf"
	answerFileName = "answer.txt"
)

func main() { //nolint:gocognit,cyclop,gocyclo,maintidx

	_, err := os.Stat(videoFileName)

	if os.IsNotExist(err) {
		panic("Could not find `" + videoFileName + "`")
	}

	mediaEngine := &webrtc.MediaEngine{}
	if err = mediaEngine.RegisterDefaultCodecs(); err != nil {
		panic(err)
	}

	interceptorRegistry := &interceptor.Registry{}

	interceptorRegistry.Add(packetDropInterceptorFactory{})

	if err = webrtc.ConfigureFlexFEC03(49, mediaEngine, interceptorRegistry); err != nil {
		panic(err)
	}

	if err = webrtc.RegisterDefaultInterceptors(mediaEngine, interceptorRegistry); err != nil {
		panic(err)
	}

	api := webrtc.NewAPI(
		webrtc.WithMediaEngine(mediaEngine),
		webrtc.WithInterceptorRegistry(interceptorRegistry),
	)

	peerConnection, err := api.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", cErr)
		}
	}()

	iceConnectedCtx, iceConnectedCtxCancel := context.WithCancel(context.Background())

	file, openErr := os.Open(videoFileName)
	if openErr != nil {
		panic(openErr)
	}

	_, header, openErr := ivfreader.NewWith(file)
	if openErr != nil {
		panic(openErr)
	}

	var trackCodec string
	switch header.FourCC {
	case "AV01":
		trackCodec = webrtc.MimeTypeAV1
	case "VP90":
		trackCodec = webrtc.MimeTypeVP9
	case "VP80":
		trackCodec = webrtc.MimeTypeVP8
	default:
		panic(fmt.Sprintf("Unable to handle FourCC %s", header.FourCC))
	}

	videoTrack, videoTrackErr := webrtc.NewTrackLocalStaticSample(
		webrtc.RTPCodecCapability{MimeType: trackCodec}, "video", "pion",
	)
	if videoTrackErr != nil {
		panic(videoTrackErr)
	}

	rtpSender, videoTrackErr := peerConnection.AddTrack(videoTrack)
	if videoTrackErr != nil {
		panic(videoTrackErr)
	}

	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()

	go func() {

		file, ivfErr := os.Open(videoFileName)
		if ivfErr != nil {
			panic(ivfErr)
		}

		ivf, header, ivfErr := ivfreader.NewWith(file)
		if ivfErr != nil {
			panic(ivfErr)
		}

		<-iceConnectedCtx.Done()

		ticker := time.NewTicker(
			time.Millisecond * time.Duration((float32(header.TimebaseNumerator)/float32(header.TimebaseDenominator))*1000),
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			frame, _, ivfErr := ivf.ParseNextFrame()
			if errors.Is(ivfErr, io.EOF) {
				fmt.Printf("All video frames parsed and sent")
				os.Exit(0)
			}

			if ivfErr != nil {
				panic(ivfErr)
			}

			if ivfErr = videoTrack.WriteSample(media.Sample{Data: frame, Duration: time.Second}); ivfErr != nil {
				panic(ivfErr)
			}
		}
	}()

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
		if connectionState == webrtc.ICEConnectionStateConnected {
			iceConnectedCtxCancel()
		}
	})

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

	fmt.Printf("Save the browser's answer to '%s' and press Enter to continue...\n", answerFileName)
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	answerData, readErr := os.ReadFile(answerFileName)
	if readErr != nil {
		panic(readErr)
	}

	answerStr := strings.TrimSpace(string(answerData))
	if len(answerStr) == 0 {
		panic("Answer file is empty")
	}

	answer := webrtc.SessionDescription{}
	decode(answerStr, &answer)

	if err = peerConnection.SetRemoteDescription(answer); err != nil {
		panic(err)
	}

	fmt.Println("Answer received and set successfully!")

	select {}
}

func encode(obj *webrtc.SessionDescription) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *webrtc.SessionDescription) { _ = "STUB: not implemented"; return }

type packetDropInterceptorFactory struct{}

func (f packetDropInterceptorFactory) NewInterceptor(_ string) (interceptor.Interceptor, error) {
	_ = "STUB: not implemented"
	return *new(interceptor.Interceptor), nil
}

type dropFilter struct {
	interceptor.NoOp
	mu                  sync.Mutex
	mediaPacketsTotal   int
	fecPacketsTotal     int
	droppedPacketsTotal int
}

func (i *dropFilter) BindLocalStream(info *interceptor.StreamInfo, writer interceptor.RTPWriter) interceptor.RTPWriter {
	_ = "STUB: not implemented"
	return *new(interceptor.RTPWriter)
}
