//go:build !js

package webrtc

type Configuration struct {
	ICEServers []ICEServer `json:"iceServers,omitempty"`

	ICETransportPolicy ICETransportPolicy `json:"iceTransportPolicy,omitempty"`

	BundlePolicy BundlePolicy `json:"bundlePolicy,omitempty"`

	RTCPMuxPolicy RTCPMuxPolicy `json:"rtcpMuxPolicy,omitempty"`

	PeerIdentity string `json:"peerIdentity,omitempty"`

	Certificates []Certificate `json:"certificates,omitempty"`

	ICECandidatePoolSize uint8 `json:"iceCandidatePoolSize,omitempty"`

	SDPSemantics SDPSemantics `json:"sdpSemantics,omitempty"`

	AlwaysNegotiateDataChannels bool `json:"alwaysNegotiateDataChannels,omitempty"`
}
