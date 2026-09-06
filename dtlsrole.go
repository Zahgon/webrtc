package webrtc

import (
	"github.com/pion/sdp/v3"
)

type DTLSRole byte

const (
	DTLSRoleUnknown DTLSRole = iota

	DTLSRoleAuto

	DTLSRoleClient

	DTLSRoleServer
)

const (
	defaultDtlsRoleAnswer = DTLSRoleClient

	defaultDtlsRoleOffer = DTLSRoleAuto
)

func (r DTLSRole) String() string { _ = "STUB: not implemented"; return "" }

func dtlsRoleFromSDP(sessionDescription *sdp.SessionDescription) DTLSRole {
	_ = "STUB: not implemented"
	return *new(DTLSRole)
}

func connectionRoleFromDtlsRole(d DTLSRole) sdp.ConnectionRole {
	_ = "STUB: not implemented"
	return *new(sdp.ConnectionRole)
}
