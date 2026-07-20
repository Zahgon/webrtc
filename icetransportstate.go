package webrtc

import "github.com/pion/ice/v4"

type ICETransportState int

const (
	ICETransportStateUnknown ICETransportState = iota

	ICETransportStateNew

	ICETransportStateChecking

	ICETransportStateConnected

	ICETransportStateCompleted

	ICETransportStateFailed

	ICETransportStateDisconnected

	ICETransportStateClosed
)

const (
	iceTransportStateNewStr          = "new"
	iceTransportStateCheckingStr     = "checking"
	iceTransportStateConnectedStr    = "connected"
	iceTransportStateCompletedStr    = "completed"
	iceTransportStateFailedStr       = "failed"
	iceTransportStateDisconnectedStr = "disconnected"
	iceTransportStateClosedStr       = "closed"
)

func newICETransportState(raw string) ICETransportState {
	_ = "STUB: not implemented"
	return *new(ICETransportState)
}

func (c ICETransportState) String() string { _ = "STUB: not implemented"; return "" }

func newICETransportStateFromICE(i ice.ConnectionState) ICETransportState {
	_ = "STUB: not implemented"
	return *new(ICETransportState)
}

func (c ICETransportState) toICE() ice.ConnectionState {
	_ = "STUB: not implemented"
	return *new(ice.ConnectionState)
}

func (c ICETransportState) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ICETransportState) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }
