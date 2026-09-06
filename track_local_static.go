//go:build !js

package webrtc

import (
	"sync"

	"github.com/pion/rtp"
	"github.com/pion/webrtc/v4/pkg/media"
)

type trackBinding struct {
	id                          string
	ssrc, ssrcRTX, ssrcFEC      SSRC
	payloadType, payloadTypeRTX PayloadType
	writeStream                 TrackLocalWriter
}

type TrackLocalStaticRTP struct {
	mu                sync.RWMutex
	bindings          []trackBinding
	codec             RTPCodecCapability
	payloader         func(RTPCodecCapability) (rtp.Payloader, error)
	id, rid, streamID string
	initalTimestamp   *uint32
	initialSeqNumber  *uint16
}

func NewTrackLocalStaticRTP(
	c RTPCodecCapability,
	id, streamID string,
	options ...func(*TrackLocalStaticRTP),
) (*TrackLocalStaticRTP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithRTPStreamID(rid string) func(*TrackLocalStaticRTP) { _ = "STUB: not implemented"; return nil }

func WithPayloader(h func(RTPCodecCapability) (rtp.Payloader, error)) func(*TrackLocalStaticRTP) {
	_ = "STUB: not implemented"
	return nil
}

func WithRTPTimestamp(timestamp uint32) func(*TrackLocalStaticRTP) {
	_ = "STUB: not implemented"
	return nil
}

func WithRTPSequenceNumber(sequenceNumber uint16) func(*TrackLocalStaticRTP) {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrackLocalStaticRTP) Bind(trackContext TrackLocalContext) (RTPCodecParameters, error) {
	_ = "STUB: not implemented"
	return *new(RTPCodecParameters), nil
}

func (s *TrackLocalStaticRTP) Unbind(t TrackLocalContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrackLocalStaticRTP) ID() string { _ = "STUB: not implemented"; return "" }

func (s *TrackLocalStaticRTP) StreamID() string { _ = "STUB: not implemented"; return "" }

func (s *TrackLocalStaticRTP) RID() string { _ = "STUB: not implemented"; return "" }

func (s *TrackLocalStaticRTP) Kind() RTPCodecType {
	_ = "STUB: not implemented"
	return *new(RTPCodecType)
}

func (s *TrackLocalStaticRTP) Codec() RTPCodecCapability {
	_ = "STUB: not implemented"
	return *new(RTPCodecCapability)
}

var rtpPacketPool = sync.Pool{
	New: func() any {
		return &rtp.Packet{}
	},
}

func resetPacketPoolAllocation(localPacket *rtp.Packet) { _ = "STUB: not implemented"; return }

func getPacketAllocationFromPool() *rtp.Packet { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func (s *TrackLocalStaticRTP) WriteRTP(p *rtp.Packet) error { _ = "STUB: not implemented"; return nil }

func (s *TrackLocalStaticRTP) writeRTP(packet *rtp.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrackLocalStaticRTP) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type TrackLocalStaticSample struct {
	mu         sync.Mutex
	packetizer rtp.Packetizer
	sequencer  rtp.Sequencer
	rtpTrack   *TrackLocalStaticRTP
	clockRate  float64
	remainder  float64
}

func NewTrackLocalStaticSample(
	c RTPCodecCapability,
	id, streamID string,
	options ...func(*TrackLocalStaticRTP),
) (*TrackLocalStaticSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *TrackLocalStaticSample) ID() string { _ = "STUB: not implemented"; return "" }

func (s *TrackLocalStaticSample) StreamID() string { _ = "STUB: not implemented"; return "" }

func (s *TrackLocalStaticSample) RID() string { _ = "STUB: not implemented"; return "" }

func (s *TrackLocalStaticSample) Kind() RTPCodecType {
	_ = "STUB: not implemented"
	return *new(RTPCodecType)
}

func (s *TrackLocalStaticSample) Codec() RTPCodecCapability {
	_ = "STUB: not implemented"
	return *new(RTPCodecCapability)
}

func (s *TrackLocalStaticSample) Bind(t TrackLocalContext) (RTPCodecParameters, error) {
	_ = "STUB: not implemented"
	return *new(RTPCodecParameters), nil
}

func (s *TrackLocalStaticSample) Unbind(t TrackLocalContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrackLocalStaticSample) WriteSample(sample media.Sample) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrackLocalStaticSample) GeneratePadding(samples uint32) error {
	_ = "STUB: not implemented"
	return nil
}
