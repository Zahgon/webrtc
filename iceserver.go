//go:build !js

package webrtc

import (
	"github.com/pion/stun/v3"
)

type ICEServer struct {
	URLs           []string          `json:"urls"`
	Username       string            `json:"username,omitempty"`
	Credential     any               `json:"credential,omitempty"`
	CredentialType ICECredentialType `json:"credentialType,omitempty"`
}

func (s ICEServer) parseURL(i int) (*stun.URI, error) { _ = "STUB: not implemented"; return nil, nil }

func (s ICEServer) validate() error { _ = "STUB: not implemented"; return nil }

func (s ICEServer) urls() ([]*stun.URI, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil, nil
}

func iceserverUnmarshalUrls(val any) (*[]string, error) { _ = "STUB: not implemented"; return nil, nil }

func iceserverUnmarshalOauth(val any) (*OAuthCredential, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ICEServer) iceserverUnmarshalFields(fields map[string]any) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

func (s *ICEServer) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s ICEServer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
