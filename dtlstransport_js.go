//go:build js && wasm
// +build js,wasm

package webrtc

import "syscall/js"

type DTLSTransport struct {
	underlying js.Value
}

func (r *DTLSTransport) JSValue() js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func (r *DTLSTransport) ICETransport() *ICETransport { _ = "STUB: not implemented"; return nil }

func (t *DTLSTransport) GetRemoteCertificate() []byte { _ = "STUB: not implemented"; return nil }
