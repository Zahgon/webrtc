//go:build !js

package main

import (
	"fmt"
	"net/http"

	"github.com/pion/webrtc/v4"
)

func main() {
	var pc *webrtc.PeerConnection

	setupOfferHandler(&pc)
	setupCandidateHandler(&pc)
	setupStaticHandler()

	fmt.Println("google-chrome-unstable --force-fieldtrials=" +
		"WebRTC-Sctp-Snap/Enabled/WebRTC-IceHandshakeDtls/Enabled/ " +
		"--disable-features=WebRtcPqcForDtls http://localhost:8080")
	fmt.Printf("Add `--enable-logging --v=1` and then " +
		"`grep SCTP_PACKET chrome_debug.log | " +
		"text2pcap -D -u 1001,2001 -t \"%%H:%%M:%%S.%%f\" - out.pcap` " +
		"for inspecting the raw packets.\n")
	fmt.Println("🚀 Signaling server started on http://localhost:8080")
	//nolint:gosec
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func setupOfferHandler(pc **webrtc.PeerConnection) { _ = "STUB: not implemented"; return }

func setupICECandidateHandler(pc *webrtc.PeerConnection) { _ = "STUB: not implemented"; return }

func setupDataChannelHandler(pc *webrtc.PeerConnection) { _ = "STUB: not implemented"; return }

func processOffer(
	pc *webrtc.PeerConnection,
	offer webrtc.SessionDescription,
	responseWriter http.ResponseWriter,
) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:err113

func setupCandidateHandler(pc **webrtc.PeerConnection) { _ = "STUB: not implemented"; return }

func setupStaticHandler() { _ = "STUB: not implemented"; return }
