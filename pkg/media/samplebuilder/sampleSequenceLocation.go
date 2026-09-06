package samplebuilder

type sampleSequenceLocation struct {
	head uint16

	tail uint16
}

func (l sampleSequenceLocation) empty() bool { _ = "STUB: not implemented"; return false }

func (l sampleSequenceLocation) hasData() bool { _ = "STUB: not implemented"; return false }

func (l sampleSequenceLocation) count() uint16 { _ = "STUB: not implemented"; return 0 }

const (
	slCompareVoid = iota
	slCompareBefore
	slCompareInside
	slCompareAfter
)

func (l sampleSequenceLocation) compare(pos uint16) int { _ = "STUB: not implemented"; return 0 }
