//go:build !js

package main

import (
	"fmt"
	"net/http"

	"github.com/pion/ice/v4"
	"github.com/pion/webrtc/v4"
)

var api *webrtc.API

func doSignaling(res http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func main() {

	settingEngine := webrtc.SettingEngine{}

	mux, err := ice.NewMultiUDPMuxFromPort(8443)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening for WebRTC traffic at %d\n", 8443)
	settingEngine.SetICEUDPMux(mux)

	api = webrtc.NewAPI(webrtc.WithSettingEngine(settingEngine))

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/doSignaling", doSignaling)

	fmt.Println("Open http://localhost:8080 to access this demo")

	panic(http.ListenAndServe(":8080", nil))
}
