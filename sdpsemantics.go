package webrtc

type SDPSemantics int

const (
	SDPSemanticsUnifiedPlan SDPSemantics = iota

	SDPSemanticsPlanB

	SDPSemanticsUnifiedPlanWithFallback
)

const (
	sdpSemanticsUnifiedPlanWithFallback = "unified-plan-with-fallback"
	sdpSemanticsUnifiedPlan             = "unified-plan"
	sdpSemanticsPlanB                   = "plan-b"
)

func newSDPSemantics(raw string) SDPSemantics { _ = "STUB: not implemented"; return *new(SDPSemantics) }

func (s SDPSemantics) String() string { _ = "STUB: not implemented"; return "" }

func (s *SDPSemantics) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s SDPSemantics) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
