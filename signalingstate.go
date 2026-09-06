package webrtc

type stateChangeOp int

const (
	stateChangeOpSetLocal stateChangeOp = iota + 1
	stateChangeOpSetRemote
)

func (op stateChangeOp) String() string { _ = "STUB: not implemented"; return "" }

type SignalingState int32

const (
	SignalingStateUnknown SignalingState = iota

	SignalingStateStable

	SignalingStateHaveLocalOffer

	SignalingStateHaveRemoteOffer

	SignalingStateHaveLocalPranswer

	SignalingStateHaveRemotePranswer

	SignalingStateClosed
)

const (
	signalingStateStableStr             = "stable"
	signalingStateHaveLocalOfferStr     = "have-local-offer"
	signalingStateHaveRemoteOfferStr    = "have-remote-offer"
	signalingStateHaveLocalPranswerStr  = "have-local-pranswer"
	signalingStateHaveRemotePranswerStr = "have-remote-pranswer"
	signalingStateClosedStr             = "closed"
)

func newSignalingState(raw string) SignalingState {
	_ = "STUB: not implemented"
	return *new(SignalingState)
}

func (t SignalingState) String() string { _ = "STUB: not implemented"; return "" }

func (t *SignalingState) Get() SignalingState {
	_ = "STUB: not implemented"
	return *new(SignalingState)
}

func (t *SignalingState) Set(state SignalingState) { _ = "STUB: not implemented"; return }

//nolint:gocognit,cyclop
func checkNextSignalingState(cur, next SignalingState, op stateChangeOp, sdpType SDPType) (SignalingState, error) {
	_ = "STUB: not implemented"
	return *new(SignalingState), nil
}
