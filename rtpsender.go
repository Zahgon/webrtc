//go:build !js

package webrtc

import (
	"sync"
	"time"

	"github.com/pion/interceptor"
	"github.com/pion/rtcp"
)

type trackEncoding struct {
	track TrackLocal

	srtpStream *srtpWriterFuture

	rtcpInterceptor interceptor.RTCPReader
	streamInfo      interceptor.StreamInfo

	context *baseTrackLocalContext

	ssrc, ssrcRTX, ssrcFEC SSRC
}

type RTPSender struct {
	trackEncodings []*trackEncoding

	transport *DTLSTransport

	payloadType PayloadType
	kind        RTPCodecType

	negotiated bool

	api *API
	id  string

	rtpTransceiver *RTPTransceiver

	mu                     sync.RWMutex
	sendCalled, stopCalled chan struct{}
}

func (api *API) NewRTPSender(track TrackLocal, transport *DTLSTransport) (*RTPSender, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RTPSender) isNegotiated() bool { _ = "STUB: not implemented"; return false }

func (r *RTPSender) setNegotiated() { _ = "STUB: not implemented"; return }

func (r *RTPSender) setRTPTransceiver(rtpTransceiver *RTPTransceiver) {
	_ = "STUB: not implemented"
	return
}

func (r *RTPSender) Transport() *DTLSTransport { _ = "STUB: not implemented"; return nil }

func (r *RTPSender) GetParameters() RTPSendParameters {
	_ = "STUB: not implemented"
	return *new(RTPSendParameters)
}

func (r *RTPSender) AddEncoding(track TrackLocal) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

func (r *RTPSender) addEncoding(track TrackLocal) { _ = "STUB: not implemented"; return }

func (r *RTPSender) Track() TrackLocal { _ = "STUB: not implemented"; return *new(TrackLocal) }

func (r *RTPSender) ReplaceTrack(track TrackLocal) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

func (r *RTPSender) Send(parameters RTPSendParameters) error { _ = "STUB: not implemented"; return nil }

func (r *RTPSender) Stop() error { _ = "STUB: not implemented"; return nil }

func (r *RTPSender) Read(b []byte) (n int, a interceptor.Attributes, err error) {
	_ = "STUB: not implemented"
	return 0, *new(interceptor.Attributes), nil
}

func (r *RTPSender) ReadRTCP() ([]rtcp.Packet, interceptor.Attributes, error) {
	_ = "STUB: not implemented"
	return nil, *new(interceptor.Attributes), nil
}

func (r *RTPSender) ReadSimulcast(b []byte, rid string) (n int, a interceptor.Attributes, err error) {
	_ = "STUB: not implemented"
	return 0, *new(interceptor.Attributes), nil
}

func (r *RTPSender) ReadSimulcastRTCP(rid string) ([]rtcp.Packet, interceptor.Attributes, error) {
	_ = "STUB: not implemented"
	return nil, *new(interceptor.Attributes), nil
}

func (r *RTPSender) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (r *RTPSender) SetReadDeadlineSimulcast(deadline time.Time, rid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RTPSender) hasSent() bool { _ = "STUB: not implemented"; return false }

func (r *RTPSender) hasStopped() bool { _ = "STUB: not implemented"; return false }

func (r *RTPSender) configureRTXAndFEC() { _ = "STUB: not implemented"; return }
