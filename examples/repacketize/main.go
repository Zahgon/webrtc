//go:build !js

package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pion/webrtc/v4"
)

//go:embed index.html index.js
var web embed.FS

func main() {
	fs := http.FileServer(http.FS(web))

	http.Handle("/", fs)

	http.HandleFunc("/sdp", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		switch r.Method {
		case "POST":
			pc := createRTCConn()

			sdp := &webrtc.SessionDescription{}
			if err := json.NewDecoder(r.Body).Decode(sdp); err != nil {
				panic(err)
			}

			if err := pc.SetRemoteDescription(*sdp); err != nil {
				panic(err)
			}

			gather := webrtc.GatheringCompletePromise(pc)

			answer, err := pc.CreateAnswer(nil)
			if err != nil {
				panic(err)
			}
			err = pc.SetLocalDescription(answer)
			if err != nil {
				panic(err)
			}

			<-gather

			resp, err := json.Marshal(pc.LocalDescription())
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "application/json")
			if _, err := w.Write(resp); err != nil {
				panic(err)
			}

			return
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Open http://localhost:8080 to access this example")
	panic(http.ListenAndServe(":8080", nil)) //nolint:gosec // example
}

func createRTCConn() *webrtc.PeerConnection {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}
