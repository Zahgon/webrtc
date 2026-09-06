//go:build js && wasm
// +build js,wasm

package webrtc

import (
	"syscall/js"
)

type RTPTransceiver struct {
	underlying js.Value
}

func (r *RTPTransceiver) JSValue() js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func (r *RTPTransceiver) Direction() RTPTransceiverDirection {
	_ = "STUB: not implemented"
	return *new(RTPTransceiverDirection)
}

func (r *RTPTransceiver) Sender() *RTPSender { _ = "STUB: not implemented"; return nil }

func (r *RTPTransceiver) Receiver() *RTPReceiver { _ = "STUB: not implemented"; return nil }
