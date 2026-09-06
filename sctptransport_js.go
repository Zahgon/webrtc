//go:build js && wasm
// +build js,wasm

package webrtc

import "syscall/js"

type SCTPTransport struct {
	underlying js.Value
}

func (r *SCTPTransport) JSValue() js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func (r *SCTPTransport) Transport() *DTLSTransport { _ = "STUB: not implemented"; return nil }

func (r *SCTPTransport) Metadata() (SCTPTransportMetadata, bool) {
	_ = "STUB: not implemented"
	return *new(SCTPTransportMetadata), false
}
