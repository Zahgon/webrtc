package webrtc

type ICEParameters struct {
	UsernameFragment string `json:"usernameFragment"`
	Password         string `json:"password"` //nolint:gosec // not a secret.
	ICELite          bool   `json:"iceLite"`
}
