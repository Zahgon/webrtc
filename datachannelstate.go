package webrtc

type DataChannelState int

const (
	DataChannelStateUnknown DataChannelState = iota

	DataChannelStateConnecting

	DataChannelStateOpen

	DataChannelStateClosing

	DataChannelStateClosed
)

const (
	dataChannelStateConnectingStr = "connecting"
	dataChannelStateOpenStr       = "open"
	dataChannelStateClosingStr    = "closing"
	dataChannelStateClosedStr     = "closed"
)

func newDataChannelState(raw string) DataChannelState {
	_ = "STUB: not implemented"
	return *new(DataChannelState)
}

func (t DataChannelState) String() string { _ = "STUB: not implemented"; return "" }

func (t DataChannelState) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *DataChannelState) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }
