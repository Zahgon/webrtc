package webrtc

import (
	"github.com/pion/ice/v4"
)

func supportedNetworkTypes() []NetworkType { _ = "STUB: not implemented"; return nil }

type NetworkType int

const (
	NetworkTypeUnknown NetworkType = iota

	NetworkTypeUDP4

	NetworkTypeUDP6

	NetworkTypeTCP4

	NetworkTypeTCP6
)

const (
	networkTypeUDP4Str = "udp4"
	networkTypeUDP6Str = "udp6"
	networkTypeTCP4Str = "tcp4"
	networkTypeTCP6Str = "tcp6"
)

func (t NetworkType) String() string { _ = "STUB: not implemented"; return "" }

func (t NetworkType) Protocol() string {
	_ = "STUB: not implemented" //nolint:staticcheck
	return ""
}

func NewNetworkType(raw string) (NetworkType, error) {
	_ = "STUB: not implemented"
	return *new(NetworkType), nil
}

func getNetworkType(iceNetworkType ice.NetworkType) (NetworkType, error) {
	_ = "STUB: not implemented"
	return *new(NetworkType), nil
}

func toICENetworkTypes(networkTypes []NetworkType) []ice.NetworkType {
	_ = "STUB: not implemented"
	return nil
}

func (networkType NetworkType) toICE() ice.NetworkType {
	_ = "STUB: not implemented"
	return *new(ice.NetworkType)
}
