//go:build !js

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media/ivfreader"
)

var peerConnection *webrtc.PeerConnection

func doSignaling(res http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func createPeerConnection(res http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func addVideo(res http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

func removeVideo(res http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	var err error
	if peerConnection, err = webrtc.NewPeerConnection(webrtc.Configuration{}); err != nil {
		panic(err)
	}
	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", cErr)
		}
	}()

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

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/createPeerConnection", createPeerConnection)
	http.HandleFunc("/addVideo", addVideo)
	http.HandleFunc("/removeVideo", removeVideo)

	go func() {
		fmt.Println("Open http://localhost:8080 to access this demo")

		panic(http.ListenAndServe(":8080", nil))
	}()

	select {}
}

func writeVideoToTrack(
	ivf *ivfreader.IVFReader, header *ivfreader.IVFFileHeader, track *webrtc.TrackLocalStaticSample,
) {
	_ = "STUB: not implemented"
	return
}
