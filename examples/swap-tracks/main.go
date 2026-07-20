//go:build !js

package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v4"
)

func main() {

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", cErr)
		}
	}()

	outputTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{
		MimeType: webrtc.MimeTypeVP8,
	}, "video", "pion")
	if err != nil {
		panic(err)
	}

	rtpSender, err := peerConnection.AddTrack(outputTrack)
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

	offer := webrtc.SessionDescription{}
	decode(readUntilNewline(), &offer)

	err = peerConnection.SetRemoteDescription(offer)
	if err != nil {
		panic(err)
	}

	currTrack := 0

	trackCount := 0

	packets := make(chan *rtp.Packet, 60)

	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) { //nolint: revive
		fmt.Printf("Track has started, of type %d: %s \n", track.PayloadType(), track.Codec().MimeType)
		trackNum := trackCount
		trackCount++

		var lastTimestamp uint32

		var isCurrTrack bool
		for {

			rtp, _, readErr := track.ReadRTP()
			if readErr != nil {
				panic(readErr)
			}

			oldTimestamp := rtp.Timestamp
			if lastTimestamp == 0 {
				rtp.Timestamp = 0
			} else {
				rtp.Timestamp -= lastTimestamp
			}
			lastTimestamp = oldTimestamp

			if currTrack == trackNum { //nolint:nestif

				if !isCurrTrack {
					isCurrTrack = true
					if track.Kind() == webrtc.RTPCodecTypeVideo {
						if writeErr := peerConnection.WriteRTCP([]rtcp.Packet{
							&rtcp.PictureLossIndication{MediaSSRC: uint32(track.SSRC())},
						}); writeErr != nil {
							fmt.Println(writeErr)
						}
					}
				}
				packets <- rtp
			} else {
				isCurrTrack = false
			}
		}
	})

	ctx, done := context.WithCancel(context.Background())

	peerConnection.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", state.String())

		if state == webrtc.PeerConnectionStateFailed {

			done()
		}

		if state == webrtc.PeerConnectionStateClosed {

			done()
		}
	})

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

	go func() {
		var currTimestamp uint32
		for i := uint16(0); ; i++ {
			packet := <-packets

			currTimestamp += packet.Timestamp
			packet.Timestamp = currTimestamp

			packet.SequenceNumber = i

			if err := outputTrack.WriteRTP(packet); err != nil {
				if errors.Is(err, io.ErrClosedPipe) {

					return
				}

				panic(err)
			}
		}
	}()

	fmt.Printf("Waiting for connection\n")
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if trackCount == 0 {
			continue
		}

		fmt.Printf("Waiting 5 seconds then changing...\n")
		time.Sleep(5 * time.Second)
		if currTrack == trackCount-1 {
			currTrack = 0
		} else {
			currTrack++
		}
		fmt.Printf("Switched to track #%v\n", currTrack+1)
	}
}

func readUntilNewline() (in string) { _ = "STUB: not implemented"; return "" }

func encode(obj *webrtc.SessionDescription) string { _ = "STUB: not implemented"; return "" }

func decode(in string, obj *webrtc.SessionDescription) { _ = "STUB: not implemented"; return }
