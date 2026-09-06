package h264writer

import (
	"io"

	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
)

type (
	H264Writer struct {
		writer       io.Writer
		hasKeyFrame  bool
		cachedPacket *codecs.H264Packet
	}
)

func New(filename string) (*H264Writer, error) {
	_ = "STUB: not implemented"
	return nil,
		//nolint:gosec
		nil
}

func NewWith(w io.Writer) *H264Writer { _ = "STUB: not implemented"; return nil }

func (h *H264Writer) WriteRTP(packet *rtp.Packet) error { _ = "STUB: not implemented"; return nil }

func (h *H264Writer) Close() error { _ = "STUB: not implemented"; return nil }

func isKeyFrame(data []byte) bool { _ = "STUB: not implemented"; return false }
