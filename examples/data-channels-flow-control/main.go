package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pion/webrtc/v4"
)

const (
	bufferedAmountLowThreshold uint64 = 512 * 1024
	maxBufferedAmount          uint64 = 1024 * 1024
)

func check(err error) { _ = "STUB: not implemented"; return }

func setRemoteDescription(pc *webrtc.PeerConnection, sdp []byte) { _ = "STUB: not implemented"; return }

func createOfferer() *webrtc.PeerConnection { _ = "STUB: not implemented"; return nil }

func createAnswerer() *webrtc.PeerConnection { _ = "STUB: not implemented"; return nil }

func main() {
	offerPC := createOfferer()
	defer func() {
		if err := offerPC.Close(); err != nil {
			fmt.Printf("cannot close offerPC: %v\n", err)
		}
	}()

	answerPC := createAnswerer()
	defer func() {
		if err := answerPC.Close(); err != nil {
			fmt.Printf("cannot close answerPC: %v\n", err)
		}
	}()

	answerPC.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			check(offerPC.AddICECandidate(candidate.ToJSON()))
		}
	})

	offerPC.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			check(answerPC.AddICECandidate(candidate.ToJSON()))
		}
	})

	offerPC.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
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

	answerPC.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
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

	offer, err := offerPC.CreateOffer(nil)
	check(err)
	check(offerPC.SetLocalDescription(offer))
	desc, err := json.Marshal(offer)
	check(err)

	setRemoteDescription(answerPC, desc)

	answer, err := answerPC.CreateAnswer(nil)
	check(err)
	check(answerPC.SetLocalDescription(answer))
	desc2, err := json.Marshal(answer)
	check(err)

	setRemoteDescription(offerPC, desc2)

	select {}
}
