//go:build js && wasm
// +build js,wasm

package webrtc

import "syscall/js"

type RTPSender struct {
	underlying js.Value
}

func (s *RTPSender) JSValue() js.Value { _ = "STUB: not implemented"; return *new(js.Value) }
