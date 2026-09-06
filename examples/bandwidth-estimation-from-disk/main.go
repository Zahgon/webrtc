//go:build !js

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/cc"
	"github.com/pion/interceptor/pkg/gcc"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
	"github.com/pion/webrtc/v4/pkg/media/ivfreader"
)

const (
	lowFile    = "low.ivf"
	lowBitrate = 300_000

	medFile    = "med.ivf"
	medBitrate = 1_000_000

	highFile    = "high.ivf"
	highBitrate = 2_500_000

	ivfHeaderSize = 32
)

func main() { //nolint:gocognit,cyclop,maintidx
	qualityLevels := []struct {
		fileName string
		bitrate  int
	}{
		{lowFile, lowBitrate},
		{medFile, medBitrate},
		{highFile, highBitrate},
	}
	currentQuality := 0

	for _, level := range qualityLevels {
		_, err := os.Stat(level.fileName)
		if os.IsNotExist(err) {
			panic(fmt.Sprintf("File %s was not found", level.fileName))
		}
	}

	interceptorRegistry := &interceptor.Registry{}
	mediaEngine := &webrtc.MediaEngine{}
	if err := mediaEngine.RegisterDefaultCodecs(); err != nil {
		panic(err)
	}

	congestionController, err := cc.NewInterceptor(func() (cc.BandwidthEstimator, error) {
		return gcc.NewSendSideBWE(gcc.SendSideBWEInitialBitrate(lowBitrate))
	})
	if err != nil {
		panic(err)
	}

	estimatorChan := make(chan cc.BandwidthEstimator, 1)
	congestionController.OnNewPeerConnection(func(id string, estimator cc.BandwidthEstimator) { //nolint: revive
		estimatorChan <- estimator
	})

	interceptorRegistry.Add(congestionController)
	if err = webrtc.ConfigureTWCCHeaderExtensionSender(mediaEngine, interceptorRegistry); err != nil {
		panic(err)
	}

	if err = webrtc.RegisterDefaultInterceptors(mediaEngine, interceptorRegistry); err != nil {
		panic(err)
	}

	peerConnection, err := webrtc.NewAPI(
		webrtc.WithInterceptorRegistry(interceptorRegistry), webrtc.WithMediaEngine(mediaEngine),
	).NewPeerConnection(webrtc.Configuration{
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

	estimator := <-estimatorChan

	videoTrack, err := webrtc.NewTrackLocalStaticSample(
		webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8}, "video", "pion",
	)
	if err != nil {
		panic(err)
	}

	rtpSender, err := peerConnection.AddTrack(videoTrack)
	if err != nil {
		panic(err)
	}

	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
	})

	peerConnection.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", state.String())
	})

	offer := webrtc.SessionDescription{}
	decode(readUntilNewline(), &offer)

	if err = peerConnection.SetRemoteDescription(offer); err != nil {
		panic(err)
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	if err = peerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}

	<-gatherComplete

	fmt.Println(encode(peerConnection.LocalDescription()))

	file, err := os.Open(qualityLevels[currentQuality].fileName)
	if err != nil {
		panic(err)
	}

	ivf, header, err := ivfreader.NewWith(file)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(
		time.Millisecond * time.Duration((float32(header.TimebaseNumerator)/float32(header.TimebaseDenominator))*1000),
	)
	defer ticker.Stop()
	frame := []byte{}
	frameHeader := &ivfreader.IVFFrameHeader{}
	currentTimestamp := uint64(0)

	switchQualityLevel := func(newQualityLevel int) {
		fmt.Printf(
			"Switching from %s to %s \n",
			qualityLevels[currentQuality].fileName,
			qualityLevels[newQualityLevel].fileName,
		)

		currentQuality = newQualityLevel
		ivf.ResetReader(setReaderFile(qualityLevels[currentQuality].fileName))
		for {
			if frame, frameHeader, err = ivf.ParseNextFrame(); err != nil {
				break
			} else if frameHeader.Timestamp >= currentTimestamp && frame[0]&0x1 == 0 {
				break
			}
		}
	}

	for ; true; <-ticker.C {
		targetBitrate := estimator.GetTargetBitrate()
		switch {

		case currentQuality != 0 && targetBitrate < qualityLevels[currentQuality].bitrate:
			switchQualityLevel(currentQuality - 1)

		case len(qualityLevels) > (currentQuality+1) && targetBitrate > qualityLevels[currentQuality+1].bitrate:
			switchQualityLevel(currentQuality + 1)

		default:
			frame, frameHeader, err = ivf.ParseNextFrame()
		}

		switch {

		case errors.Is(err, io.EOF):
			ivf.ResetReader(setReaderFile(qualityLevels[currentQuality].fileName))

		case err == nil:
			currentTimestamp = frameHeader.Timestamp
			if err = videoTrack.WriteSample(media.Sample{Data: frame, Duration: time.Second}); err != nil {
				panic(err)
			}

		default:
			panic(err)
		}
	}
}

func setReaderFile(filename string) func(_ int64) io.Reader { _ = "STUB: not implemented"; return nil }

func readUntilNewline() (in string) { _ = "STUB: not implemented"; return "" }

func encode(obj *webrtc.SessionDescription) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *webrtc.SessionDescription) { _ = "STUB: not implemented"; return }
