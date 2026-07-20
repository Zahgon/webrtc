package main

import (
	"fmt"
	"net/http"

	"github.com/pion/webrtc/v4"
)

var peerConnection *webrtc.PeerConnection

func doSignaling(res http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/doSignaling", doSignaling)

	fmt.Println("Open http://localhost:8080 to access this demo")

	panic(http.ListenAndServe(":8080", nil))
}
