package webrtc

import (
	"math"

	"github.com/pion/dtls/v3"
)

const (
	receiveMTU = 1500

	simulcastProbeCount = 10

	simulcastMaxProbeRoutines = 25

	defaultMaxSCTPMessageSize = 1073741823

	sctpMaxMessageSizeUnsetValue = math.MaxUint16

	mediaSectionApplication = "application"

	sdpAttributeRid = "rid"

	sdpAttributeSimulcast = "simulcast"

	outboundMTU = 1200

	sctpOutboundMTU = 1191

	rtpPayloadTypeBitmask = 0x7F

	incomingUnhandledRTPSsrc = "Incoming unhandled RTP ssrc(%d), OnTrack will not be fired. %v"

	useReadSimulcast = "Use ReadSimulcast(rid) instead of Read() when multiple tracks are present"

	generatedCertificateOrigin = "WebRTC"

	AttributeRtxPayloadType = "rtx_payload_type"

	AttributeRtxSsrc = "rtx_ssrc"

	AttributeRtxSequenceNumber = "rtx_sequence_number"
)

func defaultSrtpProtectionProfiles() []dtls.SRTPProtectionProfile {
	_ = "STUB: not implemented"
	return nil
}
