//go:build js && wasm
// +build js,wasm

package webrtc

import (
	"syscall/js"

	"github.com/pion/datachannel"
)

const dataChannelBufferSize = 16384

type DataChannel struct {
	underlying js.Value

	onOpenHandler       *js.Func
	onCloseHandler      *js.Func
	onClosingHandler    *js.Func
	onMessageHandler    *js.Func
	onBufferedAmountLow *js.Func
	onErrorHandler      *js.Func

	onCloseFunc func()

	closeWrapperInstalled bool

	api *API
}

func (d *DataChannel) JSValue() js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func (d *DataChannel) OnOpen(f func()) { _ = "STUB: not implemented"; return }

func (d *DataChannel) OnClose(f func()) { _ = "STUB: not implemented"; return }

func (d *DataChannel) OnClosing(f func()) { _ = "STUB: not implemented"; return }

func (d *DataChannel) OnError(f func(err error)) { _ = "STUB: not implemented"; return }

func (d *DataChannel) OnMessage(f func(msg DataChannelMessage)) { _ = "STUB: not implemented"; return }

func (d *DataChannel) Send(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) SendText(s string) (err error) { _ = "STUB: not implemented"; return nil }

func (d *DataChannel) Detach() (datachannel.ReadWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(datachannel.ReadWriteCloser), nil
}

func (d *DataChannel) Close() (err error) { _ = "STUB: not implemented"; return nil }

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

func valueToDataChannelMessage(val js.Value) DataChannelMessage {
	_ = "STUB: not implemented"
	return *new(DataChannelMessage)
}
