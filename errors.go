package webrtc

import (
	"errors"
)

var (
	ErrUnknownType = errors.New("unknown")

	ErrConnectionClosed = errors.New("connection closed")

	ErrDataChannelNotOpen = errors.New("data channel not open")

	ErrCertificateExpired = errors.New("x509Cert expired")

	ErrNoTurnCredentials = errors.New("turn server credentials required")

	ErrTurnCredentials = errors.New("invalid turn server credentials")

	ErrExistingTrack = errors.New("track already exists")

	ErrPrivateKeyType = errors.New("private key type not supported")

	ErrModifyingPeerIdentity = errors.New("peerIdentity cannot be modified")

	ErrModifyingCertificates = errors.New("certificates cannot be modified")

	ErrModifyingBundlePolicy = errors.New("bundle policy cannot be modified")

	ErrModifyingRTCPMuxPolicy = errors.New("rtcp mux policy cannot be modified")

	ErrModifyingICECandidatePoolSize = errors.New("ice candidate pool size cannot be modified")

	ErrStringSizeLimit = errors.New("data channel label exceeds size limit")

	ErrMaxDataChannelID = errors.New("maximum number ID for datachannel specified")

	ErrNegotiatedWithoutID = errors.New("negotiated set without channel id")

	ErrRetransmitsOrPacketLifeTime = errors.New("both MaxPacketLifeTime and MaxRetransmits was set")

	ErrCodecNotFound = errors.New("codec not found")

	ErrNoRemoteDescription = errors.New("remote description is not set")

	ErrIncorrectSDPSemantics = errors.New("remote SessionDescription semantics does not match configuration")

	ErrIncorrectSignalingState = errors.New("operation can not be run in current signaling state")

	ErrProtocolTooLarge = errors.New("protocol is larger then 65535 bytes")

	ErrSenderNotCreatedByConnection = errors.New("RtpSender not created by this PeerConnection")

	ErrSessionDescriptionNoFingerprint = errors.New("SetRemoteDescription called with no fingerprint")

	ErrSessionDescriptionInvalidFingerprint = errors.New("SetRemoteDescription called with an invalid fingerprint")

	ErrSessionDescriptionConflictingFingerprints = errors.New(
		"SetRemoteDescription called with multiple conflicting fingerprint",
	)

	ErrSessionDescriptionMissingIceUfrag = errors.New("SetRemoteDescription called with no ice-ufrag")

	ErrSessionDescriptionMissingIcePwd = errors.New("SetRemoteDescription called with no ice-pwd")

	ErrSessionDescriptionConflictingIceUfrag = errors.New(
		"SetRemoteDescription called with multiple conflicting ice-ufrag values",
	)

	ErrSessionDescriptionConflictingIcePwd = errors.New(
		"SetRemoteDescription called with multiple conflicting ice-pwd values",
	)

	ErrNoSRTPProtectionProfile = errors.New("DTLS Handshake completed and no SRTP Protection Profile was chosen")

	ErrFailedToGenerateCertificateFingerprint = errors.New("failed to generate certificate fingerprint")

	ErrNoCodecsAvailable = errors.New("operation failed no codecs are available")

	ErrUnsupportedCodec = errors.New("unable to start track, codec is not supported by remote")

	ErrSenderWithNoCodecs = errors.New("unable to populate media section, RTPSender created with no codecs")

	ErrCodecAlreadyRegistered = errors.New("codec already registered for same payload type")

	ErrRTPSenderNewTrackHasIncorrectKind = errors.New("new track must be of the same kind as previous")

	ErrRTPSenderNewTrackHasIncorrectEnvelope = errors.New("new track must have the same envelope as previous")

	ErrUnbindFailed = errors.New("failed to unbind TrackLocal from PeerConnection")

	ErrNoPayloaderForCodec = errors.New("the requested codec does not have a payloader")

	ErrRegisterHeaderExtensionInvalidDirection = errors.New(
		"a header extension must be registered as 'recvonly', 'sendonly' or both",
	)

	ErrSimulcastProbeOverflow = errors.New("simulcast probe limit has been reached, new SSRC has been discarded")

	ErrSDPUnmarshalling = errors.New("failed to unmarshal SDP")

	errDetachNotEnabled                 = errors.New("enable detaching by calling webrtc.DetachDataChannels()")
	errDetachBeforeOpened               = errors.New("datachannel not opened yet, try calling Detach from OnOpen")
	errDtlsTransportNotStarted          = errors.New("the DTLS transport has not started yet")
	errDtlsKeyExtractionFailed          = errors.New("failed extracting keys from DTLS for SRTP")
	errFailedToStartSRTP                = errors.New("failed to start SRTP")
	errFailedToStartSRTCP               = errors.New("failed to start SRTCP")
	errInvalidDTLSStart                 = errors.New("attempted to start DTLSTransport that is not in new state")
	errNoRemoteCertificate              = errors.New("peer didn't provide certificate via DTLS")
	errIdentityProviderNotImplemented   = errors.New("identity provider is not implemented")
	errNoMatchingCertificateFingerprint = errors.New("remote certificate does not match any fingerprint")

	errICEConnectionNotStarted        = errors.New("ICE connection not started")
	errICECandidateTypeUnknown        = errors.New("unknown candidate type")
	errICEInvalidConvertCandidateType = errors.New(
		"cannot convert ice.CandidateType into webrtc.ICECandidateType, invalid type",
	)
	errICEAgentNotExist            = errors.New("ICEAgent does not exist")
	errICECandiatesCoversionFailed = errors.New("unable to convert ICE candidates to ICECandidates")
	errICERoleUnknown              = errors.New("unknown ICE Role")
	errICEProtocolUnknown          = errors.New("unknown protocol")
	errICEGathererNotStarted       = errors.New("gatherer not started")
	errAddressRewriteWithNAT1To1   = errors.New("address rewrite rules cannot be combined with NAT1To1IPs")

	errNetworkTypeUnknown = errors.New("unknown network type")

	errSDPDoesNotMatchOffer        = errors.New("new sdp does not match previous offer")
	errSDPDoesNotMatchAnswer       = errors.New("new sdp does not match previous answer")
	errPeerConnSDPTypeInvalidValue = errors.New(
		"provided value is not a valid enum value of type SDPType",
	)
	errPeerConnStateChangeInvalid                     = errors.New("invalid state change op")
	errPeerConnStateChangeUnhandled                   = errors.New("unhandled state change op")
	errPeerConnSDPTypeInvalidValueSetLocalDescription = errors.New("invalid SDP type supplied to SetLocalDescription()")
	errPeerConnRemoteDescriptionWithoutMidValue       = errors.New(
		"remoteDescription contained media section without mid value",
	)
	errPeerConnRemoteDescriptionNil                  = errors.New("remoteDescription has not been set yet")
	errMediaSectionHasExplictSSRCAttribute           = errors.New("media section has an explicit SSRC")
	errPeerConnRemoteSSRCAddTransceiver              = errors.New("could not add transceiver for remote SSRC")
	errPeerConnSimulcastMidRTPExtensionRequired      = errors.New("mid RTP Extensions required for Simulcast")
	errPeerConnSimulcastStreamIDRTPExtensionRequired = errors.New("stream id RTP Extensions required for Simulcast")
	errPeerConnSimulcastIncomingSSRCFailed           = errors.New("incoming SSRC failed Simulcast probing")
	errPeerConnAddTransceiverFromKindOnlyAcceptsOne  = errors.New(
		"AddTransceiverFromKind only accepts one RTPTransceiverInit",
	)
	errPeerConnAddTransceiverFromTrackOnlyAcceptsOne = errors.New(
		"AddTransceiverFromTrack only accepts one RTPTransceiverInit",
	)
	errPeerConnAddTransceiverFromKindSupport = errors.New(
		"AddTransceiverFromKind currently only supports recvonly",
	)
	errPeerConnAddTransceiverFromTrackSupport = errors.New(
		"AddTransceiverFromTrack currently only supports sendonly and sendrecv",
	)
	errPeerConnSetIdentityProviderNotImplemented = errors.New("TODO SetIdentityProvider")
	errPeerConnWriteRTCPOpenWriteStream          = errors.New("WriteRTCP failed to open WriteStream")
	errPeerConnTranscieverMidNil                 = errors.New("cannot find transceiver with mid")
	errPeerConnEarlyMediaWithoutAnswer           = errors.New(
		"cannot process early media without SDP answer," +
			"use SettingEngine.SetHandleUndeclaredSSRCWithoutAnswer(true) to process without answer",
	)

	errRTPReceiverDTLSTransportNil            = errors.New("DTLSTransport must not be nil")
	errRTPReceiverReceiveAlreadyCalled        = errors.New("Receive has already been called")
	errRTPReceiverWithSSRCTrackStreamNotFound = errors.New("unable to find stream for Track with SSRC")
	errRTPReceiverForRIDTrackStreamNotFound   = errors.New("no trackStreams found for RID")

	errRTPSenderTrackNil             = errors.New("Track must not be nil")
	errRTPSenderDTLSTransportNil     = errors.New("DTLSTransport must not be nil")
	errRTPSenderSendAlreadyCalled    = errors.New("Send has already been called")
	errRTPSenderSendNotCalled        = errors.New("Send has not been called")
	errRTPSenderStopped              = errors.New("Sender has already been stopped")
	errRTPSenderTrackRemoved         = errors.New("Sender Track has been removed or replaced to nil")
	errRTPSenderRidNil               = errors.New("Sender cannot add encoding as rid is empty")
	errRTPSenderNoBaseEncoding       = errors.New("Sender cannot add encoding as there is no base track")
	errRTPSenderBaseEncodingMismatch = errors.New("Sender cannot add encoding as provided track does not match base track")
	errRTPSenderRIDCollision         = errors.New("Sender cannot encoding due to RID collision")
	errRTPSenderNoTrackForRID        = errors.New("Sender does not have track for RID")

	errRTPTransceiverCannotChangeMid        = errors.New("cannot change transceiver mid")
	errRTPTransceiverSetSendingInvalidState = errors.New("invalid state change in RTPTransceiver.setSending")
	errRTPTransceiverCodecUnsupported       = errors.New("unsupported codec type by this transceiver")

	errSCTPTransportDTLS = errors.New("DTLS not established")

	errSDPZeroTransceivers                 = errors.New("addTransceiverSDP() called with 0 transceivers")
	errSDPMediaSectionMediaDataChanInvalid = errors.New("invalid Media Section. Media + DataChannel both enabled")
	errSDPMediaSectionMultipleTrackInvalid = errors.New(
		"invalid Media Section. Can not have multiple tracks in one MediaSection in UnifiedPlan",
	)

	errSettingEngineSetAnsweringDTLSRole = errors.New("SetAnsweringDTLSRole must DTLSRoleClient or DTLSRoleServer")

	errSignalingStateCannotRollback            = errors.New("can't rollback from stable state")
	errSignalingStateProposedTransitionInvalid = errors.New("invalid proposed signaling state transition")

	errStatsICECandidateStateInvalid = errors.New(
		"cannot convert to StatsICECandidatePairStateSucceeded invalid ice candidate state",
	)

	errICECandidatePoolSizeTooLarge = errors.New("ice candidate pool size greater than 1 is not supported")

	errInvalidICECredentialTypeString = errors.New("invalid ICECredentialType")
	errInvalidICEServer               = errors.New("invalid ICEServer")

	errICETransportNotInNew = errors.New("ICETransport can only be called in ICETransportStateNew")
	errICETransportClosed   = errors.New("ICETransport closed")

	errCertificatePEMMultipleCert = errors.New("failed parsing certificate, more than 1 CERTIFICATE block in pems")
	errCertificatePEMMultiplePriv = errors.New("failed parsing certificate, more than 1 PRIVATE KEY block in pems")
	errCertificatePEMMissing      = errors.New("failed parsing certificate, pems must contain both a CERTIFICATE block and a PRIVATE KEY block")

	errRTPTooShort = errors.New("not long enough to be a RTP Packet")

	errExcessiveRetries = errors.New("excessive retries in CreateOffer")
)
