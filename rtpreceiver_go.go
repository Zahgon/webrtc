//go:build !js

package webrtc

func (r *RTPReceiver) SetRTPParameters(params RTPParameters) { _ = "STUB: not implemented"; return }
