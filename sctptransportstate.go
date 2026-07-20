package webrtc

type SCTPTransportState int

const (
	SCTPTransportStateUnknown SCTPTransportState = iota

	SCTPTransportStateConnecting

	SCTPTransportStateConnected

	SCTPTransportStateClosed
)

const (
	sctpTransportStateConnectingStr = "connecting"
	sctpTransportStateConnectedStr  = "connected"
	sctpTransportStateClosedStr     = "closed"
)

func newSCTPTransportState(raw string) SCTPTransportState {
	_ = "STUB: not implemented"
	return *new(SCTPTransportState)
}

func (s SCTPTransportState) String() string { _ = "STUB: not implemented"; return "" }
