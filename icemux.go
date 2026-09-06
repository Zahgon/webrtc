package webrtc

import (
	"net"

	"github.com/pion/ice/v4"
	"github.com/pion/logging"
)

func NewICETCPMux(logger logging.LeveledLogger, listener net.Listener, readBufferSize int) ice.TCPMux {
	_ = "STUB: not implemented"
	return *new(ice.TCPMux)
}

func NewICEUDPMux(logger logging.LeveledLogger, udpConn net.PacketConn) ice.UDPMux {
	_ = "STUB: not implemented"
	return *new(ice.UDPMux)
}
