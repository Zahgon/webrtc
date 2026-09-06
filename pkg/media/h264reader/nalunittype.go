package h264reader

type NalUnitType uint8

const (
	NalUnitTypeUnspecified              NalUnitType = 0
	NalUnitTypeCodedSliceNonIdr         NalUnitType = 1
	NalUnitTypeCodedSliceDataPartitionA NalUnitType = 2
	NalUnitTypeCodedSliceDataPartitionB NalUnitType = 3
	NalUnitTypeCodedSliceDataPartitionC NalUnitType = 4
	NalUnitTypeCodedSliceIdr            NalUnitType = 5
	NalUnitTypeSEI                      NalUnitType = 6
	NalUnitTypeSPS                      NalUnitType = 7
	NalUnitTypePPS                      NalUnitType = 8
	NalUnitTypeAUD                      NalUnitType = 9
	NalUnitTypeEndOfSequence            NalUnitType = 10
	NalUnitTypeEndOfStream              NalUnitType = 11
	NalUnitTypeFiller                   NalUnitType = 12
	NalUnitTypeSpsExt                   NalUnitType = 13
	NalUnitTypeCodedSliceAux            NalUnitType = 19
)

func (n *NalUnitType) String() string {
	_ = "STUB: not implemented" //nolint:cyclop
	return ""
}
