package webrtc

type RTPTransceiverDirection int

const (
	RTPTransceiverDirectionUnknown RTPTransceiverDirection = iota

	RTPTransceiverDirectionSendrecv

	RTPTransceiverDirectionSendonly

	RTPTransceiverDirectionRecvonly

	RTPTransceiverDirectionInactive
)

const (
	rtpTransceiverDirectionSendrecvStr = "sendrecv"
	rtpTransceiverDirectionSendonlyStr = "sendonly"
	rtpTransceiverDirectionRecvonlyStr = "recvonly"
	rtpTransceiverDirectionInactiveStr = "inactive"
)

func NewRTPTransceiverDirection(raw string) RTPTransceiverDirection {
	_ = "STUB: not implemented"
	return *new(RTPTransceiverDirection)
}

func (t RTPTransceiverDirection) String() string { _ = "STUB: not implemented"; return "" }

func (t RTPTransceiverDirection) Revers() RTPTransceiverDirection {
	_ = "STUB: not implemented"
	return *new(RTPTransceiverDirection)
}

func haveRTPTransceiverDirectionIntersection(
	haystack []RTPTransceiverDirection,
	needle []RTPTransceiverDirection,
) bool {
	_ = "STUB: not implemented"
	return false
}
