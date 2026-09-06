package webrtc

type SCTPCapabilities struct {
	MaxMessageSize uint32 `json:"maxMessageSize"`

	sctpInit string
}
