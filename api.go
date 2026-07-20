//go:build !js

package webrtc

import (
	"github.com/pion/interceptor"
)

type API struct {
	settingEngine       *SettingEngine
	mediaEngine         *MediaEngine
	interceptorRegistry *interceptor.Registry

	interceptor interceptor.Interceptor
}

func NewAPI(options ...func(*API)) *API { _ = "STUB: not implemented"; return nil }

func WithMediaEngine(m *MediaEngine) func(a *API) { _ = "STUB: not implemented"; return nil }

func WithSettingEngine(s SettingEngine) func(a *API) { _ = "STUB: not implemented"; return nil }

func WithInterceptorRegistry(ir *interceptor.Registry) func(a *API) {
	_ = "STUB: not implemented"
	return nil
}
