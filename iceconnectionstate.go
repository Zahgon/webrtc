package webrtc

type ICEConnectionState int

const (
	ICEConnectionStateUnknown ICEConnectionState = iota

	ICEConnectionStateNew

	ICEConnectionStateChecking

	ICEConnectionStateConnected

	ICEConnectionStateCompleted

	ICEConnectionStateDisconnected

	ICEConnectionStateFailed

	ICEConnectionStateClosed
)

const (
	iceConnectionStateNewStr          = "new"
	iceConnectionStateCheckingStr     = "checking"
	iceConnectionStateConnectedStr    = "connected"
	iceConnectionStateCompletedStr    = "completed"
	iceConnectionStateDisconnectedStr = "disconnected"
	iceConnectionStateFailedStr       = "failed"
	iceConnectionStateClosedStr       = "closed"
)

func NewICEConnectionState(raw string) ICEConnectionState {
	_ = "STUB: not implemented"
	return *new(ICEConnectionState)
}

func (c ICEConnectionState) String() string { _ = "STUB: not implemented"; return "" }
