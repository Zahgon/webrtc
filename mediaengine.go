//go:build !js

package webrtc

import (
	"sync"

	"github.com/pion/rtp"
	"github.com/pion/sdp/v3"
)

type mediaEngineHeaderExtension struct {
	uri              string
	isAudio, isVideo bool

	allowedDirections []RTPTransceiverDirection
}

type MediaEngine struct {
	negotiatedVideo, negotiatedAudio bool
	negotiateMultiCodecs             bool

	videoCodecs, audioCodecs                     []RTPCodecParameters
	negotiatedVideoCodecs, negotiatedAudioCodecs []RTPCodecParameters

	headerExtensions           []mediaEngineHeaderExtension
	negotiatedHeaderExtensions map[int]mediaEngineHeaderExtension

	mu sync.RWMutex
}

func (m *MediaEngine) setMultiCodecNegotiation(negotiateMultiCodecs bool) {
	_ = "STUB: not implemented"
	return
}

func (m *MediaEngine) multiCodecNegotiation() bool { _ = "STUB: not implemented"; return false }

func (m *MediaEngine) RegisterDefaultCodecs() error { _ = "STUB: not implemented"; return nil }

func (m *MediaEngine) addCodec(codecs []RTPCodecParameters, codec RTPCodecParameters) ([]RTPCodecParameters, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MediaEngine) RegisterCodec(codec RTPCodecParameters, typ RTPCodecType) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:cyclop
func (m *MediaEngine) RegisterHeaderExtension(
	extension RTPHeaderExtensionCapability,
	typ RTPCodecType,
	allowedDirections ...RTPTransceiverDirection,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MediaEngine) RegisterFeedback(feedback RTCPFeedback, typ RTPCodecType) {
	_ = "STUB: not implemented"
	return
}

func (m *MediaEngine) getHeaderExtensionID(extension RTPHeaderExtensionCapability) (
	val int,
	audioNegotiated, videoNegotiated bool,
) {
	_ = "STUB: not implemented"
	return 0, false, false
}

func (m *MediaEngine) copy() *MediaEngine { _ = "STUB: not implemented"; return nil }

func findCodecByPayload(codecs []RTPCodecParameters, payloadType PayloadType) *RTPCodecParameters {
	_ = "STUB: not implemented"
	return nil
}

func (m *MediaEngine) getCodecByPayload(payloadType PayloadType) (RTPCodecParameters, RTPCodecType, error) {
	_ = "STUB: not implemented"
	return *new(RTPCodecParameters), *new(RTPCodecType), nil
}

func (m *MediaEngine) collectStats(collector *statsReportCollector) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec // G115

//nolint:cyclop
func (m *MediaEngine) matchRemoteCodec(
	remoteCodec RTPCodecParameters,
	typ RTPCodecType,
	exactMatches, partialMatches []RTPCodecParameters,
) (RTPCodecParameters, codecMatchType, error) {
	_ = "STUB: not implemented"
	return *new(RTPCodecParameters), *new(codecMatchType), nil
}

//nolint:nestif

func (m *MediaEngine) updateHeaderExtensionFromMediaSection(media *sdp.MediaDescription) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MediaEngine) updateHeaderExtension(id int, extension string, typ RTPCodecType) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MediaEngine) pushCodecs(codecs []RTPCodecParameters, typ RTPCodecType) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MediaEngine) updateFromRemoteDescription(desc sdp.SessionDescription) error {
	_ = "STUB: not implemented" //nolint:cyclop,gocognit
	return nil
}

func (m *MediaEngine) getCodecsByKind(typ RTPCodecType) []RTPCodecParameters {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gocognit,cyclop
func (m *MediaEngine) getRTPParametersByKind(typ RTPCodecType, directions []RTPTransceiverDirection) RTPParameters {
	_ = "STUB: not implemented"
	return *new(RTPParameters)
}

//nolint:nestif

func (m *MediaEngine) getRTPParametersByPayloadType(payloadType PayloadType) (RTPParameters, error) {
	_ = "STUB: not implemented"
	return *new(RTPParameters), nil
}

func payloaderForCodec(codec RTPCodecCapability) (rtp.Payloader, error) {
	_ = "STUB: not implemented"
	return *new(rtp.Payloader), nil
}

func (m *MediaEngine) isRTXEnabled(typ RTPCodecType, directions []RTPTransceiverDirection) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *MediaEngine) isFECEnabled(typ RTPCodecType, directions []RTPTransceiverDirection) bool {
	_ = "STUB: not implemented"
	return false
}
