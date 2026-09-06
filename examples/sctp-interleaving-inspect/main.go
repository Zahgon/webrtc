//go:build !js

package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"

	"github.com/pion/datachannel"
	"github.com/pion/ice/v4"
	"github.com/pion/sdp/v3"
)

const (
	listenAddr     = ":8080"
	sctpPort       = 5000
	maxMessageSize = 64 * 1024 * 1024
	outboundMTU    = 1191
	readBufferSize = 20 * 1024 * 1024
)

var (
	errPeerProvidedNoDTLSCertificate = errors.New("peer provided no DTLS certificate")
	errNoApplicationMediaSection     = errors.New("offer has no application media section")
	errMissingICECredentials         = errors.New("offer missing ICE credentials")
	errNoICECandidates               = errors.New("offer has no ICE candidates; wait for ICE gathering to complete")
	errMissingDTLSFingerprint        = errors.New("offer missing DTLS fingerprint")
	errUnsupportedFingerprint        = errors.New("unsupported fingerprint algorithm")
	errRemoteFingerprintMismatch     = errors.New("remote certificate fingerprint mismatch")
	errDataChannelsNotReady          = errors.New("control and bulk channels are not ready")
	errTrafficShaperNotReady         = errors.New("traffic shaper is not ready")
)

func main() {
	events := newEventHub()

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		http.ServeFile(responseWriter, request, "./index.html")
	})
	http.HandleFunc("/events", events.serveHTTP)
	http.HandleFunc("/status", serveStatus)
	http.HandleFunc("/server-pressure", func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(responseWriter, "method not allowed", http.StatusMethodNotAllowed)

			return
		}

		sessionID := request.URL.Query().Get("session")
		runID := request.URL.Query().Get("run")
		if err := events.startServerPressure(sessionID, runID); err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)

			return
		}

		responseWriter.WriteHeader(http.StatusNoContent)
	})
	http.HandleFunc("/offer", func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(responseWriter, "method not allowed", http.StatusMethodNotAllowed)

			return
		}

		offer, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)

			return
		}

		sessionID := request.URL.Query().Get("session")
		answer, err := acceptBrowserOffer(request.Context(), sessionID, string(offer), events)
		if err != nil {
			events.publish(sessionID, logEvent("error", err.Error()))
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)

			return
		}

		responseWriter.Header().Set("Content-Type", "application/sdp")

		//nolint:gosec
		if _, err = io.WriteString(responseWriter, answer); err != nil {
			events.publish(sessionID, logEvent("error", err.Error()))
		}
	})

	fmt.Printf("Listening on %s\n", listenAddr)
	//nolint:gosec
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		panic(err)
	}
}

func acceptBrowserOffer(ctx context.Context, sessionID string, rawOffer string, events *eventHub) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func gatherICECandidates(ctx context.Context, sessionID string, events *eventHub) (*ice.Agent, []ice.Candidate, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func runStack(
	ctx context.Context,
	agent *ice.Agent,
	offer browserOffer,
	cert tls.Certificate,
	sessionID string,
	events *eventHub,
) error {
	_ = "STUB: not implemented"
	return nil
}

func echoDataChannel(dc *datachannel.DataChannel, sessionID string, events *eventHub) {
	_ = "STUB: not implemented"
	return
}

type browserOffer struct {
	raw         *sdp.SessionDescription
	media       *sdp.MediaDescription
	iceUfrag    string
	icePwd      string
	fingerprint fingerprint
	candidates  []ice.Candidate
}

type fingerprint struct {
	algorithm string
	value     string
}

func parseOffer(rawOffer string) (browserOffer, error) {
	_ = "STUB: not implemented"
	return *new(browserOffer), nil
}

func attr(desc *sdp.SessionDescription, media *sdp.MediaDescription, key string) string {
	_ = "STUB: not implemented"
	return ""
}

func parseFingerprint(raw string) (fingerprint, error) {
	_ = "STUB: not implemented"
	return *new(fingerprint), nil
}

func parseCandidates(media *sdp.MediaDescription) ([]ice.Candidate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildAnswer(
	localUfrag string,
	localPwd string,
	candidates []ice.Candidate,
	cert tls.Certificate,
	offer browserOffer,
) string {
	_ = "STUB: not implemented"
	return ""
}

func certificateFingerprint(cert tls.Certificate) string { _ = "STUB: not implemented"; return "" }

func verifyFingerprint(rawCert []byte, expected fingerprint) error {
	_ = "STUB: not implemented"
	return nil
}

type connectedPacketConn struct {
	net.Conn
	remote net.Addr
}

func (c *connectedPacketConn) ReadFrom(p []byte) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func (c *connectedPacketConn) WriteTo(p []byte, _ net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *connectedPacketConn) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

type sctpTapConn struct {
	net.Conn
	inspect func(direction string, packet []byte)
	shaper  *trafficShaper
}

func (c *sctpTapConn) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *sctpTapConn) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type eventHub struct {
	mu          sync.Mutex
	clients     map[string]map[chan []byte]struct{}
	shapers     map[string]*trafficShaper
	connections map[string]map[string]*datachannel.DataChannel
}

func newEventHub() *eventHub { _ = "STUB: not implemented"; return nil }

func (h *eventHub) serveHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *eventHub) publish(sessionID string, event any) { _ = "STUB: not implemented"; return }

func (h *eventHub) registerShaper(sessionID string) *trafficShaper {
	_ = "STUB: not implemented"
	return nil
}

func (h *eventHub) unregisterShaper(sessionID string, shaper *trafficShaper) {
	_ = "STUB: not implemented"
	return
}

func (h *eventHub) registerDataChannel(sessionID string, dc *datachannel.DataChannel) {
	_ = "STUB: not implemented"
	return
}

func (h *eventHub) unregisterDataChannel(sessionID string, dc *datachannel.DataChannel) {
	_ = "STUB: not implemented"
	return
}

func (h *eventHub) startServerPressure(sessionID string, runID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *eventHub) runServerPressure(
	sessionID string,
	runID string,
	shaper *trafficShaper,
	control *datachannel.DataChannel,
	bulk *datachannel.DataChannel,
) {
	_ = "STUB: not implemented"
	return
}

type trafficShaper struct {
	mu             sync.Mutex
	enabled        bool
	bytesPerSecond int
}

func (s *trafficShaper) setEnabled(enabled bool) { _ = "STUB: not implemented"; return }

func (s *trafficShaper) throttle(n int) { _ = "STUB: not implemented"; return }

func logEvent(kind, message string) map[string]any { _ = "STUB: not implemented"; return nil }

func packetEvent(direction string, raw []byte) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

type sctpChunkView struct {
	Type                   uint8  `json:"type"`
	Name                   string `json:"name"`
	Length                 uint16 `json:"length"`
	InterleavingAdvertised bool   `json:"interleavingAdvertised,omitempty"`
}

func parseSCTPChunks(raw []byte) []sctpChunkView { _ = "STUB: not implemented"; return nil }

//nolint:gosec // chunkLength is uint16 on the wire.

func chunkAdvertisesInterleaving(chunk []byte) bool { _ = "STUB: not implemented"; return false }

func chunkTypeName(chunkType uint8) string { _ = "STUB: not implemented"; return "" }

func serveStatus(responseWriter http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}
