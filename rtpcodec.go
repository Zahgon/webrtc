package webrtc

type RTPCodecType int

const (
	RTPCodecTypeUnknown RTPCodecType = iota

	RTPCodecTypeAudio

	RTPCodecTypeVideo
)

func (t RTPCodecType) String() string { _ = "STUB: not implemented"; return "" }

//nolint: goconst

//nolint: goconst

func NewRTPCodecType(r string) RTPCodecType { _ = "STUB: not implemented"; return *new(RTPCodecType) }

type RTPCodecCapability struct {
	MimeType     string
	ClockRate    uint32
	Channels     uint16
	SDPFmtpLine  string
	RTCPFeedback []RTCPFeedback
}

type RTPHeaderExtensionCapability struct {
	URI string
}

type RTPHeaderExtensionParameter struct {
	URI string
	ID  int
}

type RTPCodecParameters struct {
	RTPCodecCapability
	PayloadType PayloadType

	statsID string
}

type RTPParameters struct {
	HeaderExtensions []RTPHeaderExtensionParameter
	Codecs           []RTPCodecParameters
}

type codecMatchType int

const (
	codecMatchNone    codecMatchType = 0
	codecMatchPartial codecMatchType = 1
	codecMatchExact   codecMatchType = 2
)

func codecParametersFuzzySearch(
	needle RTPCodecParameters,
	haystack []RTPCodecParameters,
) (RTPCodecParameters, codecMatchType) {
	_ = "STUB: not implemented"
	return *new(RTPCodecParameters), *new(codecMatchType)
}

func findRTXPayloadType(needle PayloadType, haystack []RTPCodecParameters) PayloadType {
	_ = "STUB: not implemented"
	return *new(PayloadType)
}

func primaryPayloadTypeForRTXExists(needle RTPCodecParameters, haystack []RTPCodecParameters) (
	isRTX bool, primaryExists bool,
) {
	_ = "STUB: not implemented"
	return false, false
}

func filterUnattachedRTX(codecs []RTPCodecParameters) []RTPCodecParameters {
	_ = "STUB: not implemented"
	return nil
}

func findFECPayloadType(haystack []RTPCodecParameters) PayloadType {
	_ = "STUB: not implemented"
	return *new(PayloadType)
}

func rtcpFeedbackIntersection(a, b []RTCPFeedback) (out []RTCPFeedback) {
	_ = "STUB: not implemented"
	return nil
}
