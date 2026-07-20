package webrtc

type ICEGatherOptions struct {
	ICEServers           []ICEServer
	ICEGatherPolicy      ICETransportPolicy
	ICECandidatePoolSize uint8
}
