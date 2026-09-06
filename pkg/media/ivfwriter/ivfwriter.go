package ivfwriter

import (
	"errors"
	"io"

	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
)

var (
	errFileNotOpened        = errors.New("file not opened")
	errInvalidNilPacket     = errors.New("invalid nil packet")
	errCodecUnset           = errors.New("codec is unset")
	errCodecAlreadySet      = errors.New("codec is already set")
	errNoSuchCodec          = errors.New("no codec for this MimeType")
	errInvalidMediaTimebase = errors.New("invalid media timebase")
)

type (
	codec int

	IVFWriter struct {
		ioWriter     io.Writer
		count        uint64
		seenKeyFrame bool

		codec codec

		timebaseDenominator uint32
		timebaseNumerator   uint32
		firstFrameTimestamp uint32
		clockRate           uint64
		videoWidth          uint16
		videoHeight         uint16

		directPTS bool

		currentFrame []byte

		av1Depacketizer *codecs.AV1Depacketizer
	}
)

const (
	codecUnset codec = iota
	codecVP8
	codecVP9
	codecAV1

	mimeTypeVP8 = "video/VP8"
	mimeTypeVP9 = "video/VP9"
	mimeTypeAV1 = "video/AV1"
)

func New(fileName string, opts ...Option) (*IVFWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
	//nolint:gosec
}

func NewWith(out io.Writer, opts ...Option) (*IVFWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IVFWriter) writeHeader() error { _ = "STUB: not implemented"; return nil }

func (i *IVFWriter) timestampToPts(timestamp uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (i *IVFWriter) writeFrame(frame []byte, timestamp uint64) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // G115

func (i *IVFWriter) WriteRTP(packet *rtp.Packet) error { _ = "STUB: not implemented"; return nil }

func (i *IVFWriter) writeVP8(packet *rtp.Packet, timestamp uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *IVFWriter) writeVP9(packet *rtp.Packet, timestamp uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *IVFWriter) writeAV1(packet *rtp.Packet, timestamp uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *IVFWriter) Close() error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // G115

type Option func(i *IVFWriter) error

func WithCodec(mimeType string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWidthAndHeight(width, height uint16) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithFrameRate(numerator, denominator uint32) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDirectPTS() Option { _ = "STUB: not implemented"; return *new(Option) }
