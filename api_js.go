//go:build js && wasm
// +build js,wasm

package webrtc

type API struct {
	settingEngine *SettingEngine
}

func NewAPI(options ...func(*API)) *API { _ = "STUB: not implemented"; return nil }

func WithSettingEngine(s SettingEngine) func(a *API) { _ = "STUB: not implemented"; return nil }
