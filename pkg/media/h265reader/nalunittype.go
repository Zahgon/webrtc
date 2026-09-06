package h265reader

type NalUnitType uint8

const (
	NalUnitTypeTrailN   NalUnitType = 0
	NalUnitTypeTrailR   NalUnitType = 1
	NalUnitTypeTsaN     NalUnitType = 2
	NalUnitTypeTsaR     NalUnitType = 3
	NalUnitTypeStsaN    NalUnitType = 4
	NalUnitTypeStsaR    NalUnitType = 5
	NalUnitTypeRadlN    NalUnitType = 6
	NalUnitTypeRadlR    NalUnitType = 7
	NalUnitTypeRaslN    NalUnitType = 8
	NalUnitTypeRaslR    NalUnitType = 9
	NalUnitTypeBlaWLp   NalUnitType = 16
	NalUnitTypeBlaWRadl NalUnitType = 17
	NalUnitTypeBlaNLp   NalUnitType = 18
	NalUnitTypeIdrWRadl NalUnitType = 19
	NalUnitTypeIdrNLp   NalUnitType = 20
	NalUnitTypeCraNut   NalUnitType = 21

	NalUnitTypeVps       NalUnitType = 32
	NalUnitTypeSps       NalUnitType = 33
	NalUnitTypePps       NalUnitType = 34
	NalUnitTypeAud       NalUnitType = 35
	NalUnitTypeEos       NalUnitType = 36
	NalUnitTypeEob       NalUnitType = 37
	NalUnitTypeFd        NalUnitType = 38
	NalUnitTypePrefixSei NalUnitType = 39
	NalUnitTypeSuffixSei NalUnitType = 40

	NalUnitTypeReserved41 NalUnitType = 41
	NalUnitTypeReserved47 NalUnitType = 47
	NalUnitTypeUnspec48   NalUnitType = 48
	NalUnitTypeUnspec63   NalUnitType = 63
)

func (n *NalUnitType) String() string {
	_ = "STUB: not implemented" //nolint:cyclop
	return ""
}
