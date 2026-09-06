package webrtc

const (
	TypeRTCPFBTransportCC = "transport-cc"

	TypeRTCPFBGoogREMB = "goog-remb"

	TypeRTCPFBACK = "ack"

	TypeRTCPFBCCM = "ccm"

	TypeRTCPFBNACK = "nack"
)

type RTCPFeedback struct {
	Type string

	Parameter string
}
