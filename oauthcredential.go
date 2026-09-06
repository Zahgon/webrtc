package webrtc

type OAuthCredential struct {
	MACKey string

	AccessToken string //nolint:gosec // not a secret.
}
