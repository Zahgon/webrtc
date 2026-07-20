package webrtc

type RTPTransceiverInit struct {
	Direction     RTPTransceiverDirection
	SendEncodings []RTPEncodingParameters
}
