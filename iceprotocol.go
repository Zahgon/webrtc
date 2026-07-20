package webrtc

type ICEProtocol int

const (
	ICEProtocolUnknown ICEProtocol = iota

	ICEProtocolUDP

	ICEProtocolTCP
)

const (
	iceProtocolUDPStr = "udp"
	iceProtocolTCPStr = "tcp"
)

func NewICEProtocol(raw string) (ICEProtocol, error) {
	_ = "STUB: not implemented"
	return *new(ICEProtocol), nil
}

func (t ICEProtocol) String() string { _ = "STUB: not implemented"; return "" }
