package webrtc

import (
	"github.com/pion/interceptor"
	"github.com/pion/rtp"
)

type TrackLocalWriter interface {
	WriteRTP(header *rtp.Header, payload []byte) (int, error)

	Write(b []byte) (int, error)
}

type TrackLocalContext interface {
	CodecParameters() []RTPCodecParameters

	HeaderExtensions() []RTPHeaderExtensionParameter

	SSRC() SSRC

	SSRCRetransmission() SSRC

	SSRCForwardErrorCorrection() SSRC

	WriteStream() TrackLocalWriter

	ID() string

	RTCPReader() interceptor.RTCPReader
}

type baseTrackLocalContext struct {
	id                     string
	params                 RTPParameters
	ssrc, ssrcRTX, ssrcFEC SSRC
	writeStream            TrackLocalWriter
	rtcpInterceptor        interceptor.RTCPReader
}

func (t *baseTrackLocalContext) CodecParameters() []RTPCodecParameters {
	_ = "STUB: not implemented"
	return nil
}

func (t *baseTrackLocalContext) HeaderExtensions() []RTPHeaderExtensionParameter {
	_ = "STUB: not implemented"
	return nil
}

func (t *baseTrackLocalContext) SSRC() SSRC { _ = "STUB: not implemented"; return *new(SSRC) }

func (t *baseTrackLocalContext) SSRCRetransmission() SSRC {
	_ = "STUB: not implemented"
	return *new(SSRC)
}

func (t *baseTrackLocalContext) SSRCForwardErrorCorrection() SSRC {
	_ = "STUB: not implemented"
	return *new(SSRC)
}

func (t *baseTrackLocalContext) WriteStream() TrackLocalWriter {
	_ = "STUB: not implemented"
	return *new(TrackLocalWriter)
}

func (t *baseTrackLocalContext) ID() string { _ = "STUB: not implemented"; return "" }

func (t *baseTrackLocalContext) RTCPReader() interceptor.RTCPReader {
	_ = "STUB: not implemented"
	return *new(interceptor.RTCPReader)
}

type TrackLocal interface {
	Bind(TrackLocalContext) (RTPCodecParameters, error)

	Unbind(TrackLocalContext) error

	ID() string

	RID() string

	StreamID() string

	Kind() RTPCodecType
}
