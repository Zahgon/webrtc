package rtpdump

import (
	"io"
	"sync"
)

type Reader struct {
	readerMu sync.Mutex
	reader   io.Reader
}

func NewReader(r io.Reader) (*Reader, Header, error) {
	_ = "STUB: not implemented"
	return nil, *new(Header), nil
}

func (r *Reader) Next() (Packet, error) { _ = "STUB: not implemented"; return *new(Packet), nil }
