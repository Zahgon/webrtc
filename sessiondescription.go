package webrtc

import (
	"github.com/pion/sdp/v3"
)

type ICETrickleCapability int

const (
	ICETrickleCapabilityUnknown ICETrickleCapability = iota

	ICETrickleCapabilitySupported

	ICETrickleCapabilityUnsupported
)

func (t ICETrickleCapability) String() string { _ = "STUB: not implemented"; return "" }

type SessionDescription struct {
	Type SDPType `json:"type"`
	SDP  string  `json:"sdp"`

	parsed *sdp.SessionDescription
}

func (sd *SessionDescription) Unmarshal() (*sdp.SessionDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasICETrickleOption(desc *sdp.SessionDescription) bool {
	_ = "STUB: not implemented"
	return false
}

func hasTrickleOptionValue(value string) bool { _ = "STUB: not implemented"; return false }
