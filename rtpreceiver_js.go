//go:build js && wasm
// +build js,wasm

package webrtc

import "syscall/js"

type RTPReceiver struct {
	underlying js.Value
}

func (r *RTPReceiver) JSValue() js.Value { _ = "STUB: not implemented"; return *new(js.Value) }
