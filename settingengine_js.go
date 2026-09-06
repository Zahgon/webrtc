//go:build js && wasm
// +build js,wasm

package webrtc

type SettingEngine struct {
	detach struct {
		DataChannels bool
	}
}

func (e *SettingEngine) DetachDataChannels() { _ = "STUB: not implemented"; return }
