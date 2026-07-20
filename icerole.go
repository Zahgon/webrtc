package webrtc

type ICERole int

const (
	ICERoleUnknown ICERole = iota

	ICERoleControlling

	ICERoleControlled
)

const (
	iceRoleControllingStr = "controlling"
	iceRoleControlledStr  = "controlled"
)

func newICERole(raw string) ICERole { _ = "STUB: not implemented"; return *new(ICERole) }

func (t ICERole) String() string { _ = "STUB: not implemented"; return "" }

func (t ICERole) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *ICERole) UnmarshalText(b []byte) error { _ = "STUB: not implemented"; return nil }
