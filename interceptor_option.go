//go:build !js

package webrtc

import (
	"github.com/pion/interceptor/pkg/nack"
	"github.com/pion/interceptor/pkg/report"
	"github.com/pion/interceptor/pkg/stats"
	"github.com/pion/interceptor/pkg/twcc"
	"github.com/pion/logging"
)

type interceptorOptions struct {
	loggerFactory logging.LoggerFactory

	nackGeneratorOptions  []nack.GeneratorOption
	nackResponderOptions  []nack.ResponderOption
	reportReceiverOptions []report.ReceiverOption
	reportSenderOptions   []report.SenderOption
	statsOptions          []stats.Option
	twccOptions           []twcc.Option
}

type InterceptorOption func(*interceptorOptions)

func WithInterceptorLoggerFactory(loggerFactory logging.LoggerFactory) InterceptorOption {
	_ = "STUB: not implemented"
	return *new(InterceptorOption)
}

func WithNackGeneratorOptions(opts ...nack.GeneratorOption) InterceptorOption {
	_ = "STUB: not implemented"
	return *new(InterceptorOption)
}

func WithNackResponderOptions(opts ...nack.ResponderOption) InterceptorOption {
	_ = "STUB: not implemented"
	return *new(InterceptorOption)
}

func WithReportReceiverOptions(opts ...report.ReceiverOption) InterceptorOption {
	_ = "STUB: not implemented"
	return *new(InterceptorOption)
}

func WithReportSenderOptions(opts ...report.SenderOption) InterceptorOption {
	_ = "STUB: not implemented"
	return *new(InterceptorOption)
}

func WithStatsInterceptorOptions(opts ...stats.Option) InterceptorOption {
	_ = "STUB: not implemented"
	return *new(InterceptorOption)
}

func WithTWCCOptions(opts ...twcc.Option) InterceptorOption {
	_ = "STUB: not implemented"
	return *new(InterceptorOption)
}
