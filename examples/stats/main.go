//go:build !js

package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/stats"
	"github.com/pion/webrtc/v4"
)

const statsInterval = time.Second * 5

func main() {

	mediaEngine := &webrtc.MediaEngine{}

	if err := mediaEngine.RegisterDefaultCodecs(); err != nil {
		panic(err)
	}

	interceptorRegistry := &interceptor.Registry{}

	statsInterceptorFactory, err := stats.NewInterceptor()
	if err != nil {
		panic(err)
	}

	var statsGetter stats.Getter
	statsInterceptorFactory.OnNewPeerConnection(func(_ string, g stats.Getter) {
		statsGetter = g
	})
	interceptorRegistry.Add(statsInterceptorFactory)

	if err = webrtc.RegisterDefaultInterceptors(mediaEngine, interceptorRegistry); err != nil {
		panic(err)
	}

	api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine), webrtc.WithInterceptorRegistry(interceptorRegistry))

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

	if _, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio); err != nil {
		panic(err)
	} else if _, err = peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo); err != nil {
		panic(err)
	}

	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) { //nolint: revive
		fmt.Printf("New incoming track with codec: %s\n", track.Codec().MimeType)

		go func() {

			for {
				stats := statsGetter.Get(uint32(track.SSRC()))

				fmt.Printf("Stats for: %s\n", track.Codec().MimeType)
				fmt.Println(stats.InboundRTPStreamStats)

				time.Sleep(statsInterval)
			}
		}()

		rtpBuff := make([]byte, 1500)
		for {
			_, _, readErr := track.Read(rtpBuff)
			if readErr != nil {
				panic(readErr)
			}
		}
	})

	var iceConnectionState atomic.Value
	iceConnectionState.Store(webrtc.ICEConnectionStateNew)

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
		iceConnectionState.Store(connectionState)
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

	for {
		time.Sleep(statsInterval)

		if iceConnectionState.Load() == webrtc.ICEConnectionStateChecking {
			continue
		}

		for _, s := range peerConnection.GetStats() {
			switch stat := s.(type) {
			case webrtc.ICECandidateStats:
				if stat.Type == webrtc.StatsTypeRemoteCandidate {
					fmt.Printf("%s IP(%s) Port(%d)\n", stat.Type, stat.IP, stat.Port)
				}
			default:
			}
		}
	}
}

func readUntilNewline() (in string) { _ = "STUB: not implemented"; return "" }

func encode(obj *webrtc.SessionDescription) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *webrtc.SessionDescription) { _ = "STUB: not implemented"; return }
