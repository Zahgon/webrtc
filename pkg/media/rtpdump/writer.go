package rtpdump

import (
	"io"
	"sync"
)

type Writer struct {
	writerMu sync.Mutex
	writer   io.Writer
}

func NewWriter(w io.Writer, hdr Header) (*Writer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Writer) WritePacket(p Packet) error { _ = "STUB: not implemented"; return nil }
