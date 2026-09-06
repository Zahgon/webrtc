package webrtc

type DTLSTransportState int

const (
	DTLSTransportStateUnknown DTLSTransportState = iota

	DTLSTransportStateNew

	DTLSTransportStateConnecting

	DTLSTransportStateConnected

	DTLSTransportStateClosed

	DTLSTransportStateFailed
)

const (
	dtlsTransportStateNewStr        = "new"
	dtlsTransportStateConnectingStr = "connecting"
	dtlsTransportStateConnectedStr  = "connected"
	dtlsTransportStateClosedStr     = "closed"
	dtlsTransportStateFailedStr     = "failed"
)

func newDTLSTransportState(raw string) DTLSTransportState {
	_ = "STUB: not implemented"
	return *new(DTLSTransportState)
}

func (t DTLSTransportState) String() string { _ = "STUB: not implemented"; return "" }

func (t DTLSTransportState) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DTLSTransportState) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }
