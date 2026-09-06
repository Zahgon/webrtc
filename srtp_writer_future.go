//go:build !js

package webrtc

import (
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pion/rtp"
)

type srtpWriterFuture struct {
	ssrc           SSRC
	rtpSender      *RTPSender
	rtcpReadStream atomic.Value
	rtpWriteStream atomic.Value
	mu             sync.Mutex
	closed         bool
}

func (s *srtpWriterFuture) init(returnWhenNoSRTP bool) error { //nolint:cyclop
	if returnWhenNoSRTP {
		select {
		case <-s.rtpSender.stopCalled:
			return io.ErrClosedPipe
		case <-s.rtpSender.transport.srtpReady:
		default:
			return nil
		}
	} else {
		select {
		case <-s.rtpSender.stopCalled:
			return io.ErrClosedPipe
		case <-s.rtpSender.transport.srtpReady:
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return io.ErrClosedPipe
	}

	srtcpSession, err := s.rtpSender.transport.getSRTCPSession()
	if err != nil {
		return err
	}

	rtcpReadStream, err := srtcpSession.OpenReadStream(uint32(s.ssrc))
	if err != nil {
		return err
	}

	srtpSession, err := s.rtpSender.transport.getSRTPSession()
	if err != nil {
		return err
	}

	rtpWriteStream, err := srtpSession.OpenWriteStream()
	if err != nil {
		return err
	}

	s.rtcpReadStream.Store(rtcpReadStream)
	s.rtpWriteStream.Store(rtpWriteStream)

	return nil
}

func (s *srtpWriterFuture) Close() error { _ = "STUB: not implemented"; return nil }

func (s *srtpWriterFuture) Read(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *srtpWriterFuture) SetReadDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *srtpWriterFuture) WriteRTP(header *rtp.Header, payload []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *srtpWriterFuture) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
