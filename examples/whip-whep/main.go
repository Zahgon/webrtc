//go:build !js

package main

import (
	"fmt"
	"net/http"

	"github.com/pion/webrtc/v4"
)

var (
	videoTrack *webrtc.TrackLocalStaticRTP
	audioTrack *webrtc.TrackLocalStaticRTP

	peerConnectionConfiguration = webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
)

func main() {

	var err error
	if videoTrack, err = webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{
		MimeType: webrtc.MimeTypeH264,
	}, "video", "pion"); err != nil {
		panic(err)
	}
	if audioTrack, err = webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{
		MimeType: webrtc.MimeTypeOpus,
	}, "audio", "pion"); err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/whep", whepHandler)
	http.HandleFunc("/whip", whipHandler)

	fmt.Println("Open http://localhost:8080 to access this demo")
	panic(http.ListenAndServe(":8080", nil))
}

func whipHandler(res http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func whepHandler(res http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

func writeAnswer(res http.ResponseWriter, peerConnection *webrtc.PeerConnection, offer []byte, path string) {
	_ = "STUB: not implemented"
	return
}

//nolint: errcheck
