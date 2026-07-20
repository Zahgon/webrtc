package webrtc

type ICEGathererState uint32

const (
	ICEGathererStateUnknown ICEGathererState = iota

	ICEGathererStateNew

	ICEGathererStateGathering

	ICEGathererStateComplete

	ICEGathererStateClosed
)

func (s ICEGathererState) String() string { _ = "STUB: not implemented"; return "" }

func atomicStoreICEGathererState(state *ICEGathererState, newState ICEGathererState) {
	_ = "STUB: not implemented"
	return
}

func atomicLoadICEGathererState(state *ICEGathererState) ICEGathererState {
	_ = "STUB: not implemented"
	return *new(ICEGathererState)
}
