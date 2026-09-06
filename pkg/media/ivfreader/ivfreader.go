package ivfreader

import (
	"errors"
	"io"
)

const (
	ivfFileHeaderSignature = "DKIF"
	ivfFileHeaderSize      = 32
	ivfFrameHeaderSize     = 12
)

var (
	errNilStream             = errors.New("stream is nil")
	errIncompleteFrameHeader = errors.New("incomplete frame header")
	errIncompleteFrameData   = errors.New("incomplete frame data")
	errIncompleteFileHeader  = errors.New("incomplete file header")
	errSignatureMismatch     = errors.New("IVF signature mismatch")
	errUnknownIVFVersion     = errors.New("IVF version unknown, parser may not parse correctly")
	errInvalidMediaTimebase  = errors.New("invalid media timebase")
)

type IVFFileHeader struct {
	signature           string
	version             uint16
	headerSize          uint16
	FourCC              string
	Width               uint16
	Height              uint16
	TimebaseDenominator uint32
	TimebaseNumerator   uint32
	NumFrames           uint32
	unused              uint32
}

type IVFFrameHeader struct {
	FrameSize uint32
	Timestamp uint64
}

type IVFReader struct {
	stream               io.Reader
	bytesReadSuccesfully int64
	timebaseDenominator  uint32
	timebaseNumerator    uint32
}

func NewWith(stream io.Reader) (*IVFReader, *IVFFileHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (i *IVFReader) ResetReader(reset func(bytesRead int64) io.Reader) {
	_ = "STUB: not implemented"
	return
}

func (i *IVFReader) ptsToTimestamp(pts uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (i *IVFReader) ParseNextFrame() ([]byte, *IVFFrameHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (i *IVFReader) parseFileHeader() (*IVFFileHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
