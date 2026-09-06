package webrtc

type BundlePolicy int

const (
	BundlePolicyUnknown BundlePolicy = iota

	BundlePolicyBalanced

	BundlePolicyMaxCompat

	BundlePolicyMaxBundle
)

const (
	bundlePolicyBalancedStr  = "balanced"
	bundlePolicyMaxCompatStr = "max-compat"
	bundlePolicyMaxBundleStr = "max-bundle"
)

func newBundlePolicy(raw string) BundlePolicy { _ = "STUB: not implemented"; return *new(BundlePolicy) }

func (t BundlePolicy) String() string { _ = "STUB: not implemented"; return "" }

func (t *BundlePolicy) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (t BundlePolicy) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
