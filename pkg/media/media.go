package media

import (
	"time"

	"github.com/pion/rtp"
)

type Sample struct {
	Data               []byte
	Timestamp          time.Time
	Duration           time.Duration
	PacketTimestamp    uint32
	PrevDroppedPackets uint16
	Metadata           any

	RTPHeaders []*rtp.Header
}

type Writer interface {
	WriteRTP(packet *rtp.Packet) error

	Close() error
}
