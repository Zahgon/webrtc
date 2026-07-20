package webrtc

type ICECandidatePair struct {
	statsID string
	Local   *ICECandidate
	Remote  *ICECandidate
}

func newICECandidatePairStatsID(localID, remoteID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *ICECandidatePair) String() string { _ = "STUB: not implemented"; return "" }

func NewICECandidatePair(local, remote *ICECandidate) *ICECandidatePair {
	_ = "STUB: not implemented"
	return nil
}
