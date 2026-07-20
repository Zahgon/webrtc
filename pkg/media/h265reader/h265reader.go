package h265reader

import (
	"errors"
	"io"
)

type H265Reader struct {
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
	errDataIsNotH265Stream = errors.New("data is not a H265/HEVC bitstream")
)

func (reader *H265Reader) shouldSkipNAL(naluType NalUnitType) bool {
	_ = "STUB: not implemented"
	return false
}

func NewReader(in io.Reader) (*H265Reader, error) { _ = "STUB: not implemented"; return nil, nil }

type Option func(*H265Reader) error

func NewReaderWithOptions(in io.Reader, options ...Option) (*H265Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithIncludeSEI(include bool) Option { _ = "STUB: not implemented"; return *new(Option) }

type NAL struct {
	PictureOrderCount uint32

	ForbiddenZeroBit bool
	NalUnitType      NalUnitType
	LayerID          uint8
	TemporalIDPlus1  uint8

	Data []byte
}

func (reader *H265Reader) read(numToRead int) (data []byte, e error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (reader *H265Reader) bitStreamStartsWithH265Prefix() (prefixLength int, e error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (reader *H265Reader) NextNAL() (*NAL, error) { _ = "STUB: not implemented"; return nil, nil }

func (reader *H265Reader) processByte(readByte byte) (nalFound bool) {
	_ = "STUB: not implemented"
	return false
}

func newNal(data []byte) *NAL { _ = "STUB: not implemented"; return nil }

func (h *NAL) parseHeader() { _ = "STUB: not implemented"; return }
