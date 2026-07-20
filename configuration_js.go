//go:build js && wasm
// +build js,wasm

package webrtc

type Configuration struct {
	ICEServers []ICEServer

	ICETransportPolicy ICETransportPolicy

	BundlePolicy BundlePolicy

	RTCPMuxPolicy RTCPMuxPolicy

	PeerIdentity string

	ICECandidatePoolSize uint8

	AlwaysNegotiateDataChannels bool

	Certificates []Certificate `json:"certificates,omitempty"`
}
