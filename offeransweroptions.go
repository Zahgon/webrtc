package webrtc

type OfferAnswerOptions struct {
	VoiceActivityDetection bool

	ICETricklingSupported bool
}

type AnswerOptions struct {
	OfferAnswerOptions
}

type OfferOptions struct {
	OfferAnswerOptions

	ICERestart bool
}
