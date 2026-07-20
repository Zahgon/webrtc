//go:build !js

package webrtc

import (
	"sync"
	"sync/atomic"

	"github.com/pion/sdp/v3"
)

type RTPTransceiver struct {
	mid                    atomic.Value
	sender                 atomic.Value
	receiver               atomic.Value
	direction              atomic.Value
	currentDirection       atomic.Value
	currentRemoteDirection atomic.Value

	codecs []RTPCodecParameters

	kind RTPCodecType

	api *API
	mu  sync.RWMutex
}

func newRTPTransceiver(
	receiver *RTPReceiver,
	sender *RTPSender,
	direction RTPTransceiverDirection,
	kind RTPCodecType,
	api *API,
) *RTPTransceiver {
	_ = "STUB: not implemented"
	return nil
}

func (t *RTPTransceiver) SetCodecPreferences(codecs []RTPCodecParameters) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *RTPTransceiver) getCodecs() []RTPCodecParameters { _ = "STUB: not implemented"; return nil }

func (t *RTPTransceiver) setCodecPreferencesFromRemoteDescription(media *sdp.MediaDescription) {
	_ = "STUB: not implemented" //nolint:cyclop
	return
}

func (t *RTPTransceiver) Sender() *RTPSender { _ = "STUB: not implemented"; return nil }

func (t *RTPTransceiver) SetSender(s *RTPSender, track TrackLocal) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *RTPTransceiver) setSender(s *RTPSender) { _ = "STUB: not implemented"; return }

func (t *RTPTransceiver) Receiver() *RTPReceiver { _ = "STUB: not implemented"; return nil }

func (t *RTPTransceiver) SetMid(mid string) error { _ = "STUB: not implemented"; return nil }

func (t *RTPTransceiver) Mid() string { _ = "STUB: not implemented"; return "" }

func (t *RTPTransceiver) Kind() RTPCodecType { _ = "STUB: not implemented"; return *new(RTPCodecType) }

func (t *RTPTransceiver) Direction() RTPTransceiverDirection {
	_ = "STUB: not implemented"
	return *new(RTPTransceiverDirection)
}

func (t *RTPTransceiver) Stop() error { _ = "STUB: not implemented"; return nil }

func (t *RTPTransceiver) setReceiver(r *RTPReceiver) { _ = "STUB: not implemented"; return }

func (t *RTPTransceiver) setDirection(d RTPTransceiverDirection) { _ = "STUB: not implemented"; return }

func (t *RTPTransceiver) setCurrentDirection(d RTPTransceiverDirection) {
	_ = "STUB: not implemented"
	return
}

func (t *RTPTransceiver) getCurrentDirection() RTPTransceiverDirection {
	_ = "STUB: not implemented"
	return *new(RTPTransceiverDirection)
}

func (t *RTPTransceiver) setCurrentRemoteDirection(d RTPTransceiverDirection) {
	_ = "STUB: not implemented"
	return
}

func (t *RTPTransceiver) getCurrentRemoteDirection() RTPTransceiverDirection {
	_ = "STUB: not implemented"
	return *new(RTPTransceiverDirection)
}

func (t *RTPTransceiver) setSendingTrack(track TrackLocal) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

func (t *RTPTransceiver) isSendAllowed(kind RTPCodecType) bool {
	_ = "STUB: not implemented"
	return false
}

func findByMid(mid string, localTransceivers []*RTPTransceiver) (*RTPTransceiver, []*RTPTransceiver) {
	_ = "STUB: not implemented"
	return nil, nil
}

func satisfyTypeAndDirection(
	remoteKind RTPCodecType,
	remoteDirection RTPTransceiverDirection,
	localTransceivers []*RTPTransceiver,
) (*RTPTransceiver, []*RTPTransceiver) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleUnknownRTPPacket(
	buf []byte,
	midExtensionID,
	streamIDExtensionID,
	repairStreamIDExtensionID uint8,
) (mid, rid, rsid string, paddingOnly bool, err error) {
	_ = "STUB: not implemented"
	return "", "", "", false, nil
}
