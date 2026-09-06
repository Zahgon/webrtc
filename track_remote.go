//go:build !js

package webrtc

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/pion/interceptor"
	"github.com/pion/rtp"
)

type peekedPacket struct {
	payload    []byte
	attributes interceptor.Attributes
}

type TrackRemote struct {
	mu                  sync.RWMutex
	repairReadRequested atomic.Bool

	id       string
	streamID string

	payloadType PayloadType
	kind        RTPCodecType
	ssrc        SSRC
	rtxSsrc     SSRC
	codec       RTPCodecParameters
	params      RTPParameters
	rid         string

	receiver *RTPReceiver

	peekedPackets []*peekedPacket

	audioPlayoutStatsProviders []AudioPlayoutStatsProvider
}

func newTrackRemote(kind RTPCodecType, ssrc, rtxSsrc SSRC, rid string, receiver *RTPReceiver) *TrackRemote {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrackRemote) ID() string { _ = "STUB: not implemented"; return "" }

func (t *TrackRemote) RID() string { _ = "STUB: not implemented"; return "" }

func (t *TrackRemote) PayloadType() PayloadType {
	_ = "STUB: not implemented"
	return *new(PayloadType)
}

func (t *TrackRemote) Kind() RTPCodecType { _ = "STUB: not implemented"; return *new(RTPCodecType) }

func (t *TrackRemote) StreamID() string { _ = "STUB: not implemented"; return "" }

func (t *TrackRemote) SSRC() SSRC { _ = "STUB: not implemented"; return *new(SSRC) }

func (t *TrackRemote) Msid() string { _ = "STUB: not implemented"; return "" }

func (t *TrackRemote) Codec() RTPCodecParameters {
	_ = "STUB: not implemented"
	return *new(RTPCodecParameters)
}

func (t *TrackRemote) Read(b []byte) (n int, attributes interceptor.Attributes, err error) {
	_ = "STUB: not implemented"
	return 0, *new(interceptor.Attributes), nil
}

func (t *TrackRemote) read(b []byte) (n int, attributes interceptor.Attributes, err error) {
	_ = "STUB: not implemented"
	return 0, *new(interceptor.Attributes), nil
}

func (t *TrackRemote) checkAndUpdateTrack(b []byte) error { _ = "STUB: not implemented"; return nil }

func (t *TrackRemote) ReadRTP() (*rtp.Packet, interceptor.Attributes, error) {
	_ = "STUB: not implemented"
	return nil, *new(interceptor.Attributes), nil
}

func (t *TrackRemote) peek(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (t *TrackRemote) SetReadDeadline(deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrackRemote) RtxSSRC() SSRC { _ = "STUB: not implemented"; return *new(SSRC) }

func (t *TrackRemote) HasRTX() bool { _ = "STUB: not implemented"; return false }

func (t *TrackRemote) addProvider(provider AudioPlayoutStatsProvider) {
	_ = "STUB: not implemented"
	return
}

func (t *TrackRemote) removeProvider(provider AudioPlayoutStatsProvider) {
	_ = "STUB: not implemented"
	return
}

func (t *TrackRemote) pullAudioPlayoutStats(now time.Time) []AudioPlayoutStats {
	_ = "STUB: not implemented"
	return nil
}

func (t *TrackRemote) setRtxSSRC(ssrc SSRC) { _ = "STUB: not implemented"; return }
