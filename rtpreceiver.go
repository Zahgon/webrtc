//go:build !js

package webrtc

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/stats"
	"github.com/pion/logging"
	"github.com/pion/rtcp"
	"github.com/pion/srtp/v3"
)

type trackStreams struct {
	track *TrackRemote

	streamInfo, repairStreamInfo *interceptor.StreamInfo

	rtpReadStream  *srtp.ReadStreamSRTP
	rtpInterceptor interceptor.RTPReader

	rtcpReadStream  *srtp.ReadStreamSRTCP
	rtcpInterceptor interceptor.RTCPReader

	repairReadStream    *srtp.ReadStreamSRTP
	repairInterceptor   interceptor.RTPReader
	repairStreamChannel chan rtxPacketWithAttributes

	repairRtcpReadStream  *srtp.ReadStreamSRTCP
	repairRtcpInterceptor interceptor.RTCPReader
}

type rtxPacketWithAttributes struct {
	pkt        []byte
	attributes interceptor.Attributes
	pool       *sync.Pool
}

func (p *rtxPacketWithAttributes) release() { _ = "STUB: not implemented"; return }

type RTPReceiver struct {
	kind      RTPCodecType
	transport *DTLSTransport

	tracks []trackStreams

	closed               atomic.Bool
	closedChan, received chan any
	mu                   sync.RWMutex

	tr *RTPTransceiver

	api *API

	rtxPool sync.Pool

	log logging.LeveledLogger
}

func (api *API) NewRTPReceiver(kind RTPCodecType, transport *DTLSTransport) (*RTPReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RTPReceiver) setRTPTransceiver(tr *RTPTransceiver) { _ = "STUB: not implemented"; return }

func (r *RTPReceiver) Transport() *DTLSTransport { _ = "STUB: not implemented"; return nil }

func (r *RTPReceiver) getParameters() RTPParameters {
	_ = "STUB: not implemented"
	return *new(RTPParameters)
}

func (r *RTPReceiver) GetParameters() RTPParameters {
	_ = "STUB: not implemented"
	return *new(RTPParameters)
}

func (r *RTPReceiver) Track() *TrackRemote { _ = "STUB: not implemented"; return nil }

func (r *RTPReceiver) Tracks() []*TrackRemote { _ = "STUB: not implemented"; return nil }

func (r *RTPReceiver) RTPTransceiver() *RTPTransceiver { _ = "STUB: not implemented"; return nil }

func (r *RTPReceiver) configureReceive(parameters RTPReceiveParameters) {
	_ = "STUB: not implemented"
	return
}

func (r *RTPReceiver) startReceive(parameters RTPReceiveParameters) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

func (r *RTPReceiver) Receive(parameters RTPReceiveParameters) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RTPReceiver) Read(b []byte) (n int, a interceptor.Attributes, err error) {
	_ = "STUB: not implemented"
	return 0, *new(interceptor.Attributes), nil
}

func (r *RTPReceiver) ReadSimulcast(b []byte, rid string) (n int, a interceptor.Attributes, err error) {
	_ = "STUB: not implemented"
	return 0, *new(interceptor.Attributes), nil
}

func (r *RTPReceiver) ReadRTCP() ([]rtcp.Packet, interceptor.Attributes, error) {
	_ = "STUB: not implemented"
	return nil, *new(interceptor.Attributes), nil
}

func (r *RTPReceiver) ReadSimulcastRTCP(rid string) ([]rtcp.Packet, interceptor.Attributes, error) {
	_ = "STUB: not implemented"
	return nil, *new(interceptor.Attributes), nil
}

func (r *RTPReceiver) haveReceived() bool { _ = "STUB: not implemented"; return false }

func (r *RTPReceiver) haveClosed() bool { _ = "STUB: not implemented"; return false }

func (r *RTPReceiver) Stop() error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

func (r *RTPReceiver) collectStats(collector *statsReportCollector, statsGetter stats.Getter) {
	_ = "STUB: not implemented"
	return
}

func (r *RTPReceiver) populateInboundStats(
	inboundStats *InboundRTPStreamStats,
	statsGetter stats.Getter,
	remoteTrack *TrackRemote,
) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

//nolint:gosec

func (r *RTPReceiver) collectAudioPlayoutStats(
	collector *statsReportCollector,
	nowTime time.Time,
	remoteTrack *TrackRemote,
) {
	_ = "STUB: not implemented"
	return
}

func (r *RTPReceiver) streamsForTrack(t *TrackRemote) *trackStreams {
	_ = "STUB: not implemented"
	return nil
}

func (r *RTPReceiver) readRTP(b []byte, reader *TrackRemote) (n int, a interceptor.Attributes, err error) {
	_ = "STUB: not implemented"
	return 0, *new(interceptor.Attributes), nil
}

func (r *RTPReceiver) receiveForRid(
	rid string,
	params RTPParameters,
	streamInfo *interceptor.StreamInfo,
	rtpReadStream *srtp.ReadStreamSRTP,
	rtpInterceptor interceptor.RTPReader,
	rtcpReadStream *srtp.ReadStreamSRTCP,
	rtcpInterceptor interceptor.RTCPReader,
	peekedPackets []*peekedPacket,
) (*TrackRemote, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RTPReceiver) receiveForRtx(
	ssrc SSRC,
	rsid string,
	streamInfo *interceptor.StreamInfo,
	rtpReadStream *srtp.ReadStreamSRTP,
	rtpInterceptor interceptor.RTPReader,
	rtcpReadStream *srtp.ReadStreamSRTCP,
	rtcpInterceptor interceptor.RTCPReader,
) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gocognit,cyclop
func (r *RTPReceiver) receiveForRtxInternal(
	ssrc SSRC,
	rsid string,
	streamInfo *interceptor.StreamInfo,
	rtpReadStream *srtp.ReadStreamSRTP,
	rtpInterceptor interceptor.RTPReader,
	rtcpReadStream *srtp.ReadStreamSRTCP,
	rtcpInterceptor interceptor.RTCPReader,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RTPReceiver) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (r *RTPReceiver) SetReadDeadlineSimulcast(deadline time.Time, rid string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RTPReceiver) setRTPReadDeadline(deadline time.Time, reader *TrackRemote) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RTPReceiver) readRTX(reader *TrackRemote) *rtxPacketWithAttributes {
	_ = "STUB: not implemented"
	return nil
}
