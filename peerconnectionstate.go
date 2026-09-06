package webrtc

type PeerConnectionState int

const (
	PeerConnectionStateUnknown PeerConnectionState = iota

	PeerConnectionStateNew

	PeerConnectionStateConnecting

	PeerConnectionStateConnected

	PeerConnectionStateDisconnected

	PeerConnectionStateFailed

	PeerConnectionStateClosed
)

const (
	peerConnectionStateNewStr          = "new"
	peerConnectionStateConnectingStr   = "connecting"
	peerConnectionStateConnectedStr    = "connected"
	peerConnectionStateDisconnectedStr = "disconnected"
	peerConnectionStateFailedStr       = "failed"
	peerConnectionStateClosedStr       = "closed"
)

func newPeerConnectionState(raw string) PeerConnectionState {
	_ = "STUB: not implemented"
	return *new(PeerConnectionState)
}

func (t PeerConnectionState) String() string { _ = "STUB: not implemented"; return "" }
