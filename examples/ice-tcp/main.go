//go:build !js

package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/pion/webrtc/v4"
)

var api *webrtc.API

func doSignaling(res http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

//nolint:cyclop
func main() {
	settingEngine := webrtc.SettingEngine{}

	settingEngine.SetNetworkTypes([]webrtc.NetworkType{
		webrtc.NetworkTypeTCP4,
		webrtc.NetworkTypeTCP6,
	})

	tcpListener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IP{0, 0, 0, 0},
		Port: 8443,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening for ICE TCP at %s\n", tcpListener.Addr())

	tcpMux := webrtc.NewICETCPMux(nil, tcpListener, 8)
	settingEngine.SetICETCPMux(tcpMux)

	api = webrtc.NewAPI(webrtc.WithSettingEngine(settingEngine))

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/doSignaling", doSignaling)

	fmt.Println("Open http://localhost:8080 to access this demo")

	panic(http.ListenAndServe(":8080", nil))
}
