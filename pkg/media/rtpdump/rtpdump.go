package rtpdump

import (
	"errors"
	"net"
	"time"
)

const (
	pktHeaderLen = 8
	headerLen    = 16
	preambleLen  = 36
)

var errMalformed = errors.New("malformed rtpdump")

type Header struct {
	Start time.Time

	Source net.IP

	Port uint16
}

func (h Header) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // G115
//nolint:gosec // G115

func (h *Header) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

type Packet struct {
	Offset time.Duration

	IsRTCP bool

	Payload []byte
}

func (p Packet) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // G115
//nolint:gosec // G115

func (p *Packet) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func (p *Packet) offsetMs() uint32 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // G115

type packetHeader struct {
	Length uint16

	PacketLength uint16

	Offset uint32
}

func (p packetHeader) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *packetHeader) Unmarshal(d []byte) error { _ = "STUB: not implemented"; return nil }

func (p packetHeader) offset() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }
