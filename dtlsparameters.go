package webrtc

type DTLSParameters struct {
	Role         DTLSRole          `json:"role"`
	Fingerprints []DTLSFingerprint `json:"fingerprints"`
}
