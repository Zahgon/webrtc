package webrtc

import (
	"sync"
	"time"

	"github.com/pion/ice/v4"
)

type Stats interface {
	statsMarker()
}

func UnmarshalStatsJSON(b []byte) (Stats, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return *new(Stats), nil
}

type StatsType string

const (
	StatsTypeCodec StatsType = "codec"

	StatsTypeInboundRTP StatsType = "inbound-rtp"

	StatsTypeOutboundRTP StatsType = "outbound-rtp"

	StatsTypeRemoteInboundRTP StatsType = "remote-inbound-rtp"

	StatsTypeRemoteOutboundRTP StatsType = "remote-outbound-rtp"

	StatsTypeCSRC StatsType = "csrc"

	StatsTypeMediaSource = "media-source"

	StatsTypeMediaPlayout StatsType = "media-playout"

	StatsTypePeerConnection StatsType = "peer-connection"

	StatsTypeDataChannel StatsType = "data-channel"

	StatsTypeStream StatsType = "stream"

	StatsTypeTrack StatsType = "track"

	StatsTypeSender StatsType = "sender"

	StatsTypeReceiver StatsType = "receiver"

	StatsTypeTransport StatsType = "transport"

	StatsTypeCandidatePair StatsType = "candidate-pair"

	StatsTypeLocalCandidate StatsType = "local-candidate"

	StatsTypeRemoteCandidate StatsType = "remote-candidate"

	StatsTypeCertificate StatsType = "certificate"

	StatsTypeSCTPTransport StatsType = "sctp-transport"
)

type MediaKind string

const (
	MediaKindAudio MediaKind = "audio"

	MediaKindVideo MediaKind = "video"
)

type StatsTimestamp float64

func (s StatsTimestamp) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func statsTimestampFrom(t time.Time) StatsTimestamp {
	_ = "STUB: not implemented"
	return *new(StatsTimestamp)
}

func statsTimestampNow() StatsTimestamp { _ = "STUB: not implemented"; return *new(StatsTimestamp) }

type StatsReport map[string]Stats

type statsReportCollector struct {
	collectingGroup sync.WaitGroup
	report          StatsReport
	mux             sync.Mutex
}

type SCTPTransportPartialReliabilityMode string

const (
	SCTPTransportPartialReliabilityModeNone        SCTPTransportPartialReliabilityMode = "none"
	SCTPTransportPartialReliabilityModeForwardTSN  SCTPTransportPartialReliabilityMode = "forward-tsn"
	SCTPTransportPartialReliabilityModeIForwardTSN SCTPTransportPartialReliabilityMode = "i-forward-tsn"
)

type SCTPTransportMetadata struct {
	MessageInterleavingEnabled bool `json:"messageInterleavingEnabled"`

	PartialReliabilityMode SCTPTransportPartialReliabilityMode `json:"partialReliabilityMode"`

	ZeroChecksumSendingEnabled bool `json:"zeroChecksumSendingEnabled"`

	ZeroChecksumReceivingEnabled bool `json:"zeroChecksumReceivingEnabled"`
}

func newStatsReportCollector() *statsReportCollector { _ = "STUB: not implemented"; return nil }

func (src *statsReportCollector) Collecting() { _ = "STUB: not implemented"; return }

func (src *statsReportCollector) Collect(id string, stats Stats) { _ = "STUB: not implemented"; return }

func (src *statsReportCollector) Done() { _ = "STUB: not implemented"; return }

func (src *statsReportCollector) Ready() StatsReport {
	_ = "STUB: not implemented"
	return *new(StatsReport)
}

type CodecType string

const (
	CodecTypeEncode CodecType = "encode"

	CodecTypeDecode CodecType = "decode"
)

type CodecStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	PayloadType PayloadType `json:"payloadType"`

	CodecType CodecType `json:"codecType"`

	TransportID string `json:"transportId"`

	MimeType string `json:"mimeType"`

	ClockRate uint32 `json:"clockRate"`

	Channels uint8 `json:"channels"`

	SDPFmtpLine string `json:"sdpFmtpLine"`

	Implementation string `json:"implementation"`
}

func (s CodecStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalCodecStats(b []byte) (CodecStats, error) {
	_ = "STUB: not implemented"
	return *new(CodecStats), nil
}

type InboundRTPStreamStats struct {
	Mid string `json:"mid"`

	Rid string `json:"rid,omitempty"`

	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	SSRC SSRC `json:"ssrc"`

	Kind string `json:"kind"`

	TransportID string `json:"transportId"`

	CodecID string `json:"codecId"`

	FIRCount uint32 `json:"firCount"`

	PLICount uint32 `json:"pliCount"`

	TotalProcessingDelay float64 `json:"totalProcessingDelay"`

	NACKCount uint32 `json:"nackCount"`

	JitterBufferDelay float64 `json:"jitterBufferDelay"`

	JitterBufferTargetDelay float64 `json:"jitterBufferTargetDelay"`

	JitterBufferEmittedCount uint64 `json:"jitterBufferEmittedCount"`

	JitterBufferMinimumDelay float64 `json:"jitterBufferMinimumDelay"`

	TotalSamplesReceived uint64 `json:"totalSamplesReceived"`

	ConcealedSamples uint64 `json:"concealedSamples"`

	SilentConcealedSamples uint64 `json:"silentConcealedSamples"`

	ConcealmentEvents uint64 `json:"concealmentEvents"`

	InsertedSamplesForDeceleration uint64 `json:"insertedSamplesForDeceleration"`

	RemovedSamplesForAcceleration uint64 `json:"removedSamplesForAcceleration"`

	AudioLevel float64 `json:"audioLevel"`

	TotalAudioEnergy float64 `json:"totalAudioEnergy"`

	TotalSamplesDuration float64 `json:"totalSamplesDuration"`

	SLICount uint32 `json:"sliCount"`

	QPSum uint64 `json:"qpSum"`

	TotalDecodeTime float64 `json:"totalDecodeTime"`

	TotalInterFrameDelay float64 `json:"totalInterFrameDelay"`

	TotalSquaredInterFrameDelay float64 `json:"totalSquaredInterFrameDelay"`

	PacketsReceived uint32 `json:"packetsReceived"`

	PacketsLost int32 `json:"packetsLost"`

	Jitter float64 `json:"jitter"`

	PacketsDiscarded uint32 `json:"packetsDiscarded"`

	PacketsRepaired uint32 `json:"packetsRepaired"`

	BurstPacketsLost uint32 `json:"burstPacketsLost"`

	BurstPacketsDiscarded uint32 `json:"burstPacketsDiscarded"`

	BurstLossCount uint32 `json:"burstLossCount"`

	BurstDiscardCount uint32 `json:"burstDiscardCount"`

	BurstLossRate float64 `json:"burstLossRate"`

	BurstDiscardRate float64 `json:"burstDiscardRate"`

	GapLossRate float64 `json:"gapLossRate"`

	GapDiscardRate float64 `json:"gapDiscardRate"`

	TrackID string `json:"trackId"`

	ReceiverID string `json:"receiverId"`

	RemoteID string `json:"remoteId"`

	FramesDecoded uint32 `json:"framesDecoded"`

	KeyFramesDecoded uint32 `json:"keyFramesDecoded"`

	FramesRendered uint32 `json:"framesRendered"`

	FramesDropped uint32 `json:"framesDropped"`

	FrameWidth uint32 `json:"frameWidth"`

	FrameHeight uint32 `json:"frameHeight"`

	LastPacketReceivedTimestamp StatsTimestamp `json:"lastPacketReceivedTimestamp"`

	HeaderBytesReceived uint64 `json:"headerBytesReceived"`

	AverageRTCPInterval float64 `json:"averageRtcpInterval"`

	FECPacketsReceived uint32 `json:"fecPacketsReceived"`

	FECPacketsDiscarded uint64 `json:"fecPacketsDiscarded"`

	BytesReceived uint64 `json:"bytesReceived"`

	FramesReceived uint32 `json:"framesReceived"`

	PacketsFailedDecryption uint32 `json:"packetsFailedDecryption"`

	PacketsDuplicated uint32 `json:"packetsDuplicated"`

	PerDSCPPacketsReceived map[string]uint32 `json:"perDscpPacketsReceived"`

	DecoderImplementation string `json:"decoderImplementation"`

	PauseCount uint32 `json:"pauseCount"`

	TotalPausesDuration float64 `json:"totalPausesDuration"`

	FreezeCount uint32 `json:"freezeCount"`

	TotalFreezesDuration float64 `json:"totalFreezesDuration"`

	PowerEfficientDecoder bool `json:"powerEfficientDecoder"`
}

func (s InboundRTPStreamStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalInboundRTPStreamStats(b []byte) (InboundRTPStreamStats, error) {
	_ = "STUB: not implemented"
	return *new(InboundRTPStreamStats), nil
}

type QualityLimitationReason string

const (
	QualityLimitationReasonNone QualityLimitationReason = "none"

	QualityLimitationReasonCPU QualityLimitationReason = "cpu"

	QualityLimitationReasonBandwidth QualityLimitationReason = "bandwidth"

	QualityLimitationReasonOther QualityLimitationReason = "other"
)

type OutboundRTPStreamStats struct {
	Mid string `json:"mid"`

	Rid string `json:"rid"`

	MediaSourceID string `json:"mediaSourceId"`

	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	SSRC SSRC `json:"ssrc"`

	Kind string `json:"kind"`

	TransportID string `json:"transportId"`

	CodecID string `json:"codecId"`

	HeaderBytesSent uint64 `json:"headerBytesSent"`

	RetransmittedPacketsSent uint64 `json:"retransmittedPacketsSent"`

	RetransmittedBytesSent uint64 `json:"retransmittedBytesSent"`

	FIRCount uint32 `json:"firCount"`

	PLICount uint32 `json:"pliCount"`

	NACKCount uint32 `json:"nackCount"`

	SLICount uint32 `json:"sliCount"`

	QPSum uint64 `json:"qpSum"`

	PacketsSent uint32 `json:"packetsSent"`

	PacketsDiscardedOnSend uint32 `json:"packetsDiscardedOnSend"`

	FECPacketsSent uint32 `json:"fecPacketsSent"`

	BytesSent uint64 `json:"bytesSent"`

	BytesDiscardedOnSend uint64 `json:"bytesDiscardedOnSend"`

	TrackID string `json:"trackId"`

	SenderID string `json:"senderId"`

	RemoteID string `json:"remoteId"`

	LastPacketSentTimestamp StatsTimestamp `json:"lastPacketSentTimestamp"`

	TargetBitrate float64 `json:"targetBitrate"`

	TotalEncodedBytesTarget uint64 `json:"totalEncodedBytesTarget"`

	FrameWidth uint32 `json:"frameWidth"`

	FrameHeight uint32 `json:"frameHeight"`

	FramesPerSecond float64 `json:"framesPerSecond"`

	FramesSent uint32 `json:"framesSent"`

	HugeFramesSent uint32 `json:"hugeFramesSent"`

	FramesEncoded uint32 `json:"framesEncoded"`

	KeyFramesEncoded uint32 `json:"keyFramesEncoded"`

	TotalEncodeTime float64 `json:"totalEncodeTime"`

	TotalPacketSendDelay float64 `json:"totalPacketSendDelay"`

	AverageRTCPInterval float64 `json:"averageRtcpInterval"`

	QualityLimitationReason QualityLimitationReason `json:"qualityLimitationReason"`

	QualityLimitationDurations map[string]float64 `json:"qualityLimitationDurations"`

	QualityLimitationResolutionChanges uint32 `json:"qualityLimitationResolutionChanges"`

	PerDSCPPacketsSent map[string]uint32 `json:"perDscpPacketsSent"`

	Active bool `json:"active"`

	EncoderImplementation string `json:"encoderImplementation"`

	PowerEfficientEncoder bool `json:"powerEfficientEncoder"`

	ScalabilityMode string `json:"scalabilityMode"`
}

func (s OutboundRTPStreamStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalOutboundRTPStreamStats(b []byte) (OutboundRTPStreamStats, error) {
	_ = "STUB: not implemented"
	return *new(OutboundRTPStreamStats), nil
}

type RemoteInboundRTPStreamStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	SSRC SSRC `json:"ssrc"`

	Kind string `json:"kind"`

	TransportID string `json:"transportId"`

	CodecID string `json:"codecId"`

	FIRCount uint32 `json:"firCount"`

	PLICount uint32 `json:"pliCount"`

	NACKCount uint32 `json:"nackCount"`

	SLICount uint32 `json:"sliCount"`

	QPSum uint64 `json:"qpSum"`

	PacketsReceived uint32 `json:"packetsReceived"`

	PacketsLost int32 `json:"packetsLost"`

	Jitter float64 `json:"jitter"`

	PacketsDiscarded uint32 `json:"packetsDiscarded"`

	PacketsRepaired uint32 `json:"packetsRepaired"`

	BurstPacketsLost uint32 `json:"burstPacketsLost"`

	BurstPacketsDiscarded uint32 `json:"burstPacketsDiscarded"`

	BurstLossCount uint32 `json:"burstLossCount"`

	BurstDiscardCount uint32 `json:"burstDiscardCount"`

	BurstLossRate float64 `json:"burstLossRate"`

	BurstDiscardRate float64 `json:"burstDiscardRate"`

	GapLossRate float64 `json:"gapLossRate"`

	GapDiscardRate float64 `json:"gapDiscardRate"`

	LocalID string `json:"localId"`

	RoundTripTime float64 `json:"roundTripTime"`

	TotalRoundTripTime float64 `json:"totalRoundTripTime"`

	FractionLost float64 `json:"fractionLost"`

	RoundTripTimeMeasurements uint64 `json:"roundTripTimeMeasurements"`
}

func (s RemoteInboundRTPStreamStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalRemoteInboundRTPStreamStats(b []byte) (RemoteInboundRTPStreamStats, error) {
	_ = "STUB: not implemented"
	return *new(RemoteInboundRTPStreamStats), nil
}

type RemoteOutboundRTPStreamStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	SSRC SSRC `json:"ssrc"`

	Kind string `json:"kind"`

	TransportID string `json:"transportId"`

	CodecID string `json:"codecId"`

	FIRCount uint32 `json:"firCount"`

	PLICount uint32 `json:"pliCount"`

	NACKCount uint32 `json:"nackCount"`

	SLICount uint32 `json:"sliCount"`

	QPSum uint64 `json:"qpSum"`

	PacketsSent uint32 `json:"packetsSent"`

	PacketsDiscardedOnSend uint32 `json:"packetsDiscardedOnSend"`

	FECPacketsSent uint32 `json:"fecPacketsSent"`

	BytesSent uint64 `json:"bytesSent"`

	BytesDiscardedOnSend uint64 `json:"bytesDiscardedOnSend"`

	LocalID string `json:"localId"`

	RemoteTimestamp StatsTimestamp `json:"remoteTimestamp"`

	ReportsSent uint64 `json:"reportsSent"`

	RoundTripTime float64 `json:"roundTripTime"`

	TotalRoundTripTime float64 `json:"totalRoundTripTime"`

	RoundTripTimeMeasurements uint64 `json:"roundTripTimeMeasurements"`
}

func (s RemoteOutboundRTPStreamStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalRemoteOutboundRTPStreamStats(b []byte) (RemoteOutboundRTPStreamStats, error) {
	_ = "STUB: not implemented"
	return *new(RemoteOutboundRTPStreamStats), nil
}

type RTPContributingSourceStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	ContributorSSRC SSRC `json:"contributorSsrc"`

	InboundRTPStreamID string `json:"inboundRtpStreamId"`

	PacketsContributedTo uint32 `json:"packetsContributedTo"`

	AudioLevel float64 `json:"audioLevel"`
}

func (s RTPContributingSourceStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalCSRCStats(b []byte) (RTPContributingSourceStats, error) {
	_ = "STUB: not implemented"
	return *new(RTPContributingSourceStats), nil
}

type AudioSourceStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	TrackIdentifier string `json:"trackIdentifier"`

	Kind string `json:"kind"`

	AudioLevel float64 `json:"audioLevel"`

	TotalAudioEnergy float64 `json:"totalAudioEnergy"`

	TotalSamplesDuration float64 `json:"totalSamplesDuration"`

	EchoReturnLoss float64 `json:"echoReturnLoss"`

	EchoReturnLossEnhancement float64 `json:"echoReturnLossEnhancement"`

	DroppedSamplesDuration float64 `json:"droppedSamplesDuration"`

	DroppedSamplesEvents uint64 `json:"droppedSamplesEvents"`

	TotalCaptureDelay float64 `json:"totalCaptureDelay"`

	TotalSamplesCaptured uint64 `json:"totalSamplesCaptured"`
}

func (s AudioSourceStats) statsMarker() { _ = "STUB: not implemented"; return }

type VideoSourceStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	TrackIdentifier string `json:"trackIdentifier"`

	Kind string `json:"kind"`

	Width uint32 `json:"width"`

	Height uint32 `json:"height"`

	Frames uint32 `json:"frames"`

	FramesPerSecond float64 `json:"framesPerSecond"`
}

func (s VideoSourceStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalMediaSourceStats(b []byte) (Stats, error) {
	_ = "STUB: not implemented"
	return *new(Stats), nil
}

type AudioPlayoutStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	Kind string `json:"kind"`

	SynthesizedSamplesDuration float64 `json:"synthesizedSamplesDuration"`

	SynthesizedSamplesEvents uint64 `json:"synthesizedSamplesEvents"`

	TotalSamplesDuration float64 `json:"totalSamplesDuration"`

	TotalPlayoutDelay float64 `json:"totalPlayoutDelay"`

	TotalSamplesCount uint64 `json:"totalSamplesCount"`
}

func (s AudioPlayoutStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalMediaPlayoutStats(b []byte) (Stats, error) {
	_ = "STUB: not implemented"
	return *new(Stats), nil
}

type PeerConnectionStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	DataChannelsOpened uint32 `json:"dataChannelsOpened"`

	DataChannelsClosed uint32 `json:"dataChannelsClosed"`

	DataChannelsRequested uint32 `json:"dataChannelsRequested"`

	DataChannelsAccepted uint32 `json:"dataChannelsAccepted"`
}

func (s PeerConnectionStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalPeerConnectionStats(b []byte) (PeerConnectionStats, error) {
	_ = "STUB: not implemented"
	return *new(PeerConnectionStats), nil
}

type DataChannelStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	Label string `json:"label"`

	Protocol string `json:"protocol"`

	DataChannelIdentifier int32 `json:"dataChannelIdentifier"`

	TransportID string `json:"transportId"`

	State DataChannelState `json:"state"`

	MessagesSent uint32 `json:"messagesSent"`

	BytesSent uint64 `json:"bytesSent"`

	MessagesReceived uint32 `json:"messagesReceived"`

	BytesReceived uint64 `json:"bytesReceived"`
}

func (s DataChannelStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalDataChannelStats(b []byte) (DataChannelStats, error) {
	_ = "STUB: not implemented"
	return *new(DataChannelStats), nil
}

type MediaStreamStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	StreamIdentifier string `json:"streamIdentifier"`

	TrackIDs []string `json:"trackIds"`
}

func (s MediaStreamStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalStreamStats(b []byte) (MediaStreamStats, error) {
	_ = "STUB: not implemented"
	return *new(MediaStreamStats), nil
}

type AudioSenderStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	TrackIdentifier string `json:"trackIdentifier"`

	RemoteSource bool `json:"remoteSource"`

	Ended bool `json:"ended"`

	Kind string `json:"kind"`

	AudioLevel float64 `json:"audioLevel"`

	TotalAudioEnergy float64 `json:"totalAudioEnergy"`

	VoiceActivityFlag bool `json:"voiceActivityFlag"`

	TotalSamplesDuration float64 `json:"totalSamplesDuration"`

	EchoReturnLoss float64 `json:"echoReturnLoss"`

	EchoReturnLossEnhancement float64 `json:"echoReturnLossEnhancement"`

	TotalSamplesSent uint64 `json:"totalSamplesSent"`
}

func (s AudioSenderStats) statsMarker() { _ = "STUB: not implemented"; return }

type SenderAudioTrackAttachmentStats AudioSenderStats

func (s SenderAudioTrackAttachmentStats) statsMarker() { _ = "STUB: not implemented"; return }

type VideoSenderStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	Kind string `json:"kind"`

	FramesCaptured uint32 `json:"framesCaptured"`

	FramesSent uint32 `json:"framesSent"`

	HugeFramesSent uint32 `json:"hugeFramesSent"`

	KeyFramesSent uint32 `json:"keyFramesSent"`
}

func (s VideoSenderStats) statsMarker() { _ = "STUB: not implemented"; return }

type SenderVideoTrackAttachmentStats VideoSenderStats

func (s SenderVideoTrackAttachmentStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalSenderStats(b []byte) (Stats, error) {
	_ = "STUB: not implemented"
	return *new(Stats), nil
}

func unmarshalTrackStats(b []byte) (Stats, error) {
	_ = "STUB: not implemented"
	return *new(Stats), nil
}

type AudioReceiverStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	Kind string `json:"kind"`

	AudioLevel float64 `json:"audioLevel"`

	TotalAudioEnergy float64 `json:"totalAudioEnergy"`

	VoiceActivityFlag bool `json:"voiceActivityFlag"`

	TotalSamplesDuration float64 `json:"totalSamplesDuration"`

	EstimatedPlayoutTimestamp StatsTimestamp `json:"estimatedPlayoutTimestamp"`

	JitterBufferDelay float64 `json:"jitterBufferDelay"`

	JitterBufferEmittedCount uint64 `json:"jitterBufferEmittedCount"`

	TotalSamplesReceived uint64 `json:"totalSamplesReceived"`

	ConcealedSamples uint64 `json:"concealedSamples"`

	ConcealmentEvents uint64 `json:"concealmentEvents"`
}

func (s AudioReceiverStats) statsMarker() { _ = "STUB: not implemented"; return }

type VideoReceiverStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	Kind string `json:"kind"`

	FrameWidth uint32 `json:"frameWidth"`

	FrameHeight uint32 `json:"frameHeight"`

	FramesPerSecond float64 `json:"framesPerSecond"`

	EstimatedPlayoutTimestamp StatsTimestamp `json:"estimatedPlayoutTimestamp"`

	JitterBufferDelay float64 `json:"jitterBufferDelay"`

	JitterBufferEmittedCount uint64 `json:"jitterBufferEmittedCount"`

	FramesReceived uint32 `json:"framesReceived"`

	KeyFramesReceived uint32 `json:"keyFramesReceived"`

	FramesDecoded uint32 `json:"framesDecoded"`

	FramesDropped uint32 `json:"framesDropped"`

	PartialFramesLost uint32 `json:"partialFramesLost"`

	FullFramesLost uint32 `json:"fullFramesLost"`
}

func (s VideoReceiverStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalReceiverStats(b []byte) (Stats, error) {
	_ = "STUB: not implemented"
	return *new(Stats), nil
}

type TransportStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	PacketsSent uint32 `json:"packetsSent"`

	PacketsReceived uint32 `json:"packetsReceived"`

	BytesSent uint64 `json:"bytesSent"`

	BytesReceived uint64 `json:"bytesReceived"`

	RTCPTransportStatsID string `json:"rtcpTransportStatsId"`

	ICERole ICERole `json:"iceRole"`

	DTLSState DTLSTransportState `json:"dtlsState"`

	ICEState ICETransportState `json:"iceState"`

	SelectedCandidatePairID string `json:"selectedCandidatePairId"`

	LocalCertificateID string `json:"localCertificateId"`

	RemoteCertificateID string `json:"remoteCertificateId"`

	DTLSCipher string `json:"dtlsCipher"`

	SRTPCipher string `json:"srtpCipher"`
}

func (s TransportStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalTransportStats(b []byte) (TransportStats, error) {
	_ = "STUB: not implemented"
	return *new(TransportStats), nil
}

type StatsICECandidatePairState string

func toStatsICECandidatePairState(state ice.CandidatePairState) (StatsICECandidatePairState, error) {
	_ = "STUB: not implemented"
	return *new(StatsICECandidatePairState), nil
}

func toICECandidatePairStats(candidatePairStats ice.CandidatePairStats) (ICECandidatePairStats, error) {
	_ = "STUB: not implemented"
	return *new(ICECandidatePairStats), nil
}

const (
	StatsICECandidatePairStateFrozen StatsICECandidatePairState = "frozen"

	StatsICECandidatePairStateWaiting StatsICECandidatePairState = "waiting"

	StatsICECandidatePairStateInProgress StatsICECandidatePairState = "in-progress"

	StatsICECandidatePairStateFailed StatsICECandidatePairState = "failed"

	StatsICECandidatePairStateSucceeded StatsICECandidatePairState = "succeeded"
)

type ICECandidatePairStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	TransportID string `json:"transportId"`

	LocalCandidateID string `json:"localCandidateId"`

	RemoteCandidateID string `json:"remoteCandidateId"`

	State StatsICECandidatePairState `json:"state"`

	Nominated bool `json:"nominated"`

	PacketsSent uint32 `json:"packetsSent"`

	PacketsReceived uint32 `json:"packetsReceived"`

	BytesSent uint64 `json:"bytesSent"`

	BytesReceived uint64 `json:"bytesReceived"`

	LastPacketSentTimestamp StatsTimestamp `json:"lastPacketSentTimestamp"`

	LastPacketReceivedTimestamp StatsTimestamp `json:"lastPacketReceivedTimestamp"`

	FirstRequestTimestamp StatsTimestamp `json:"firstRequestTimestamp"`

	LastRequestTimestamp StatsTimestamp `json:"lastRequestTimestamp"`

	FirstResponseTimestamp StatsTimestamp `json:"firstResponseTimestamp"`

	LastResponseTimestamp StatsTimestamp `json:"lastResponseTimestamp"`

	FirstRequestReceivedTimestamp StatsTimestamp `json:"firstRequestReceivedTimestamp"`

	LastRequestReceivedTimestamp StatsTimestamp `json:"lastRequestReceivedTimestamp"`

	TotalRoundTripTime float64 `json:"totalRoundTripTime"`

	CurrentRoundTripTime float64 `json:"currentRoundTripTime"`

	AvailableOutgoingBitrate float64 `json:"availableOutgoingBitrate"`

	AvailableIncomingBitrate float64 `json:"availableIncomingBitrate"`

	CircuitBreakerTriggerCount uint32 `json:"circuitBreakerTriggerCount"`

	RequestsReceived uint64 `json:"requestsReceived"`

	RequestsSent uint64 `json:"requestsSent"`

	ResponsesReceived uint64 `json:"responsesReceived"`

	ResponsesSent uint64 `json:"responsesSent"`

	RetransmissionsReceived uint64 `json:"retransmissionsReceived"`

	RetransmissionsSent uint64 `json:"retransmissionsSent"`

	ConsentRequestsSent uint64 `json:"consentRequestsSent"`

	ConsentExpiredTimestamp StatsTimestamp `json:"consentExpiredTimestamp"`

	PacketsDiscardedOnSend uint32 `json:"packetsDiscardedOnSend"`

	BytesDiscardedOnSend uint32 `json:"bytesDiscardedOnSend"`
}

func (s ICECandidatePairStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalICECandidatePairStats(b []byte) (ICECandidatePairStats, error) {
	_ = "STUB: not implemented"
	return *new(ICECandidatePairStats), nil
}

type ICECandidateStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	TransportID string `json:"transportId"`

	NetworkType string `json:"networkType,omitempty"`

	IP string `json:"ip"`

	Port int32 `json:"port"`

	Protocol string `json:"protocol"`

	CandidateType ICECandidateType `json:"candidateType"`

	Priority int32 `json:"priority"`

	URL string `json:"url"`

	RelayProtocol string `json:"relayProtocol"`

	Deleted bool `json:"deleted"`
}

func (s ICECandidateStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalICECandidateStats(b []byte) (ICECandidateStats, error) {
	_ = "STUB: not implemented"
	return *new(ICECandidateStats), nil
}

type CertificateStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	Fingerprint string `json:"fingerprint"`

	FingerprintAlgorithm string `json:"fingerprintAlgorithm"`

	Base64Certificate string `json:"base64Certificate"`

	IssuerCertificateID string `json:"issuerCertificateId"`
}

func (s CertificateStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalCertificateStats(b []byte) (CertificateStats, error) {
	_ = "STUB: not implemented"
	return *new(CertificateStats), nil
}

type SCTPTransportStats struct {
	Timestamp StatsTimestamp `json:"timestamp"`

	Type StatsType `json:"type"`

	ID string `json:"id"`

	TransportID string `json:"transportId"`

	SmoothedRoundTripTime float64 `json:"smoothedRoundTripTime"`

	CongestionWindow uint32 `json:"congestionWindow"`

	ReceiverWindow uint32 `json:"receiverWindow"`

	MTU uint32 `json:"mtu"`

	UNACKData uint32 `json:"unackData"`

	Metadata *SCTPTransportMetadata `json:"metadata,omitempty"`

	BytesSent uint64 `json:"bytesSent"`

	BytesReceived uint64 `json:"bytesReceived"`
}

func (s SCTPTransportStats) statsMarker() { _ = "STUB: not implemented"; return }

func unmarshalSCTPTransportStats(b []byte) (SCTPTransportStats, error) {
	_ = "STUB: not implemented"
	return *new(SCTPTransportStats), nil
}
