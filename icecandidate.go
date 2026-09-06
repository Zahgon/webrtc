package webrtc

import (
	"github.com/pion/ice/v4"
)

type ICECandidate struct {
	statsID        string
	Foundation     string           `json:"foundation"`
	Priority       uint32           `json:"priority"`
	Address        string           `json:"address"`
	Protocol       ICEProtocol      `json:"protocol"`
	Port           uint16           `json:"port"`
	Typ            ICECandidateType `json:"type"`
	Component      uint16           `json:"component"`
	RelatedAddress string           `json:"relatedAddress"`
	RelatedPort    uint16           `json:"relatedPort"`
	TCPType        string           `json:"tcpType"`
	SDPMid         string           `json:"sdpMid"`
	SDPMLineIndex  uint16           `json:"sdpMLineIndex"`
	extensions     string
}

func newICECandidatesFromICE(
	iceCandidates []ice.Candidate,
	sdpMid string,
	sdpMLineIndex uint16,
) ([]ICECandidate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newICECandidateFromICE(candidate ice.Candidate, sdpMid string, sdpMLineIndex uint16) (ICECandidate, error) {
	_ = "STUB: not implemented"
	return *new(ICECandidate), nil
}

//nolint:gosec // G115

//nolint:gosec // G115

func (c ICECandidate) ToICE() (cand ice.Candidate, err error) {
	_ = "STUB: not implemented"
	return *new(ice.Candidate), nil
}

func (c *ICECandidate) setExtensions(ext []ice.CandidateExtension) {
	_ = "STUB: not implemented"
	return
}

func (c *ICECandidate) exportExtensions(cand ice.Candidate) error {
	_ = "STUB: not implemented"
	return nil
}

func convertTypeFromICE(t ice.CandidateType) (ICECandidateType, error) {
	_ = "STUB: not implemented"
	return *new(ICECandidateType), nil
}

func (c ICECandidate) String() string { _ = "STUB: not implemented"; return "" }

func (c ICECandidate) ToJSON() ICECandidateInit {
	_ = "STUB: not implemented"
	return *new(ICECandidateInit)
}
