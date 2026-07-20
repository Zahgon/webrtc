//go:build !js

package webrtc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"sync"
	"sync/atomic"

	"github.com/pion/dtls/v3"
	"github.com/pion/interceptor"
	"github.com/pion/logging"
	"github.com/pion/rtcp"
	"github.com/pion/srtp/v3"
	"github.com/pion/webrtc/v4/internal/mux"
)

type DTLSTransport struct {
	lock sync.RWMutex

	iceTransport          *ICETransport
	certificates          []Certificate
	remoteParameters      DTLSParameters
	remoteCertificate     []byte
	state                 DTLSTransportState
	srtpProtectionProfile srtp.ProtectionProfile

	onStateChangeHandler   func(DTLSTransportState)
	internalOnCloseHandler func()

	conn *dtls.Conn

	srtpSession, srtcpSession   atomic.Value
	srtpEndpoint, srtcpEndpoint *mux.Endpoint
	simulcastStreams            []simulcastStreamPair
	srtpReady                   chan struct{}

	dtlsMatcher mux.MatchFunc

	api *API
	log logging.LeveledLogger
}

type simulcastStreamPair struct {
	srtp  *srtp.ReadStreamSRTP
	srtcp *srtp.ReadStreamSRTCP
}

type streamsForSSRCResult struct {
	rtpReadStream   *srtp.ReadStreamSRTP
	rtpInterceptor  interceptor.RTPReader
	rtcpReadStream  *srtp.ReadStreamSRTCP
	rtcpInterceptor interceptor.RTCPReader
}

func (api *API) NewDTLSTransport(transport *ICETransport, certificates []Certificate) (*DTLSTransport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DTLSTransport) ICETransport() *ICETransport { _ = "STUB: not implemented"; return nil }

func (t *DTLSTransport) onStateChange(state DTLSTransportState) { _ = "STUB: not implemented"; return }

func (t *DTLSTransport) OnStateChange(f func(DTLSTransportState)) {
	_ = "STUB: not implemented"
	return
}

func (t *DTLSTransport) State() DTLSTransportState {
	_ = "STUB: not implemented"
	return *new(DTLSTransportState)
}

func (t *DTLSTransport) WriteRTCP(pkts []rtcp.Packet) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t *DTLSTransport) GetLocalParameters() (DTLSParameters, error) {
	_ = "STUB: not implemented"
	return *new(DTLSParameters), nil
}

func (t *DTLSTransport) GetRemoteCertificate() []byte { _ = "STUB: not implemented"; return nil }

func (t *DTLSTransport) startSRTP() error { _ = "STUB: not implemented"; return nil }

func (t *DTLSTransport) getSRTPSession() (*srtp.SessionSRTP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DTLSTransport) getSRTCPSession() (*srtp.SessionSRTCP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DTLSTransport) role() DTLSRole { _ = "STUB: not implemented"; return *new(DTLSRole) }

func (t *DTLSTransport) Start(remoteParameters DTLSParameters) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) StartContext(ctx context.Context, remoteParameters DTLSParameters) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) start(remoteParameters DTLSParameters, handshake func(*dtls.Conn) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) prepareStart(remoteParameters DTLSParameters) (DTLSRole, tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(DTLSRole), *new(tls.Certificate), nil
}

func (t *DTLSTransport) dtlsSharedOptions(certificate tls.Certificate) []dtls.Option {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // G115

func (t *DTLSTransport) srtpProtectionProfiles() []dtls.SRTPProtectionProfile {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) verifyPeerCertificateFunc() func([][]byte, [][]*x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) connectDTLS(
	dtlsEndpoint *mux.Endpoint,
	role DTLSRole,
	sharedOpts []dtls.Option,
) (*dtls.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DTLSTransport) toDTLSServerOptions(sharedOpts []dtls.Option) []dtls.ServerOption {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) toDTLSClientOptions(sharedOpts []dtls.Option) []dtls.ClientOption {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) handshakeDTLS(dtlsConn *dtls.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) completeStart(dtlsConn *dtls.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) failStart(err error) error { _ = "STUB: not implemented"; return nil }

func srtpProtectionProfileFromDTLSConn(dtlsConn *dtls.Conn) (srtp.ProtectionProfile, error) {
	_ = "STUB: not implemented"
	return *new(srtp.ProtectionProfile), nil
}

func srtpProtectionProfileFromDTLS(srtpProfile dtls.SRTPProtectionProfile) (srtp.ProtectionProfile, error) {
	_ = "STUB: not implemented"
	return *new(srtp.ProtectionProfile), nil
}

func (t *DTLSTransport) Stop() error { _ = "STUB: not implemented"; return nil }

func (t *DTLSTransport) validateFingerPrint(remoteCert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *DTLSTransport) ensureICEConn() error { _ = "STUB: not implemented"; return nil }

func (t *DTLSTransport) storeSimulcastStream(
	srtpReadStream *srtp.ReadStreamSRTP,
	srtcpReadStream *srtp.ReadStreamSRTCP,
) {
	_ = "STUB: not implemented"
	return
}

func (t *DTLSTransport) streamsForSSRC(
	ssrc SSRC,
	streamInfo interceptor.StreamInfo,
) (*streamsForSSRCResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
