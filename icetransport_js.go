//go:build js && wasm
// +build js,wasm

package webrtc

import "syscall/js"

type ICETransport struct {
	underlying js.Value
}

func (t *ICETransport) JSValue() js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func (t *ICETransport) GetSelectedCandidatePair() (*ICECandidatePair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
