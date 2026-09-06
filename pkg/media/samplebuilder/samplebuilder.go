package samplebuilder

import (
	"math"
	"time"

	"github.com/pion/rtp"
	"github.com/pion/webrtc/v4/pkg/media"
)

type SampleBuilder struct {
	maxLate          uint16
	maxLateTimestamp uint32
	buffer           [math.MaxUint16 + 1]*rtp.Packet
	preparedSamples  [math.MaxUint16 + 1]*media.Sample

	depacketizer rtp.Depacketizer

	sampleRate uint32

	packetReleaseHandler func(*rtp.Packet)

	filled sampleSequenceLocation

	active sampleSequenceLocation

	prepared sampleSequenceLocation

	lastSampleTimestamp *uint32

	droppedPackets uint16

	paddingPackets uint16

	packetHeadHandler func(headPacket any) any

	returnRTPHeaders bool
}

func New(maxLate uint16, depacketizer rtp.Depacketizer, sampleRate uint32, opts ...Option) *SampleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *SampleBuilder) tooOld(location sampleSequenceLocation) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SampleBuilder) fetchTimestamp(location sampleSequenceLocation) (timestamp uint32, hasData bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (s *SampleBuilder) releasePacket(i uint16) { _ = "STUB: not implemented"; return }

func (s *SampleBuilder) purgeConsumedBuffers() { _ = "STUB: not implemented"; return }

func (s *SampleBuilder) purgeConsumedLocation(consume sampleSequenceLocation, forceConsume bool) {
	_ = "STUB: not implemented"
	return
}

func (s *SampleBuilder) purgeBuffers(flush bool) { _ = "STUB: not implemented"; return }

func (s *SampleBuilder) Push(packet *rtp.Packet) { _ = "STUB: not implemented"; return }

func (s *SampleBuilder) Flush() { _ = "STUB: not implemented"; return }

const secondToNanoseconds = 1000000000

//nolint:gocognit,cyclop
func (s *SampleBuilder) buildSample(purgingBuffers bool) *media.Sample {
	_ = "STUB: not implemented"
	return nil
}

func (s *SampleBuilder) Pop() *media.Sample { _ = "STUB: not implemented"; return nil }

func seqnumDistance(x, y uint16) uint16 {
	_ = "STUB: not implemented"
	//nolint:gosec // G115
	return 0
}

func timestampDistance(x, y uint32) uint32 {
	_ = "STUB: not implemented"
	//nolint:gosec // G115
	return 0
}

type Option func(o *SampleBuilder)

func WithPacketReleaseHandler(h func(*rtp.Packet)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPacketHeadHandler(h func(headPacket any) any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMaxTimeDelay(maxLateDuration time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

//nolint:gosec // G5G115

func WithRTPHeaders(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }
