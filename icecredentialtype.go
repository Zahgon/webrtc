package webrtc

type ICECredentialType int

const (
	ICECredentialTypePassword ICECredentialType = iota

	ICECredentialTypeOauth
)

const (
	iceCredentialTypePasswordStr = "password"
	iceCredentialTypeOauthStr    = "oauth"
)

func newICECredentialType(raw string) (ICECredentialType, error) {
	_ = "STUB: not implemented"
	return *new(ICECredentialType), nil
}

func (t ICECredentialType) String() string { _ = "STUB: not implemented"; return "" }

func (t *ICECredentialType) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (t ICECredentialType) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
