//go:build !js

package webrtc

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/pion/ice/v4"
	"github.com/pion/logging"
	"github.com/pion/webrtc/v4/internal/mux"
)

type ICETransport struct {
	lock sync.RWMutex

	role ICERole

	onConnectionStateChangeHandler         atomic.Value
	internalOnConnectionStateChangeHandler atomic.Value
	onSelectedCandidatePairChangeHandler   atomic.Value

	state atomic.Value

	gatherer *ICEGatherer
	conn     *ice.Conn
	mux      *mux.Mux

	ctxCancel func()

	loggerFactory logging.LoggerFactory

	log logging.LeveledLogger
}

func (t *ICETransport) GetSelectedCandidatePair() (*ICECandidatePair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nilnil

func (t *ICETransport) GetSelectedCandidatePairStats() (ICECandidatePairStats, bool) {
	_ = "STUB: not implemented"
	return *new(ICECandidatePairStats), false
}

func NewICETransport(gatherer *ICEGatherer, loggerFactory logging.LoggerFactory) *ICETransport {
	_ = "STUB: not implemented"
	return nil
}

func (t *ICETransport) Start(gatherer *ICEGatherer, params ICEParameters, role *ICERole) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:cyclop
func (t *ICETransport) StartContext(
	ctx context.Context,
	gatherer *ICEGatherer,
	params ICEParameters,
	role *ICERole,
) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // G115

func (t *ICETransport) restart() error { _ = "STUB: not implemented"; return nil }

func (t *ICETransport) Stop() error { _ = "STUB: not implemented"; return nil }

func (t *ICETransport) GracefulStop() error { _ = "STUB: not implemented"; return nil }

func (t *ICETransport) stop(shouldGracefullyClose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *ICETransport) OnSelectedCandidatePairChange(f func(*ICECandidatePair)) {
	_ = "STUB: not implemented"
	return
}

func (t *ICETransport) onSelectedCandidatePairChange(pair *ICECandidatePair) {
	_ = "STUB: not implemented"
	return
}

func (t *ICETransport) OnConnectionStateChange(f func(ICETransportState)) {
	_ = "STUB: not implemented"
	return
}

func (t *ICETransport) onConnectionStateChange(state ICETransportState) {
	_ = "STUB: not implemented"
	return
}

func (t *ICETransport) Role() ICERole { _ = "STUB: not implemented"; return *new(ICERole) }

func (t *ICETransport) SetRemoteCandidates(remoteCandidates []ICECandidate) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *ICETransport) AddRemoteCandidate(remoteCandidate *ICECandidate) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *ICETransport) State() ICETransportState {
	_ = "STUB: not implemented"
	return *new(ICETransportState)
}

func (t *ICETransport) GetLocalParameters() (ICEParameters, error) {
	_ = "STUB: not implemented"
	return *new(ICEParameters), nil
}

func (t *ICETransport) GetRemoteParameters() (ICEParameters, error) {
	_ = "STUB: not implemented"
	return *new(ICEParameters), nil
}

func (t *ICETransport) setState(i ICETransportState) { _ = "STUB: not implemented"; return }

func (t *ICETransport) newEndpoint(f mux.MatchFunc) *mux.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

func (t *ICETransport) ensureGatherer() error { _ = "STUB: not implemented"; return nil }

func (t *ICETransport) Stats() TransportStats {
	_ = "STUB: not implemented"
	return *new(TransportStats)
}

func (t *ICETransport) collectStats(collector *statsReportCollector) {
	_ = "STUB: not implemented"
	return
}

func (t *ICETransport) haveRemoteCredentialsChange(newUfrag, newPwd string) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *ICETransport) setRemoteCredentials(newUfrag, newPwd string) error {
	_ = "STUB: not implemented"
	return nil
}
