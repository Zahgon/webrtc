package webrtc

type ICEComponent int

const (
	ICEComponentUnknown ICEComponent = iota

	ICEComponentRTP

	ICEComponentRTCP
)

const (
	iceComponentRTPStr  = "rtp"
	iceComponentRTCPStr = "rtcp"
)

func newICEComponent(raw string) ICEComponent { _ = "STUB: not implemented"; return *new(ICEComponent) }

func (t ICEComponent) String() string { _ = "STUB: not implemented"; return "" }
