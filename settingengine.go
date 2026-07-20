//go:build !js

package webrtc

import (
	"context"
	"crypto/x509"
	"errors"
	"io"
	"net"
	"time"

	"github.com/pion/dtls/v3"
	dtlsElliptic "github.com/pion/dtls/v3/pkg/crypto/elliptic"
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
	"github.com/pion/ice/v4"
	"github.com/pion/logging"
	"github.com/pion/stun/v3"
	"github.com/pion/transport/v4"
	"github.com/pion/transport/v4/packetio"
	"golang.org/x/net/proxy"
)

type SettingEngine struct {
	ephemeralUDP struct {
		PortMin uint16
		PortMax uint16
	}
	detach struct {
		DataChannels bool
	}
	timeout struct {
		ICEDisconnectedTimeout    *time.Duration
		ICEFailedTimeout          *time.Duration
		ICEKeepaliveInterval      *time.Duration
		ICEHostAcceptanceMinWait  *time.Duration
		ICESrflxAcceptanceMinWait *time.Duration
		ICEPrflxAcceptanceMinWait *time.Duration
		ICERelayAcceptanceMinWait *time.Duration
		ICESTUNGatherTimeout      *time.Duration
	}
	renomination renominationSettings
	candidates   struct {
		ICELite                  bool
		ICENetworkTypes          []NetworkType
		InterfaceFilter          func(string) (keep bool)
		IPFilter                 func(net.IP) (keep bool)
		RemoteIPFilter           func(net.IP) (keep bool)
		NAT1To1IPs               []string
		NAT1To1IPCandidateType   ICECandidateType
		addressRewriteRules      []ice.AddressRewriteRule
		MulticastDNSMode         ice.MulticastDNSMode
		MulticastDNSHostName     string
		UsernameFragment         string
		Password                 string //nolint:gosec // not a secret.
		IncludeLoopbackCandidate bool
	}
	replayProtection struct {
		DTLS  *uint
		SRTP  *uint
		SRTCP *uint
	}
	dtls struct {
		insecureSkipHelloVerify       bool
		disableInsecureSkipVerify     bool
		retransmissionInterval        time.Duration
		ellipticCurves                []dtlsElliptic.Curve
		connectContextMaker           func() (context.Context, func())
		extendedMasterSecret          dtls.ExtendedMasterSecretType
		clientAuth                    *dtls.ClientAuthType
		clientCAs                     *x509.CertPool
		rootCAs                       *x509.CertPool
		keyLogWriter                  io.Writer
		cipherSuites                  []dtls.CipherSuiteID
		customCipherSuites            func() []dtls.CipherSuite
		clientHelloMessageHook        func(handshake.MessageClientHello) handshake.Message
		serverHelloMessageHook        func(handshake.MessageServerHello) handshake.Message
		certificateRequestMessageHook func(handshake.MessageCertificateRequest) handshake.Message
		supportedProtocols            []string
	}
	sctp struct {
		maxReceiveBufferSize uint32
		enableZeroChecksum   bool
		rtoMax               time.Duration
		maxMessageSize       uint32
		minCwnd              uint32
		fastRtxWnd           uint32
		cwndCAStep           uint32
		enableSnap           bool
	}
	sdpMediaLevelFingerprints                 bool
	answeringDTLSRole                         DTLSRole
	disableCertificateFingerprintVerification bool
	disableSRTPReplayProtection               bool
	disableSRTCPReplayProtection              bool
	net                                       transport.Net
	BufferFactory                             func(packetType packetio.BufferPacketType, ssrc uint32) io.ReadWriteCloser
	LoggerFactory                             logging.LoggerFactory
	iceTCPMux                                 ice.TCPMux
	iceUDPMux                                 ice.UDPMux
	iceProxyDialer                            proxy.Dialer
	iceDisableActiveTCP                       bool
	iceUseCandidateCheckPriority              bool
	iceBindingRequestHandler                  func(m *stun.Message, local, remote ice.Candidate, pair *ice.CandidatePair) bool //nolint:lll
	disableMediaEngineCopy                    bool
	disableMediaEngineMultipleCodecs          bool
	srtpProtectionProfiles                    []dtls.SRTPProtectionProfile
	receiveMTU                                uint
	iceMaxBindingRequests                     *uint16
	fireOnTrackBeforeFirstRTP                 bool
	disableCloseByDTLS                        bool
	dataChannelBlockWrite                     bool
	handleUndeclaredSSRCWithoutAnswer         bool
	ignoreRidPauseForRecv                     bool
}

type renominationSettings struct {
	enabled           bool
	generator         ice.NominationValueGenerator
	automatic         bool
	automaticInterval *time.Duration
	attributeType     *uint16
}

type NominationValueGenerator func() uint32

func (f NominationValueGenerator) toIce() ice.NominationValueGenerator {
	_ = "STUB: not implemented"
	return *new(ice.NominationValueGenerator)
}

type RenominationOption func(*renominationSettings)

func WithRenominationGenerator(generator NominationValueGenerator) RenominationOption {
	_ = "STUB: not implemented"
	return *new(RenominationOption)
}

func WithRenominationInterval(interval time.Duration) RenominationOption {
	_ = "STUB: not implemented"
	return *new(RenominationOption)
}

func WithRenominationNominationAttribute(attrType uint16) RenominationOption {
	_ = "STUB: not implemented"
	return *new(RenominationOption)
}

var errInvalidRenominationInterval = errors.New("renomination interval must be greater than zero")

func (e *SettingEngine) SetICERenomination(options ...RenominationOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *SettingEngine) getSCTPMaxMessageSize() uint32 { _ = "STUB: not implemented"; return 0 }

func (e *SettingEngine) getReceiveMTU() uint { _ = "STUB: not implemented"; return 0 }

func (e *SettingEngine) DetachDataChannels() { _ = "STUB: not implemented"; return }

func (e *SettingEngine) EnableDataChannelBlockWrite(nonblockWrite bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetSRTPProtectionProfiles(profiles ...dtls.SRTPProtectionProfile) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetICETimeouts(disconnectedTimeout, failedTimeout, keepAliveInterval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetHostAcceptanceMinWait(t time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetSrflxAcceptanceMinWait(t time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetPrflxAcceptanceMinWait(t time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetRelayAcceptanceMinWait(t time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetSTUNGatherTimeout(t time.Duration) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetEphemeralUDPPortRange(portMin, portMax uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *SettingEngine) SetLite(lite bool) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetNetworkTypes(candidateTypes []NetworkType) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetInterfaceFilter(filter func(string) (keep bool)) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetIPFilter(filter func(net.IP) (keep bool)) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetRemoteIPFilter(filter func(net.IP) (keep bool)) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetNAT1To1IPs(ips []string, candidateType ICECandidateType) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetICEAddressRewriteRules(rules ...ICEAddressRewriteRule) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *SettingEngine) SetIncludeLoopbackCandidate(include bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetAnsweringDTLSRole(role DTLSRole) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *SettingEngine) SetNet(net transport.Net) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetICEMulticastDNSMode(multicastDNSMode ice.MulticastDNSMode) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetMulticastDNSHostName(hostName string) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetICECredentials(usernameFragment, password string) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) DisableCertificateFingerprintVerification(isDisabled bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSReplayProtectionWindow(n uint) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetSRTPReplayProtectionWindow(n uint) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetSRTCPReplayProtectionWindow(n uint) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) DisableSRTPReplayProtection(isDisabled bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) DisableSRTCPReplayProtection(isDisabled bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetSDPMediaLevelFingerprints(sdpMediaLevelFingerprints bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetICETCPMux(tcpMux ice.TCPMux) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetICEUDPMux(udpMux ice.UDPMux) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetICEProxyDialer(d proxy.Dialer) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetICEMaxBindingRequests(d uint16) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetICEUseCandidateCheckPriority(enabled bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) DisableActiveTCP(isDisabled bool) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) DisableMediaEngineCopy(isDisabled bool) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) DisableMediaEngineMultipleCodecs(isDisabled bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetReceiveMTU(receiveMTU uint) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetDTLSRetransmissionInterval(interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSInsecureSkipHelloVerify(skip bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSDisableInsecureSkipVerify(disable bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSEllipticCurves(ellipticCurves ...dtlsElliptic.Curve) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSConnectContextMaker(connectContextMaker func() (context.Context, func())) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSExtendedMasterSecret(extendedMasterSecret dtls.ExtendedMasterSecretType) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSClientAuth(clientAuth dtls.ClientAuthType) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSClientCAs(clientCAs *x509.CertPool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSRootCAs(rootCAs *x509.CertPool) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetDTLSKeyLogWriter(writer io.Writer) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetSCTPMaxReceiveBufferSize(maxReceiveBufferSize uint32) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) EnableSCTPZeroChecksum(isEnabled bool) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) EnableSctpSnap(isEnabled bool) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetSCTPMaxMessageSize(maxMessageSize uint32) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSCipherSuites(cipherSuites ...dtls.CipherSuiteID) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSCustomerCipherSuites(customCipherSuites func() []dtls.CipherSuite) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSClientHelloMessageHook(hook func(handshake.MessageClientHello) handshake.Message) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSServerHelloMessageHook(hook func(handshake.MessageServerHello) handshake.Message) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSCertificateRequestMessageHook(
	hook func(handshake.MessageCertificateRequest) handshake.Message,
) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetDTLSSupportedProtocols(protocols ...string) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetSCTPRTOMax(rtoMax time.Duration) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetSCTPMinCwnd(minCwnd uint32) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetSCTPFastRtxWnd(fastRtxWnd uint32) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetSCTPCwndCAStep(cwndCAStep uint32) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetICEBindingRequestHandler(
	bindingRequestHandler func(m *stun.Message, local, remote ice.Candidate, pair *ice.CandidatePair) bool,
) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetFireOnTrackBeforeFirstRTP(fireOnTrackBeforeFirstRTP bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) DisableCloseByDTLS(isEnabled bool) { _ = "STUB: not implemented"; return }

func (e *SettingEngine) SetHandleUndeclaredSSRCWithoutAnswer(handleUndeclaredSSRCWithoutAnswer bool) {
	_ = "STUB: not implemented"
	return
}

func (e *SettingEngine) SetIgnoreRidPauseForRecv(ignoreRidPauseForRecv bool) {
	_ = "STUB: not implemented"
	return
}
