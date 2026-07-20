//go:build !js

package webrtc

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/pion/datachannel"
	"github.com/pion/logging"
)

var errSCTPNotEstablished = errors.New("SCTP not established")

type DataChannel struct {
	mu sync.RWMutex

	statsID                    string
	label                      string
	ordered                    bool
	maxPacketLifeTime          *uint16
	maxRetransmits             *uint16
	protocol                   string
	negotiated                 bool
	id                         *uint16
	readyState                 atomic.Value
	bufferedAmountLowThreshold uint64
	detachCalled               bool
	readLoopActive             chan struct{}
	isGracefulClosed           bool

	onMessageHandler    func(DataChannelMessage)
	openHandlerOnce     sync.Once
	onOpenHandler       func()
	dialHandlerOnce     sync.Once
	onDialHandler       func()
	closeHandlerOnce    sync.Once
	onCloseHandler      func()
	onBufferedAmountLow func()
	onErrorHandler      func(error)

	sctpTransport *SCTPTransport
	dataChannel   *datachannel.DataChannel

	api *API
	log logging.LeveledLogger
}

func (api *API) NewDataChannel(transport *SCTPTransport, params *DataChannelParameters) (*DataChannel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *API) newDataChannel(
	params *DataChannelParameters,
	sctpTransport *SCTPTransport,
	log logging.LeveledLogger,
) (*DataChannel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DataChannel) open(sctpTransport *SCTPTransport) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

func (d *DataChannel) Transport() *SCTPTransport { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) checkDetachAfterOpen() { _ = "STUB: not implemented"; return }

func (d *DataChannel) OnOpen(f func()) { _ = "STUB: not implemented"; return }

func (d *DataChannel) onOpen() { _ = "STUB: not implemented"; return }

func (d *DataChannel) OnDial(f func()) { _ = "STUB: not implemented"; return }

func (d *DataChannel) onDial() { _ = "STUB: not implemented"; return }

func (d *DataChannel) OnClose(f func()) { _ = "STUB: not implemented"; return }

func (d *DataChannel) onClose() { _ = "STUB: not implemented"; return }

func (d *DataChannel) OnMessage(f func(msg DataChannelMessage)) { _ = "STUB: not implemented"; return }

func (d *DataChannel) onMessage(msg DataChannelMessage) { _ = "STUB: not implemented"; return }

func (d *DataChannel) handleOpen(dc *datachannel.DataChannel, isRemote, isAlreadyNegotiated bool) {
	_ = "STUB: not implemented"
	return
}

func (d *DataChannel) OnError(f func(err error)) { _ = "STUB: not implemented"; return }

func (d *DataChannel) onError(err error) { _ = "STUB: not implemented"; return }

func (d *DataChannel) readLoop() { _ = "STUB: not implemented"; return }

func (d *DataChannel) Send(data []byte) error { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) SendText(s string) error { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) ensureOpen() error { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) Detach() (datachannel.ReadWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(datachannel.ReadWriteCloser), nil
}

func (d *DataChannel) DetachWithDeadline() (datachannel.ReadWriteCloserDeadliner, error) {
	_ = "STUB: not implemented"
	return *new(datachannel.ReadWriteCloserDeadliner), nil
}

func (d *DataChannel) Close() error { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) GracefulClose() error { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) close(shouldGracefullyClose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DataChannel) Label() string { _ = "STUB: not implemented"; return "" }

func (d *DataChannel) Ordered() bool { _ = "STUB: not implemented"; return false }

func (d *DataChannel) MaxPacketLifeTime() *uint16 { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) MaxRetransmits() *uint16 { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) Protocol() string { _ = "STUB: not implemented"; return "" }

func (d *DataChannel) Negotiated() bool { _ = "STUB: not implemented"; return false }

func (d *DataChannel) ID() *uint16 { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) ReadyState() DataChannelState {
	_ = "STUB: not implemented"
	return *new(DataChannelState)
}

func (d *DataChannel) BufferedAmount() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *DataChannel) BufferedAmountLowThreshold() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *DataChannel) SetBufferedAmountLowThreshold(th uint64) { _ = "STUB: not implemented"; return }

func (d *DataChannel) OnBufferedAmountLow(f func()) { _ = "STUB: not implemented"; return }

func (d *DataChannel) makeBufferedAmountLowHandler(f func()) func() {
	_ = "STUB: not implemented"
	return nil
}

func (d *DataChannel) getStatsID() string { _ = "STUB: not implemented"; return "" }

func (d *DataChannel) collectStats(collector *statsReportCollector) {
	_ = "STUB: not implemented"
	return
}

func (d *DataChannel) setReadyState(r DataChannelState) { _ = "STUB: not implemented"; return }
