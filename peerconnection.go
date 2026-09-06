//go:build !js

package webrtc

import (
	"sync"
	"sync/atomic"

	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/stats"
	"github.com/pion/logging"
	"github.com/pion/rtcp"
	"github.com/pion/sdp/v3"
	"github.com/pion/srtp/v3"
)

type PeerConnection struct {
	id string
	mu sync.RWMutex

	sdpOrigin sdp.Origin

	ops *operations

	configuration Configuration

	currentLocalDescription  *SessionDescription
	pendingLocalDescription  *SessionDescription
	currentRemoteDescription *SessionDescription
	pendingRemoteDescription *SessionDescription
	signalingState           SignalingState
	iceConnectionState       atomic.Value
	connectionState          atomic.Value

	idpLoginURL *string

	isClosed                                *atomic.Bool
	isGracefullyClosingOrClosed             bool
	isCloseDone                             chan struct{}
	isGracefulCloseDone                     chan struct{}
	isNegotiationNeeded                     *atomic.Bool
	updateNegotiationNeededFlagOnEmptyChain *atomic.Bool

	lastOffer  string
	lastAnswer string

	canTrickleICECandidates ICETrickleCapability

	greaterMid int

	rtpTransceivers        []*RTPTransceiver
	nonMediaBandwidthProbe atomic.Value

	onSignalingStateChangeHandler     func(SignalingState)
	onICEConnectionStateChangeHandler atomic.Value
	onConnectionStateChangeHandler    atomic.Value
	onTrackHandler                    func(*TrackRemote, *RTPReceiver)
	onDataChannelHandler              func(*DataChannel)
	onNegotiationNeededHandler        atomic.Value

	iceGatherer   *ICEGatherer
	iceTransport  *ICETransport
	dtlsTransport *DTLSTransport
	sctpTransport *SCTPTransport

	api *API
	log logging.LeveledLogger

	interceptorRTCPWriter interceptor.RTCPWriter
	statsGetter           stats.Getter
}

func NewPeerConnection(configuration Configuration) (*PeerConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *API) NewPeerConnection(configuration Configuration) (*PeerConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *PeerConnection) initConfiguration(configuration Configuration) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

func (pc *PeerConnection) OnSignalingStateChange(f func(SignalingState)) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) onSignalingStateChange(newState SignalingState) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) OnDataChannel(f func(*DataChannel)) { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) defaultOnDataChannelHandler(d *DataChannel) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) OnNegotiationNeeded(f func()) { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) onNegotiationNeeded() { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) negotiationNeededOp() { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) checkNegotiationNeeded() bool {
	_ = "STUB: not implemented" //nolint:gocognit,cyclop
	return false
}

func (pc *PeerConnection) OnICECandidate(f func(*ICECandidate)) { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) OnICEGatheringStateChange(f func(ICEGatheringState)) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) OnTrack(f func(*TrackRemote, *RTPReceiver)) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) onTrack(t *TrackRemote, r *RTPReceiver) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) OnICEConnectionStateChange(f func(ICEConnectionState)) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) onICEConnectionStateChange(cs ICEConnectionState) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) OnConnectionStateChange(f func(PeerConnectionState)) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) onConnectionStateChange(cs PeerConnectionState) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) SetConfiguration(configuration Configuration) error {
	_ = "STUB: not implemented" //nolint:gocognit,cyclop
	return nil
}

func (pc *PeerConnection) GetConfiguration() Configuration {
	_ = "STUB: not implemented"
	return *new(Configuration)
}

func (pc *PeerConnection) ID() string { _ = "STUB: not implemented"; return "" }

func (pc *PeerConnection) hasLocalDescriptionChanged(desc *SessionDescription) bool {
	_ = "STUB: not implemented"
	return false
}

//nolint:gocognit,cyclop
func (pc *PeerConnection) CreateOffer(options *OfferOptions) (SessionDescription, error) {
	_ = "STUB: not implemented"
	return *new(SessionDescription), nil
}

//nolint:nestif

func (pc *PeerConnection) createICEGatherer() (*ICEGatherer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:cyclop
func (pc *PeerConnection) updateConnectionState(
	iceConnectionState ICEConnectionState,
	dtlsTransportState DTLSTransportState,
) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) createICETransport() *ICETransport { _ = "STUB: not implemented"; return nil }

//nolint:cyclop
func (pc *PeerConnection) CreateAnswer(options *AnswerOptions) (SessionDescription, error) {
	_ = "STUB: not implemented"
	return *new(SessionDescription), nil
}

//nolint:gocognit,cyclop
func (pc *PeerConnection) setDescription(sd *SessionDescription, op stateChangeOp) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:cyclop
func (pc *PeerConnection) SetLocalDescription(desc SessionDescription) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) LocalDescription() *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gocognit,gocyclo,cyclop,maintidx
func (pc *PeerConnection) SetRemoteDescription(desc SessionDescription) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:nestif

func (pc *PeerConnection) configureReceiver(incoming trackDetails, receiver *RTPReceiver) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) startReceiver(incoming trackDetails, receiver *RTPReceiver) {
	_ = "STUB: not implemented"
	return
}

//nolint:cyclop
func setRTPTransceiverCurrentDirection(
	answer *SessionDescription,
	currentTransceivers []*RTPTransceiver,
	weOffer bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func runIfNewReceiver(
	incomingTrack trackDetails,
	transceivers []*RTPTransceiver,
	callbackFunc func(incomingTrack trackDetails, receiver *RTPReceiver),
) bool {
	_ = "STUB: not implemented"
	return false
}

//nolint:gocognit,cyclop
func (pc *PeerConnection) configureRTPReceivers(
	isRenegotiation bool,
	remoteDesc *SessionDescription,
	currentTransceivers []*RTPTransceiver,
) {
	_ = "STUB: not implemented"
	return
}

//nolint:nestif

func (pc *PeerConnection) startRTPReceivers(remoteDesc *SessionDescription, currentTransceivers []*RTPTransceiver) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) startRTPSenders(currentTransceivers []*RTPTransceiver) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) startSCTP(maxMessageSize uint32, remoteSctpInit []byte) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) handleUndeclaredSSRC(
	ssrc SSRC,
	mediaSection *sdp.MediaDescription,
) (handled bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pc *PeerConnection) findMediaSectionByPayloadType(
	payloadType PayloadType,
	remoteDescription *SessionDescription,
) (selectedMediaSection *sdp.MediaDescription, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pc *PeerConnection) handleNonMediaBandwidthProbe() { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) handleIncomingSSRC(rtpStream *srtp.ReadStreamSRTP, ssrc SSRC) error {
	_ = "STUB: not implemented" //nolint:gocyclo,gocognit,cyclop,lll
	return nil
}

//nolint:gosec // G115
//nolint:gosec // G115
//nolint:gosec // G115

//nolint:gosec // G115
//nolint:gosec // G115
//nolint:gosec // G115

func (pc *PeerConnection) undeclaredMediaProcessor() { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) undeclaredRTPMediaProcessor() {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

func (pc *PeerConnection) undeclaredRTCPMediaProcessor() { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) RemoteDescription() *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) AddICECandidate(candidate ICECandidateInit) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) descriptionContainsUfrag(sdp *sdp.SessionDescription, matchUfrag string) bool {
	_ = "STUB: not implemented"
	return false
}

func (pc *PeerConnection) ICEConnectionState() ICEConnectionState {
	_ = "STUB: not implemented"
	return *new(ICEConnectionState)
}

func (pc *PeerConnection) GetSenders() (result []*RTPSender) { _ = "STUB: not implemented"; return nil }

func (pc *PeerConnection) GetReceivers() (receivers []*RTPReceiver) {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) GetTransceivers() []*RTPTransceiver {
	_ = "STUB: not implemented"
	return nil
}

//nolint:cyclop
func (pc *PeerConnection) AddTrack(track TrackLocal) (*RTPSender, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *PeerConnection) RemoveTrack(sender *RTPSender) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:cyclop
func (pc *PeerConnection) newTransceiverFromTrack(
	direction RTPTransceiverDirection,
	track TrackLocal,
	init ...RTPTransceiverInit,
) (t *RTPTransceiver, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:cyclop
func (pc *PeerConnection) AddTransceiverFromKind(
	kind RTPCodecType,
	init ...RTPTransceiverInit,
) (t *RTPTransceiver, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *PeerConnection) AddTransceiverFromTrack(
	track TrackLocal,
	init ...RTPTransceiverInit,
) (t *RTPTransceiver, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:cyclop
func (pc *PeerConnection) CreateDataChannel(label string, options *DataChannelInit) (*DataChannel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nestif

func (pc *PeerConnection) SetIdentityProvider(string) error { _ = "STUB: not implemented"; return nil }

func (pc *PeerConnection) WriteRTCP(pkts []rtcp.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) writeRTCP(pkts []rtcp.Packet, _ interceptor.Attributes) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pc *PeerConnection) Close() error { _ = "STUB: not implemented"; return nil }

func (pc *PeerConnection) GracefulClose() error { _ = "STUB: not implemented"; return nil }

func (pc *PeerConnection) close(shouldGracefullyClose bool) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

func (pc *PeerConnection) addRTPTransceiver(t *RTPTransceiver) { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) CurrentLocalDescription() *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) PendingLocalDescription() *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) CurrentRemoteDescription() *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) PendingRemoteDescription() *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) CanTrickleICECandidates() ICETrickleCapability {
	_ = "STUB: not implemented"
	return *new(ICETrickleCapability)
}

func (pc *PeerConnection) SignalingState() SignalingState {
	_ = "STUB: not implemented"
	return *new(SignalingState)
}

func (pc *PeerConnection) ICEGatheringState() ICEGatheringState {
	_ = "STUB: not implemented"
	return *new(ICEGatheringState)
}

func (pc *PeerConnection) ConnectionState() PeerConnectionState {
	_ = "STUB: not implemented"
	return *new(PeerConnectionState)
}

func (pc *PeerConnection) GetStats() StatsReport {
	_ = "STUB: not implemented"
	return *new(StatsReport)
}

func (pc *PeerConnection) startTransports(
	iceRole ICERole,
	dtlsRole DTLSRole,
	remoteIsLite bool,
	remoteUfrag, remotePwd, fingerprint, fingerprintHash string,
) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) startRTP(
	isRenegotiation bool,
	remoteDesc *SessionDescription,
	currentTransceivers []*RTPTransceiver,
) {
	_ = "STUB: not implemented"
	return
}

//nolint:cyclop
func (pc *PeerConnection) generateUnmatchedSDP(
	transceivers []*RTPTransceiver,
	useIdentity bool,
) (*sdp.SessionDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nestif

//nolint:gocognit,gocyclo,cyclop,maintidx
func (pc *PeerConnection) generateMatchedSDP(
	transceivers []*RTPTransceiver,
	useIdentity, includeUnmatched bool,
	connectionRole sdp.ConnectionRole,
	ignoreRidPauseForRecv bool,
) (*sdp.SessionDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nestif

func (pc *PeerConnection) setGatherCompleteHandler(handler func()) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) SCTP() *SCTPTransport { _ = "STUB: not implemented"; return nil }
