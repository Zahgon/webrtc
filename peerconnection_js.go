//go:build js && wasm
// +build js,wasm

package webrtc

import (
	"syscall/js"
)

type PeerConnection struct {
	underlying js.Value

	onSignalingStateChangeHandler     *js.Func
	onDataChannelHandler              *js.Func
	onNegotiationNeededHandler        *js.Func
	onConnectionStateChangeHandler    *js.Func
	onICEConnectionStateChangeHandler *js.Func
	onICECandidateHandler             *js.Func
	onICEGatheringStateChangeHandler  *js.Func

	onGatherCompleteHandler func()

	api *API
}

func NewPeerConnection(configuration Configuration) (*PeerConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *API) NewPeerConnection(configuration Configuration) (_ *PeerConnection, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *PeerConnection) JSValue() js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func (pc *PeerConnection) OnSignalingStateChange(f func(SignalingState)) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) OnDataChannel(f func(*DataChannel)) { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) OnNegotiationNeeded(f func()) { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) OnICEConnectionStateChange(f func(ICEConnectionState)) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) OnConnectionStateChange(f func(PeerConnectionState)) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) checkConfiguration(configuration Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) SetConfiguration(configuration Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) GetConfiguration() Configuration {
	_ = "STUB: not implemented"
	return *new(Configuration)
}

func (pc *PeerConnection) CreateOffer(options *OfferOptions) (_ SessionDescription, err error) {
	_ = "STUB: not implemented"
	return *new(SessionDescription), nil
}

func (pc *PeerConnection) CreateAnswer(options *AnswerOptions) (_ SessionDescription, err error) {
	_ = "STUB: not implemented"
	return *new(SessionDescription), nil
}

func (pc *PeerConnection) SetLocalDescription(desc SessionDescription) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) LocalDescription() *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) SetRemoteDescription(desc SessionDescription) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) RemoteDescription() *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) AddICECandidate(candidate ICECandidateInit) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) ICEConnectionState() ICEConnectionState {
	_ = "STUB: not implemented"
	return *new(ICEConnectionState)
}

func (pc *PeerConnection) OnICECandidate(f func(candidate *ICECandidate)) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) OnICEGatheringStateChange(f func()) { _ = "STUB: not implemented"; return }

func (pc *PeerConnection) CreateDataChannel(label string, options *DataChannelInit) (_ *DataChannel, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *PeerConnection) SetIdentityProvider(provider string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) Close() (err error) { _ = "STUB: not implemented"; return nil }

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

func (pc *PeerConnection) setGatherCompleteHandler(handler func()) {
	_ = "STUB: not implemented"
	return
}

func (pc *PeerConnection) AddTransceiverFromKind(kind RTPCodecType, init ...RTPTransceiverInit) (transceiver *RTPTransceiver, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *PeerConnection) GetTransceivers() (transceivers []*RTPTransceiver) {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PeerConnection) SCTP() *SCTPTransport { _ = "STUB: not implemented"; return nil }

func configurationToValue(configuration Configuration) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func iceServersToValue(iceServers []ICEServer) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func oauthCredentialToValue(o OAuthCredential) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func iceServerToValue(server ICEServer) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func valueToConfiguration(configValue js.Value) Configuration {
	_ = "STUB: not implemented"
	return *new(Configuration)
}

func valueToICEServers(iceServersValue js.Value) []ICEServer { _ = "STUB: not implemented"; return nil }

func valueToICECredential(iceCredentialValue js.Value) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func valueToICEServer(iceServerValue js.Value) ICEServer {
	_ = "STUB: not implemented"
	return *new(ICEServer)
}

func valueToICECandidate(val js.Value) *ICECandidate { _ = "STUB: not implemented"; return nil }

func stringToComponentIDOrZero(val string) uint16 { _ = "STUB: not implemented"; return 0 }

func sessionDescriptionToValue(desc *SessionDescription) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func valueToSessionDescription(descValue js.Value) *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

func offerOptionsToValue(offerOptions *OfferOptions) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func answerOptionsToValue(answerOptions *AnswerOptions) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func iceCandidateInitToValue(candidate ICECandidateInit) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func dataChannelInitToValue(options *DataChannelInit) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func rtpTransceiverInitInitToValue(init RTPTransceiverInit) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}
