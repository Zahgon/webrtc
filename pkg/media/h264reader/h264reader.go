package h264reader

import (
	"errors"
	"io"
)

type H264Reader struct {
	stream                      io.Reader
	nalBuffer                   []byte
	countOfConsecutiveZeroBytes int
	nalPrefixParsed             bool
	readBuffer                  []byte
	tmpReadBuf                  []byte
	includeSEI                  bool
}

var (
	errNilReader           = errors.New("stream is nil")
	errDataIsNotH264Stream = errors.New("data is not a H264 bitstream")
)

func NewReader(in io.Reader) (*H264Reader, error) { _ = "STUB: not implemented"; return nil, nil }

type Option func(*H264Reader) error

func NewReaderWithOptions(in io.Reader, options ...Option) (*H264Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithIncludeSEI(include bool) Option { _ = "STUB: not implemented"; return *new(Option) }

type NAL struct {
	PictureOrderCount uint32

	ForbiddenZeroBit bool
	RefIdc           uint8
	UnitType         NalUnitType

	Data []byte
}

func (reader *H264Reader) read(numToRead int) (data []byte, e error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (reader *H264Reader) bitStreamStartsWithH264Prefix() (prefixLength int, e error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (reader *H264Reader) NextNAL() (*NAL, error) { _ = "STUB: not implemented"; return nil, nil }

func (reader *H264Reader) processByte(readByte byte) (nalFound bool) {
	_ = "STUB: not implemented"
	return false
}

func newNal(data []byte) *NAL { _ = "STUB: not implemented"; return nil }

func (h *NAL) parseHeader() { _ = "STUB: not implemented"; return }
