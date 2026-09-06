package webrtc

type RTCPMuxPolicy int

const (
	RTCPMuxPolicyUnknown RTCPMuxPolicy = iota

	RTCPMuxPolicyNegotiate

	RTCPMuxPolicyRequire
)

const (
	rtcpMuxPolicyNegotiateStr = "negotiate"
	rtcpMuxPolicyRequireStr   = "require"
)

func newRTCPMuxPolicy(raw string) RTCPMuxPolicy {
	_ = "STUB: not implemented"
	return *new(RTCPMuxPolicy)
}

func (t RTCPMuxPolicy) String() string { _ = "STUB: not implemented"; return "" }

func (t *RTCPMuxPolicy) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (t RTCPMuxPolicy) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
