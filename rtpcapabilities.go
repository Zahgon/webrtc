package webrtc

type RTPCapabilities struct {
	Codecs           []RTPCodecCapability
	HeaderExtensions []RTPHeaderExtensionCapability
}
