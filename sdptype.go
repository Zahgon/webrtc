package webrtc

type SDPType int

const (
	SDPTypeUnknown SDPType = iota

	SDPTypeOffer

	SDPTypePranswer

	SDPTypeAnswer

	SDPTypeRollback
)

const (
	sdpTypeOfferStr    = "offer"
	sdpTypePranswerStr = "pranswer"
	sdpTypeAnswerStr   = "answer"
	sdpTypeRollbackStr = "rollback"
)

func NewSDPType(raw string) SDPType { _ = "STUB: not implemented"; return *new(SDPType) }

func (t SDPType) String() string { _ = "STUB: not implemented"; return "" }

func (t SDPType) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *SDPType) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
