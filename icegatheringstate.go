package webrtc

type ICEGatheringState int

const (
	ICEGatheringStateUnknown ICEGatheringState = iota

	ICEGatheringStateNew

	ICEGatheringStateGathering

	ICEGatheringStateComplete
)

const (
	iceGatheringStateNewStr       = "new"
	iceGatheringStateGatheringStr = "gathering"
	iceGatheringStateCompleteStr  = "complete"
)

func NewICEGatheringState(raw string) ICEGatheringState {
	_ = "STUB: not implemented"
	return *new(ICEGatheringState)
}

func (t ICEGatheringState) String() string { _ = "STUB: not implemented"; return "" }
