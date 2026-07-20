package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/pion/webrtc/v4"
)

func signalCandidate(addr string, candidate *webrtc.ICECandidate) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	offerAddr := flag.String("offer-address", "127.0.0.1:50000", "Address that the Offer HTTP server is hosted on.")
	answerAddr := flag.String("answer-address", ":60000", "Address that the Answer HTTP server is hosted on.")
	flag.Parse()

	var candidatesMux sync.Mutex
	pendingCandidates := make([]*webrtc.ICECandidate, 0)

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
		if closeErr := peerConnection.Close(); closeErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", closeErr)
		}
	}()

	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate == nil {
			return
		}

		candidatesMux.Lock()
		defer candidatesMux.Unlock()

		desc := peerConnection.RemoteDescription()
		if desc == nil {
			pendingCandidates = append(pendingCandidates, candidate)
		} else if onICECandidateErr := signalCandidate(*offerAddr, candidate); onICECandidateErr != nil {
			panic(onICECandidateErr)
		}
	})

	http.HandleFunc("/candidate", func(res http.ResponseWriter, req *http.Request) {
		candidate, candidateErr := io.ReadAll(req.Body)
		if candidateErr != nil {
			panic(candidateErr)
		}
		if candidateErr := peerConnection.AddICECandidate(
			webrtc.ICECandidateInit{Candidate: string(candidate)},
		); candidateErr != nil {
			panic(candidateErr)
		}
	})

	http.HandleFunc("/sdp", func(res http.ResponseWriter, req *http.Request) {
		sdp := webrtc.SessionDescription{}
		if decodeErr := json.NewDecoder(req.Body).Decode(&sdp); decodeErr != nil {
			panic(decodeErr)
		}

		if setErr := peerConnection.SetRemoteDescription(sdp); setErr != nil {
			panic(setErr)
		}

		answer, err := peerConnection.CreateAnswer(nil)
		if err != nil {
			panic(err)
		}

		payload, err := json.Marshal(answer)
		if err != nil {
			panic(err)
		}
		resp, err := http.Post(
			fmt.Sprintf("http://%s/sdp", *offerAddr),
			"application/json; charset=utf-8",
			bytes.NewReader(payload),
		)
		if err != nil {
			panic(err)
		} else if closeErr := resp.Body.Close(); closeErr != nil {
			panic(closeErr)
		}

		err = peerConnection.SetLocalDescription(answer)
		if err != nil {
			panic(err)
		}

		candidatesMux.Lock()
		defer candidatesMux.Unlock()

		for _, c := range pendingCandidates {
			if onICECandidateErr := signalCandidate(*offerAddr, c); onICECandidateErr != nil {
				panic(onICECandidateErr)
			}
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

	peerConnection.OnDataChannel(func(dataChannel *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", dataChannel.Label(), dataChannel.ID())
		setupDataChannel(dataChannel)
	})

	go func() { panic(http.ListenAndServe(*answerAddr, nil)) }()

	select {}
}

func setupDataChannel(dataChannel *webrtc.DataChannel) { _ = "STUB: not implemented"; return }
