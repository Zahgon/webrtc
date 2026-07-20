//go:build !js

package webrtc

import (
	"net"
	"sync"

	"github.com/pion/logging"
	"github.com/pion/sctp"
)

const sctpMaxChannels = uint16(65535)

func newSCTPTransportMetadata(metadata sctp.AssociationMetadata) SCTPTransportMetadata {
	_ = "STUB: not implemented"
	return *new(SCTPTransportMetadata)
}

type SCTPTransport struct {
	lock sync.RWMutex

	dtlsTransport *DTLSTransport

	state SCTPTransportState

	isStarted bool

	maxChannels *uint16

	onErrorHandler func(error)
	onCloseHandler func(error)

	sctpAssociation            *sctp.Association
	onDataChannelHandler       func(*DataChannel)
	onDataChannelOpenedHandler func(*DataChannel)

	dataChannels          []*DataChannel
	dataChannelIDsUsed    map[uint16]struct{}
	dataChannelsOpened    uint32
	dataChannelsRequested uint32
	dataChannelsAccepted  uint32

	localSctpInit []byte

	api *API
	log logging.LeveledLogger
}

func (api *API) NewSCTPTransport(dtls *DTLSTransport) *SCTPTransport {
	_ = "STUB: not implemented"
	return nil
}

func (r *SCTPTransport) Transport() *DTLSTransport { _ = "STUB: not implemented"; return nil }

func (r *SCTPTransport) GetCapabilities() SCTPCapabilities {
	_ = "STUB: not implemented"
	return *new(SCTPCapabilities)
}

//nolint:cyclop
func (r *SCTPTransport) Start(capabilities SCTPCapabilities) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *SCTPTransport) sctpClientOptions(netConn net.Conn, maxMessageSize uint32) []sctp.ClientOption {
	_ = "STUB: not implemented"
	return nil
}

func (r *SCTPTransport) optionalSCTPClientOptions() []sctp.ClientOption {
	_ = "STUB: not implemented"
	return nil
}

func (r *SCTPTransport) Stop() error { _ = "STUB: not implemented"; return nil }

//nolint:cyclop
func (r *SCTPTransport) acceptDataChannels(
	assoc *sctp.Association,
	existingDataChannels []*DataChannel,
) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec //G115

func (r *SCTPTransport) OnError(f func(err error)) { _ = "STUB: not implemented"; return }

func (r *SCTPTransport) onError(err error) { _ = "STUB: not implemented"; return }

func (r *SCTPTransport) OnClose(f func(err error)) { _ = "STUB: not implemented"; return }

func (r *SCTPTransport) onClose(err error) { _ = "STUB: not implemented"; return }

func (r *SCTPTransport) OnDataChannel(f func(*DataChannel)) { _ = "STUB: not implemented"; return }

func (r *SCTPTransport) OnDataChannelOpened(f func(*DataChannel)) {
	_ = "STUB: not implemented"
	return
}

func (r *SCTPTransport) onDataChannel(dc *DataChannel) (done chan struct{}) {
	_ = "STUB: not implemented"
	return nil
}

func (r *SCTPTransport) updateMaxChannels() { _ = "STUB: not implemented"; return }

func (r *SCTPTransport) MaxChannels() uint16 { _ = "STUB: not implemented"; return 0 }

func (r *SCTPTransport) State() SCTPTransportState {
	_ = "STUB: not implemented"
	return *new(SCTPTransportState)
}

func (r *SCTPTransport) Metadata() (SCTPTransportMetadata, bool) {
	_ = "STUB: not implemented"
	return *new(SCTPTransportMetadata), false
}

func (r *SCTPTransport) Stats() SCTPTransportStats {
	_ = "STUB: not implemented"
	return *new(SCTPTransportStats)
}

func (r *SCTPTransport) collectStats(collector *statsReportCollector) {
	_ = "STUB: not implemented"
	return
}

func (r *SCTPTransport) generateAndSetDataChannelID(dtlsRole DTLSRole, idOut **uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *SCTPTransport) association() *sctp.Association { _ = "STUB: not implemented"; return nil }

func (r *SCTPTransport) BufferedAmount() int { _ = "STUB: not implemented"; return 0 }

func (r *SCTPTransport) GetSctpInit() []byte { _ = "STUB: not implemented"; return nil }
