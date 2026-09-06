//go:build !js

package webrtc

import (
	"github.com/pion/logging"
	"github.com/pion/sdp/v3"
)

type trackDetails struct {
	mid      string
	kind     RTPCodecType
	streamID string
	id       string
	ssrcs    []SSRC
	rtxSsrc  *SSRC
	fecSsrc  *SSRC
	rids     []string
}

func trackDetailsForSSRC(trackDetails []trackDetails, ssrc SSRC) *trackDetails {
	_ = "STUB: not implemented"
	return nil
}

func trackDetailsForRID(trackDetails []trackDetails, mid, rid string) *trackDetails {
	_ = "STUB: not implemented"
	return nil
}

func filterTrackWithSSRC(incomingTracks []trackDetails, ssrc SSRC) []trackDetails {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gocognit,gocyclo,cyclop
func trackDetailsFromSDP(
	log logging.LeveledLogger,
	s *sdp.SessionDescription,
) (incomingTracks []trackDetails) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:nestif

//nolint:gosec // G115

//nolint:gosec // G115

func trackDetailsToRTPReceiveParameters(trackDetails *trackDetails) RTPReceiveParameters {
	_ = "STUB: not implemented"
	return *new(RTPReceiveParameters)
}

func getRids(media *sdp.MediaDescription) []*simulcastRid { _ = "STUB: not implemented"; return nil }

func addCandidatesToMediaDescriptions(
	candidates []ICECandidate,
	mediaDescr *sdp.MediaDescription,
	iceGatheringState ICEGatheringState,
) error {
	_ = "STUB: not implemented"
	return nil
}

func addDataMediaSection(
	descr *sdp.SessionDescription,
	shouldAddCandidates bool,
	dtlsFingerprints []DTLSFingerprint,
	midValue string,
	iceParams ICEParameters,
	candidates []ICECandidate,
	dtlsRole sdp.ConnectionRole,
	iceGatheringState ICEGatheringState,
	sctpMaxMessageSize uint32,
	sctpInit []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func populateLocalCandidates(
	sessionDescription *SessionDescription,
	i *ICEGatherer,
	iceGatheringState ICEGatheringState,
) *SessionDescription {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gocognit,cyclop
func addSenderSDP(
	mediaSection mediaSection,
	isPlanB bool,
	media *sdp.MediaDescription,
) {
	_ = "STUB: not implemented"
	return
}

//nolint:cyclop, gocognit
func addTransceiverSDP(
	descr *sdp.SessionDescription,
	isPlanB bool,
	shouldAddCandidates bool,
	dtlsFingerprints []DTLSFingerprint,
	mediaEngine *MediaEngine,
	midValue string,
	iceParams ICEParameters,
	candidates []ICECandidate,
	dtlsRole sdp.ConnectionRole,
	iceGatheringState ICEGatheringState,
	mediaSection mediaSection,
	ignoreRidPauseForRecv bool,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type simulcastRid struct {
	id        string
	attrValue string
	paused    bool
}

type mediaSection struct {
	id              string
	transceivers    []*RTPTransceiver
	data            bool
	sctpInit        []byte
	matchExtensions map[string]int
	rids            []*simulcastRid
}

func bundleMatchFromRemote(matchBundleGroup *string) func(mid string) bool {
	_ = "STUB: not implemented"
	return nil
}

//nolint:cyclop
func populateSDP(
	descr *sdp.SessionDescription,
	isPlanB bool,
	dtlsFingerprints []DTLSFingerprint,
	mediaDescriptionFingerprint bool,
	isICELite bool,
	isExtmapAllowMixed bool,
	mediaEngine *MediaEngine,
	connectionRole sdp.ConnectionRole,
	candidates []ICECandidate,
	iceParams ICEParameters,
	mediaSections []mediaSection,
	iceGatheringState ICEGatheringState,
	matchBundleGroup *string,
	sctpMaxMessageSize uint32,
	ignoreRidPauseForRecv bool,
) (*sdp.SessionDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMidValue(media *sdp.MediaDescription) string { _ = "STUB: not implemented"; return "" }

func descriptionIsPlanB(desc *SessionDescription, log logging.LeveledLogger) bool {
	_ = "STUB: not implemented"
	return false
}

func descriptionPossiblyPlanB(desc *SessionDescription) bool {
	_ = "STUB: not implemented"
	return false
}

func getPeerDirection(media *sdp.MediaDescription) RTPTransceiverDirection {
	_ = "STUB: not implemented"
	return *new(RTPTransceiverDirection)
}

func extractBundleID(desc *sdp.SessionDescription) string { _ = "STUB: not implemented"; return "" }

func extractFingerprint(desc *sdp.SessionDescription) (string, string, error) {
	_ = "STUB: not implemented" //nolint:gocognit,cyclop
	return "", "", nil
}

//nolint:nestif

type identifiedMediaDescription struct {
	MediaDescription *sdp.MediaDescription
	SDPMid           string
	SDPMLineIndex    uint16
}

func extractICEDetailsFromMedia( //nolint:cyclop
	media *identifiedMediaDescription,
	log logging.LeveledLogger,
) (string, string, []ICECandidate, error) {
	_ = "STUB: not implemented"
	return "", "", nil, nil
}

type sdpICEDetails struct {
	Ufrag      string
	Password   string //nolint:gosec // not a secret.
	Candidates []ICECandidate
}

func extractICEDetails(
	desc *sdp.SessionDescription,
	log logging.LeveledLogger,
) (*sdpICEDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func selectCandidateMediaSection(sessionDescription *sdp.SessionDescription) (
	descr *identifiedMediaDescription,
	ok bool,
) {
	_ = "STUB: not implemented"
	return nil, false
}

//nolint:gosec // G115

//nolint:gosec // G115

func getByMid(searchMid string, desc *SessionDescription) *sdp.MediaDescription {
	_ = "STUB: not implemented"
	return nil
}

func haveDataChannel(desc *SessionDescription) *sdp.MediaDescription {
	_ = "STUB: not implemented"
	return nil
}

func codecsFromMediaDescription(mediaDescr *sdp.MediaDescription) (out []RTPCodecParameters, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rtpExtensionsFromMediaDescription(m *sdp.MediaDescription) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateSDPOrigin(origin *sdp.Origin, descr *sdp.SessionDescription) {
	_ = "STUB: not implemented"
	return
}

func isIceLiteSet(desc *sdp.SessionDescription) bool { _ = "STUB: not implemented"; return false }

func isExtMapAllowMixedSet(desc *sdp.SessionDescription) bool {
	_ = "STUB: not implemented"
	return false
}

func getMaxMessageSize(desc *sdp.MediaDescription) uint32 { _ = "STUB: not implemented"; return 0 }

func getSctpInit(desc *sdp.MediaDescription) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
