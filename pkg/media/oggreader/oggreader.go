package oggreader

import (
	"errors"
	"io"
)

const (
	pageHeaderTypeBeginningOfStream = 0x02
	pageHeaderSignature             = "OggS"

	idPageBasePayloadLength = 19
	pageHeaderLen           = 27
)

var (
	errNilStream                       = errors.New("stream is nil")
	errBadIDPageSignature              = errors.New("bad header signature")
	errBadOpusTagsSignature            = errors.New("bad opus tags signature")
	errBadIDPageType                   = errors.New("wrong header, expected beginning of stream")
	errBadIDPageLength                 = errors.New("payload for id page must be 19 bytes")
	errBadIDPagePayloadSignature       = errors.New("bad payload signature")
	errShortPageHeader                 = errors.New("not enough data for payload header")
	errChecksumMismatch                = errors.New("expected and actual checksum do not match")
	errUnsupportedChannelMappingFamily = errors.New("unsupported channel mapping family")
)

type OggReader struct {
	stream               io.Reader
	bytesReadSuccesfully int64
	checksumTable        *[256]uint32
	doChecksum           bool
}

type OggHeader struct {
	ChannelMap   uint8
	Channels     uint8
	OutputGain   uint16
	PreSkip      uint16
	SampleRate   uint32
	Version      uint8
	StreamCount  uint8
	CoupledCount uint8

	ChannelMapping string
}

func ParseOpusHead(payload []byte) (*OggHeader, error) { _ = "STUB: not implemented"; return nil, nil }

type OggPageHeader struct {
	GranulePosition uint64

	sig           [4]byte
	version       uint8
	headerType    uint8
	Serial        uint32
	index         uint32
	segmentsCount uint8
}

type HeaderType string

const (
	headerUnknown  HeaderType = ""
	HeaderOpusID   HeaderType = "OpusHead"
	HeaderOpusTags HeaderType = "OpusTags"
)

func opusPayloadSignature(payload []byte) (HeaderType, bool) {
	_ = "STUB: not implemented"
	return *new(HeaderType), false
}

func (p *OggPageHeader) HeaderType(payload []byte) (HeaderType, bool) {
	_ = "STUB: not implemented"
	return *new(HeaderType), false
}

type Option func(*OggReader) error

func NewWith(in io.Reader) (*OggReader, *OggHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func NewWithOptions(in io.Reader, options ...Option) (*OggReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithDoChecksum(doChecksum bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func newWith(in io.Reader, doChecksum bool) (*OggReader, *OggHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (o *OggReader) readOpusHeader() (*OggHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateOpusPageHeader(pageHeader *OggPageHeader, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func parseBasicHeaderFields(payload []byte) *OggHeader { _ = "STUB: not implemented"; return nil }

func parseChannelMapping(header *OggHeader, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePayloadLength(payload []byte, expectedLen int) error {
	_ = "STUB: not implemented"
	return nil
}

func parseExtendedChannelMapping(header *OggHeader, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *OggReader) ParseNextPage() ([]byte, *OggPageHeader, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil, nil, nil
}

func (o *OggReader) ResetReader(reset func(bytesRead int64) io.Reader) {
	_ = "STUB: not implemented"
	return
}

func generateChecksumTable() *[256]uint32 { _ = "STUB: not implemented"; return nil }

//nolint:gosec // G115

//nolint:gosec // no out of bounds access here.

type OpusTags struct {
	Vendor       string
	UserComments []UserComment
}

type UserComment struct {
	Comment string
	Value   string
}

func ParseOpusTags(payload []byte) (*OpusTags, error) { _ = "STUB: not implemented"; return nil, nil }

func validateOpusTagsHeader(payload []byte, minHeaderLen int) error {
	_ = "STUB: not implemented"
	return nil
}

func parseVendorString(payload []byte, headerMagicLen, u32Size, minHeaderLen int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func parseUserComments(payload []byte, vendorEnd, u32Size int) ([]UserComment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSingleUserComment(payload []byte, pos, u32Size, index int) (UserComment, int, error) {
	_ = "STUB: not implemented"
	return *new(UserComment), 0, nil
}
