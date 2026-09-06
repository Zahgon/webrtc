//go:build !js

package webrtc

import (
	"sync"
	"sync/atomic"

	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/flexfec"
	"github.com/pion/interceptor/pkg/nack"
	"github.com/pion/interceptor/pkg/report"
	"github.com/pion/interceptor/pkg/rfc8888"
	"github.com/pion/interceptor/pkg/stats"
	"github.com/pion/interceptor/pkg/twcc"
	"github.com/pion/rtp"
)

func RegisterDefaultInterceptors(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry) error {
	_ = "STUB: not implemented"
	return nil
}

func RegisterDefaultInterceptorsWithOptions(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry,
	opts ...InterceptorOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureStatsInterceptor(interceptorRegistry *interceptor.Registry) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureStatsInterceptorWithOptions(interceptorRegistry *interceptor.Registry, opts ...stats.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func lookupStats(id string) (stats.Getter, bool) {
	_ = "STUB: not implemented"
	return *new(stats.Getter), false
}

func cleanupStats(id string) { _ = "STUB: not implemented"; return }

var statsGetter sync.Map

func ConfigureRTCPReports(interceptorRegistry *interceptor.Registry) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureRTCPReportsWithOptions(interceptorRegistry *interceptor.Registry, recvOpts []report.ReceiverOption,
	sendOpts ...report.SenderOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureNack(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureNackWithOptions(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry,
	genOpts []nack.GeneratorOption, respOpts ...nack.ResponderOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureTWCCHeaderExtensionSender(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureTWCCSender(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureTWCCSenderWithOptions(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry,
	opts ...twcc.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureCongestionControlFeedback(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureCongestionControlFeedbackWithOptions(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry,
	opts ...rfc8888.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureSimulcastExtensionHeaders(mediaEngine *MediaEngine) error {
	_ = "STUB: not implemented"
	return nil
}

func ConfigureFlexFEC03(
	payloadType PayloadType,
	mediaEngine *MediaEngine,
	interceptorRegistry *interceptor.Registry,
	options ...flexfec.FecOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

type interceptorToTrackLocalWriter struct{ interceptor atomic.Value }

func (i *interceptorToTrackLocalWriter) WriteRTP(header *rtp.Header, payload []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *interceptorToTrackLocalWriter) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:unparam
func createStreamInfo(
	id string,
	ssrc, ssrcRTX, ssrcFEC SSRC,
	payloadType, payloadTypeRTX, payloadTypeFEC PayloadType,
	codec RTPCodecCapability,
	webrtcHeaderExtensions []RTPHeaderExtensionParameter,
) *interceptor.StreamInfo {
	_ = "STUB: not implemented"
	return nil
}
