//go:build !js

package webrtc

import (
	"context"
	"sync"
	"time"
)

func (r StatsReport) GetConnectionStats(conn *PeerConnection) (PeerConnectionStats, bool) {
	_ = "STUB: not implemented"
	return *new(PeerConnectionStats), false
}

func (r StatsReport) GetDataChannelStats(dc *DataChannel) (DataChannelStats, bool) {
	_ = "STUB: not implemented"
	return *new(DataChannelStats), false
}

func (r StatsReport) GetICECandidateStats(c *ICECandidate) (ICECandidateStats, bool) {
	_ = "STUB: not implemented"
	return *new(ICECandidateStats), false
}

func (r StatsReport) GetICECandidatePairStats(c *ICECandidatePair) (ICECandidatePairStats, bool) {
	_ = "STUB: not implemented"
	return *new(ICECandidatePairStats), false
}

func (r StatsReport) GetCertificateStats(c *Certificate) (CertificateStats, bool) {
	_ = "STUB: not implemented"
	return *new(CertificateStats), false
}

func (r StatsReport) GetCodecStats(c *RTPCodecParameters) (CodecStats, bool) {
	_ = "STUB: not implemented"
	return *new(CodecStats), false
}

type AudioPlayoutStatsProvider interface {
	AddTrack(track *TrackRemote) error

	RemoveTrack(track *TrackRemote)

	Snapshot(now time.Time) (AudioPlayoutStats, bool)
}

type trackContext struct {
	cancel context.CancelFunc
}

type defaultAudioPlayoutStatsProvider struct {
	mu sync.Mutex

	stats           AudioPlayoutStats
	lastSynthesized bool
	tracks          map[*TrackRemote]*trackContext
}

func NewAudioPlayoutStatsProvider(id string) *defaultAudioPlayoutStatsProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p *defaultAudioPlayoutStatsProvider) Accumulate(
	samples int, sampleRate uint32, deviceDelay time.Duration, synthesized bool,
) {
	_ = "STUB: not implemented"
	return
}

func (p *defaultAudioPlayoutStatsProvider) Snapshot(now time.Time) (AudioPlayoutStats, bool) {
	_ = "STUB: not implemented"
	return *new(AudioPlayoutStats), false
}

func (p *defaultAudioPlayoutStatsProvider) AddTrack(track *TrackRemote) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *defaultAudioPlayoutStatsProvider) RemoveTrack(track *TrackRemote) {
	_ = "STUB: not implemented"
	return
}

func (p *defaultAudioPlayoutStatsProvider) removeTrackInternal(track *TrackRemote) {
	_ = "STUB: not implemented"
	return
}
