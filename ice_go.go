//go:build !js

package webrtc

func (api *API) NewICETransport(gatherer *ICEGatherer) *ICETransport {
	_ = "STUB: not implemented"
	return nil
}
