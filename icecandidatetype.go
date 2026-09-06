package webrtc

import (
	"github.com/pion/ice/v4"
)

type ICECandidateType int

const (
	ICECandidateTypeUnknown ICECandidateType = iota

	ICECandidateTypeHost

	ICECandidateTypeSrflx

	ICECandidateTypePrflx

	ICECandidateTypeRelay
)

const (
	iceCandidateTypeHostStr  = "host"
	iceCandidateTypeSrflxStr = "srflx"
	iceCandidateTypePrflxStr = "prflx"
	iceCandidateTypeRelayStr = "relay"
)

func NewICECandidateType(raw string) (ICECandidateType, error) {
	_ = "STUB: not implemented"
	return *new(ICECandidateType), nil
}

func (t ICECandidateType) String() string { _ = "STUB: not implemented"; return "" }

func getCandidateType(candidateType ice.CandidateType) (ICECandidateType, error) {
	_ = "STUB: not implemented"
	return *new(ICECandidateType), nil
}

func (t ICECandidateType) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented" //nolint:staticcheck
	return nil, nil
}

func (t *ICECandidateType) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }

func (r ICECandidateType) toICE() ice.CandidateType {
	_ = "STUB: not implemented"
	//nolint:gosec // G115, no overflow, ICECandidateType matches ice.CandidateType in granularity.
	return *new(ice.CandidateType)
}
