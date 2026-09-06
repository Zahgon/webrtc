package webrtc

type DataChannelInit struct {
	Ordered *bool

	MaxPacketLifeTime *uint16

	MaxRetransmits *uint16

	Protocol *string

	Negotiated *bool

	ID *uint16
}
