//go:build js && wasm
// +build js,wasm

package webrtc

import (
	"github.com/pion/ice/v4"
)

type ICEServer struct {
	URLs     []string
	Username string

	Credential     any
	CredentialType ICECredentialType
}

func (s ICEServer) parseURL(i int) (*ice.URL, error) { _ = "STUB: not implemented"; return nil, nil }

func (s ICEServer) validate() ([]*ice.URL, error) { _ = "STUB: not implemented"; return nil, nil }
