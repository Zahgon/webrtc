package oggwriter

import (
	"errors"
	"io"
	"os"
	"sync"

	"github.com/pion/rtp"
)

const (
	pageHeaderTypeContinuationOfStream = 0x00
	pageHeaderTypeContinuationOfPacket = 0x01
	pageHeaderTypeBeginningOfStream    = 0x02
	pageHeaderTypeEndOfStream          = 0x04
	defaultPreSkip                     = 3840
	defaultSampleRate                  = 48000
	opusGranuleSampleRate              = 48000
	maxOpusPacketSamples               = opusGranuleSampleRate * 120 / 1000
	defaultChannelCount                = 2
	channelMappingFamily0              = 0
	channelMappingFamily1              = 1
	channelMappingFamily2              = 2
	channelMappingFamily255            = 255
	idPageSignature                    = "OpusHead"
	commentPageSignature               = "OpusTags"
	defaultVendor                      = "pion"
	pageHeaderSignature                = "OggS"
	pageHeaderSize                     = 27
	maxOggPageSegments                 = 255
	noGranulePosition                  = ^uint64(0)
	maxUint32Length                    = uint64(1<<32 - 1)
)

var (
	errFileNotOpened        = errors.New("file not opened")
	errOutputNotOpened      = errors.New("output not opened")
	errInvalidNilPacket     = errors.New("invalid nil packet")
	errDuplicateTrackSSRC   = errors.New("duplicate Ogg track SSRC")
	errDuplicateTrackSerial = errors.New("duplicate Ogg track serial")
	errTracksStarted        = errors.New("cannot add Ogg tracks after writing has started")
	errPacketSSRCMismatch   = errors.New("RTP packet SSRC does not match Ogg track SSRC")
	errInvalidOpusPacket    = errors.New("invalid Opus packet")
	errInvalidChannelCount  = errors.New("invalid channel count")
	errInvalidChannelMap    = errors.New("invalid channel mapping")
	errInvalidOpusTags      = errors.New("invalid OpusTags")
)

type pageRewriter interface {
	io.Seeker
	io.WriterAt
}

type writerConfig struct {
	sampleRate     uint32
	channelMapping channelMapping
	pageRewriter   pageRewriter
	opusTags       OpusTags
}

type trackConfig struct {
	sampleRate        uint32
	channelMapping    channelMapping
	channelMappingSet bool
	serial            uint32
	serialSet         bool
	opusTags          OpusTags
}

type channelMapping struct {
	family       uint8
	channelCount uint8
	streamCount  uint8
	coupledCount uint8
	mapping      []byte
}

type WriterOption interface {
	applyWriterConfig(*writerConfig) error
}

type TrackOption interface {
	applyTrackConfig(*trackConfig) error
}

type WriterTrackOption interface {
	WriterOption
	TrackOption
}

type writerOptionFunc func(*writerConfig) error

func (f writerOptionFunc) applyWriterConfig(config *writerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

type trackOptionFunc func(*trackConfig) error

func (f trackOptionFunc) applyTrackConfig(config *trackConfig) error {
	_ = "STUB: not implemented"
	return nil
}

type sampleRateOption uint32

func (o sampleRateOption) applyWriterConfig(config *writerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (o sampleRateOption) applyTrackConfig(config *trackConfig) error {
	_ = "STUB: not implemented"
	return nil
}

type (
	channelCountOption uint16
	channelMapOption   struct {
		family       uint8
		streamCount  uint8
		coupledCount uint8
		mapping      []byte
	}
	vendorOption       string
	userCommentsOption []UserComment
)

func (o channelCountOption) applyWriterConfig(config *writerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (o channelCountOption) applyTrackConfig(config *trackConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (o channelMapOption) applyWriterConfig(config *writerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (o channelMapOption) applyTrackConfig(config *trackConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (o vendorOption) applyWriterConfig(config *writerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (o vendorOption) applyTrackConfig(config *trackConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (o userCommentsOption) applyWriterConfig(config *writerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (o userCommentsOption) applyTrackConfig(config *trackConfig) error {
	_ = "STUB: not implemented"
	return nil
}

type OpusTags struct {
	Vendor       string
	UserComments []UserComment
}

type UserComment struct {
	Comment string
	Value   string
}

type oggTrack struct {
	sampleRate              uint32
	channelMapping          channelMapping
	preSkip                 uint16
	serial                  uint32
	opusTags                OpusTags
	pageIndex               uint32
	previousGranulePosition uint64
	lastPayload             []byte
	lastGranulePosition     uint64
	lastPageIndex           uint32
	lastPageOffset          int64
	lastPageHeaderType      uint8
	lastPageWritten         bool
}

type oggPage struct {
	data       []byte
	payload    []byte
	headerType uint8
	granulePos uint64
	pageIndex  uint32
}

type OggWriter struct {
	mu            sync.Mutex
	stream        io.Writer
	fd            *os.File
	checksumTable *[256]uint32
	track         *oggTrack
}

type Writer struct {
	mu             sync.Mutex
	stream         io.Writer
	pageRewriter   pageRewriter
	sampleRate     uint32
	channelMapping channelMapping
	opusTags       OpusTags
	checksumTable  *[256]uint32
	tracks         map[uint32]*Track
	trackOrder     []*Track
	started        bool
}

type Track struct {
	parent *Writer
	ssrc   uint32
	track  *oggTrack
}

func WithSerial(serial uint32) TrackOption { _ = "STUB: not implemented"; return *new(TrackOption) }

func WithSampleRate(sampleRate uint32) WriterTrackOption {
	_ = "STUB: not implemented"
	return *new(WriterTrackOption)
}

func WithChannelCount(channelCount uint16) WriterTrackOption {
	_ = "STUB: not implemented"
	return *new(WriterTrackOption)
}

func WithChannelMapping(family, streamCount, coupledCount uint8, mapping []byte) WriterTrackOption {
	_ = "STUB: not implemented"
	return *new(WriterTrackOption)
}

func WithVendor(vendor string) WriterTrackOption {
	_ = "STUB: not implemented"
	return *new(WriterTrackOption)
}

func WithUserComments(comments ...UserComment) WriterTrackOption {
	_ = "STUB: not implemented"
	return *new(WriterTrackOption)
}

func WithSeekableOutput(output interface {
	io.Seeker
	io.WriterAt
},
) WriterOption {
	_ = "STUB: not implemented"
	return *new(WriterOption)
}

func New(fileName string, sampleRate uint32, channelCount uint16) (*OggWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

func NewWith(out io.Writer, sampleRate uint32, channelCount uint16) (*OggWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWith(out io.Writer, fd *os.File, config *trackConfig) (*OggWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWriter(out io.Writer, opts ...WriterOption) (*Writer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Writer) NewTrack(ssrc uint32, opts ...TrackOption) (*Track, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *OggWriter) WriteRTP(packet *rtp.Packet) error { _ = "STUB: not implemented"; return nil }

func (w *Track) WriteRTP(packet *rtp.Packet) error { _ = "STUB: not implemented"; return nil }

func (w *OggWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }

func newTrackConfig(sampleRate uint32, channelCount uint16) (*trackConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTrackState(sampleRate uint32, mapping channelMapping, serial uint32, opusTags OpusTags) *oggTrack {
	_ = "STUB: not implemented"
	return nil
}

func defaultChannelMapping(channelCount uint16) (channelMapping, error) {
	_ = "STUB: not implemented"
	return *new(channelMapping), nil
}

func validateChannelMapping(family, streamCount, coupledCount uint8, mapping []byte) (channelMapping, error) {
	_ = "STUB: not implemented"
	return *new(channelMapping), nil
}

//nolint:gosec // validated <= MaxUint8.

func validateChannelMappingFamily(family uint8) error { _ = "STUB: not implemented"; return nil }

func validateChannelMappingLayout(streamCount, coupledCount uint8, mapping []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFamilySpecificChannelMapping(
	family uint8,
	streamCount uint8,
	coupledCount uint8,
	mapping []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateVorbisChannelMapping(streamCount, coupledCount uint8, mapping []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func cloneChannelMapping(mapping channelMapping) channelMapping {
	_ = "STUB: not implemented"
	return *new(channelMapping)
}

func defaultOpusTags() OpusTags { _ = "STUB: not implemented"; return *new(OpusTags) }

func cloneOpusTags(opusTags OpusTags) OpusTags { _ = "STUB: not implemented"; return *new(OpusTags) }

func cloneUserComments(comments []UserComment) []UserComment { _ = "STUB: not implemented"; return nil }

func applyVendor(opusTags *OpusTags, vendor string) error { _ = "STUB: not implemented"; return nil }

func applyUserComments(opusTags *OpusTags, comments []UserComment) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOpusTags(opusTags OpusTags) error { _ = "STUB: not implemented"; return nil }

func validateOpusTagsWithMaxHeaderLen(opusTags OpusTags, maxHeaderLen uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOpusTagsHeaderLen(opusTags OpusTags, maxHeaderLen uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func validateUserComments(comments []UserComment) error { _ = "STUB: not implemented"; return nil }

func validateUserComment(comment UserComment) error { _ = "STUB: not implemented"; return nil }

func validateOpusTagString(field, value string) error { _ = "STUB: not implemented"; return nil }

func isValidCommentName(comment string) bool { _ = "STUB: not implemented"; return false }

func (w *Writer) serialInUse(serial uint32) bool { _ = "STUB: not implemented"; return false }

func (w *Writer) allocateSerial() uint32 { _ = "STUB: not implemented"; return 0 }

func (w *OggWriter) writeTrackHeaders(track *oggTrack) error { _ = "STUB: not implemented"; return nil }

func (w *Writer) writeTrackIDHeader(track *oggTrack) error { _ = "STUB: not implemented"; return nil }

func (w *Writer) writeTrackCommentHeader(track *oggTrack) error {
	_ = "STUB: not implemented"
	return nil
}

func writeTrackHeaders(track *oggTrack, writePageFunc func(*oggTrack, []byte, uint8, uint64) error) error {
	_ = "STUB: not implemented"
	return nil
}

func writeTrackIDHeader(track *oggTrack, writePageFunc func(*oggTrack, []byte, uint8, uint64) error) error {
	_ = "STUB: not implemented"
	return nil
}

func writeTrackCommentHeader(track *oggTrack, writePageFunc func(*oggTrack, []byte, uint8, uint64) error) error {
	_ = "STUB: not implemented"
	return nil
}

func buildIDHeader(track *oggTrack) []byte { _ = "STUB: not implemented"; return nil }

func buildCommentHeader(opusTags OpusTags) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec // validated to fit in uint32.

//nolint:gosec // validated to fit in uint32.

//nolint:gosec // validated to fit in uint32.

func writePage(
	stream io.Writer,
	rewriter pageRewriter,
	checksumTable *[256]uint32,
	track *oggTrack,
	payload []byte,
	headerType uint8,
	granulePos uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // Page counts are bounded by the packet size in memory.

func createPagesForSerial(
	checksumTable *[256]uint32,
	payload []byte,
	headerType uint8,
	granulePos uint64,
	serial uint32,
	pageIndex uint32,
) []oggPage {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // remainingPayload is < 255 here.

func packetPageHeaderType(headerType uint8, firstPage, packetComplete bool) uint8 {
	_ = "STUB: not implemented"
	return 0
}

func createPageForSerial(
	checksumTable *[256]uint32,
	payload []byte,
	headerType uint8,
	granulePos uint64,
	serial uint32,
	pageIndex uint32,
) []byte {
	_ = "STUB: not implemented"
	return nil
}

func createPageForSerialWithSegments(
	checksumTable *[256]uint32,
	payload []byte,
	segmentTable []byte,
	headerType uint8,
	granulePos uint64,
	serial uint32,
	pageIndex uint32,
) []byte {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // segmentTable is capped at maxOggPageSegments.

func (w *OggWriter) writeRTP(track *oggTrack, packet *rtp.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Writer) writeOpusPayload(track *oggTrack, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func writeRTP(
	stream io.Writer,
	rewriter pageRewriter,
	checksumTable *[256]uint32,
	track *oggTrack,
	packet *rtp.Packet,
) error {
	_ = "STUB: not implemented"
	return nil
}

func opusPayloadFromPacket(packet *rtp.Packet) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func writeOpusPayload(
	stream io.Writer,
	rewriter pageRewriter,
	checksumTable *[256]uint32,
	track *oggTrack,
	payload []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func opusPacketSampleCount(payload []byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func opusPacketFrameCount(payload []byte) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

func opusSamplesPerFrame(toc byte) uint32 { _ = "STUB: not implemented"; return 0 }

func (w *Writer) startLocked() error { _ = "STUB: not implemented"; return nil }

func markTrackEndOfStream(rewriter pageRewriter, checksumTable *[256]uint32, track *oggTrack) error {
	_ = "STUB: not implemented"
	return nil
}

func writeNilEndOfStreamPage(stream io.Writer, checksumTable *[256]uint32, track *oggTrack) error {
	_ = "STUB: not implemented"
	return nil
}

func pageRewriterForFile(fd *os.File) pageRewriter {
	_ = "STUB: not implemented"
	return *new(pageRewriter)
}

func writeToStream(stream io.Writer, p []byte) error { _ = "STUB: not implemented"; return nil }

func generateChecksumTable() *[256]uint32 { _ = "STUB: not implemented"; return nil }

//nolint:gosec // G115

//nolint:gosec // no out of bounds access here.
