//go:build js && wasm
// +build js,wasm

package webrtc

type detachedDataChannel struct {
	dc *DataChannel

	read chan DataChannelMessage
	done chan struct{}
}

func newDetachedDataChannel(dc *DataChannel) *detachedDataChannel {
	_ = "STUB: not implemented"
	return nil
}

func (c *detachedDataChannel) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *detachedDataChannel) ReadDataChannel(p []byte) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func (c *detachedDataChannel) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *detachedDataChannel) WriteDataChannel(p []byte, isString bool) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *detachedDataChannel) Close() error { _ = "STUB: not implemented"; return nil }
