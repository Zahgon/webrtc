package webrtc

type ICETransportPolicy int

type ICEGatherPolicy = ICETransportPolicy

const (
	ICETransportPolicyAll ICETransportPolicy = iota

	ICETransportPolicyRelay

	ICETransportPolicyNoHost
)

const (
	iceTransportPolicyRelayStr  = "relay"
	iceTransportPolicyNoHostStr = "nohost"
	iceTransportPolicyAllStr    = "all"
)

func NewICETransportPolicy(raw string) ICETransportPolicy {
	_ = "STUB: not implemented"
	return *new(ICETransportPolicy)
}

func (t ICETransportPolicy) String() string { _ = "STUB: not implemented"; return "" }

func (t *ICETransportPolicy) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (t ICETransportPolicy) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
