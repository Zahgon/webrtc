//go:build !js

package webrtc

import (
	"sync"
	"sync/atomic"

	"github.com/pion/ice/v4"
	"github.com/pion/logging"
	"github.com/pion/stun/v3"
)

type ICEGatherer struct {
	lock  sync.RWMutex
	log   logging.LeveledLogger
	state ICEGathererState

	validatedServers []*stun.URI
	gatherPolicy     ICETransportPolicy

	agent *ice.Agent

	onLocalCandidateHandler atomic.Value
	onStateChangeHandler    atomic.Value

	onGatheringCompleteHandler atomic.Value

	api *API

	sdpMid        atomic.Value
	sdpMLineIndex atomic.Uint32

	candidatePoolLock    sync.Mutex
	candidatePool        []ice.Candidate
	iceCandidatePoolSize uint8
}

type ICEAddressRewriteMode byte

const (
	ICEAddressRewriteModeUnspecified ICEAddressRewriteMode = iota
	ICEAddressRewriteReplace
	ICEAddressRewriteAppend
)

func (r ICEAddressRewriteMode) toICE() ice.AddressRewriteMode {
	_ = "STUB: not implemented"
	return *new(ice.AddressRewriteMode)
}

type ICEAddressRewriteRule struct {
	External        []string
	Local           string
	Iface           string
	CIDR            string
	AsCandidateType ICECandidateType
	Mode            ICEAddressRewriteMode
	Networks        []NetworkType
}

func (r ICEAddressRewriteRule) toICE() ice.AddressRewriteRule {
	_ = "STUB: not implemented"
	return *new(ice.AddressRewriteRule)
}

func (api *API) NewICEGatherer(opts ICEGatherOptions) (*ICEGatherer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *ICEGatherer) updateServers(servers []ICEServer, policy ICETransportPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *ICEGatherer) validatedServersCount() int { _ = "STUB: not implemented"; return 0 }

func (g *ICEGatherer) createAgent() error { _ = "STUB: not implemented"; return nil }

func (g *ICEGatherer) buildAgentOptions() ([]ice.AgentOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *ICEGatherer) resolveCandidateTypes() []ice.CandidateType {
	_ = "STUB: not implemented"
	return nil
}

func (g *ICEGatherer) resolveNAT1To1CandidateType() ice.CandidateType {
	_ = "STUB: not implemented"
	return *new(ice.CandidateType)
}

func (g *ICEGatherer) sanitizedMDNSMode() ice.MulticastDNSMode {
	_ = "STUB: not implemented"
	return *new(ice.MulticastDNSMode)
}

func (g *ICEGatherer) baseAgentOptions(mDNSMode ice.MulticastDNSMode) []ice.AgentOption {
	_ = "STUB: not implemented"
	return nil
}

func (g *ICEGatherer) credentialOptions() []ice.AgentOption { _ = "STUB: not implemented"; return nil }

func (g *ICEGatherer) addressRewriteOptions(candidateType ice.CandidateType) ([]ice.AgentOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *ICEGatherer) timeoutOptions() []ice.AgentOption { _ = "STUB: not implemented"; return nil }

func (g *ICEGatherer) miscOptions() []ice.AgentOption { _ = "STUB: not implemented"; return nil }

func (g *ICEGatherer) renominationOptions() []ice.AgentOption {
	_ = "STUB: not implemented"
	return nil
}

func legacyNAT1To1AddressRewriteRules(ips []string, candidateType ice.CandidateType) []ice.AddressRewriteRule {
	_ = "STUB: not implemented"
	return nil
}

func (g *ICEGatherer) Gather() error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

//nolint:gosec // G115

func (g *ICEGatherer) setMediaStreamIdentification(mid string, mLineIndex uint16) {
	_ = "STUB: not implemented"
	return
}

func (g *ICEGatherer) flushCandidates() { _ = "STUB: not implemented"; return }

//nolint:gosec // G115

func (g *ICEGatherer) Close() error { _ = "STUB: not implemented"; return nil }

func (g *ICEGatherer) GracefulClose() error { _ = "STUB: not implemented"; return nil }

func (g *ICEGatherer) close(shouldGracefullyClose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *ICEGatherer) GetLocalParameters() (ICEParameters, error) {
	_ = "STUB: not implemented"
	return *new(ICEParameters), nil
}

func (g *ICEGatherer) GetLocalCandidates() ([]ICECandidate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // G115

func (g *ICEGatherer) OnLocalCandidate(f func(*ICECandidate)) { _ = "STUB: not implemented"; return }

func (g *ICEGatherer) OnStateChange(f func(ICEGathererState)) { _ = "STUB: not implemented"; return }

func (g *ICEGatherer) State() ICEGathererState {
	_ = "STUB: not implemented"
	return *new(ICEGathererState)
}

func (g *ICEGatherer) setState(s ICEGathererState) { _ = "STUB: not implemented"; return }

func (g *ICEGatherer) getAgent() *ice.Agent { _ = "STUB: not implemented"; return nil }

func (g *ICEGatherer) collectStats(collector *statsReportCollector) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // G115, no overflow, port

//nolint:gosec

//nolint:gosec // G115, no overflow, port

//nolint:gosec // G115

func (g *ICEGatherer) getSelectedCandidatePairStats() (ICECandidatePairStats, bool) {
	_ = "STUB: not implemented"
	return *new(ICECandidatePairStats), false
}
