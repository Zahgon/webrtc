package h265writer

import (
	"io"

	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v4/pkg/media/h265reader"
)

const (
	typeAP = 48
	typeFU = 49
)

type H265Writer struct {
	writer       io.Writer
	hasKeyFrame  bool
	cachedPacket *codecs.H265Depacketizer
}

func New(filename string) (*H265Writer, error) {
	_ = "STUB: not implemented"
	return nil,
		//nolint:gosec
		nil
}

func NewWith(w io.Writer) *H265Writer { _ = "STUB: not implemented"; return nil }

func (h *H265Writer) WriteRTP(packet *rtp.Packet) error { _ = "STUB: not implemented"; return nil }

func (h *H265Writer) Close() error { _ = "STUB: not implemented"; return nil }

func isKeyFrame(data []byte) bool { _ = "STUB: not implemented"; return false }

func checkAggregationPacketForKeyFrame(data []byte) bool { _ = "STUB: not implemented"; return false }

func isKeyFrameNalu(naluType h265reader.NalUnitType) bool { _ = "STUB: not implemented"; return false }
