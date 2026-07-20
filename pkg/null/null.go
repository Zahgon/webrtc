package null

type Bool struct {
	Valid bool
	Bool  bool
}

func NewBool(value bool) Bool { _ = "STUB: not implemented"; return *new(Bool) }

type Byte struct {
	Valid bool
	Byte  byte
}

func NewByte(value byte) Byte { _ = "STUB: not implemented"; return *new(Byte) }

type Complex128 struct {
	Valid      bool
	Complex128 complex128
}

func NewComplex128(value complex128) Complex128 { _ = "STUB: not implemented"; return *new(Complex128) }

type Complex64 struct {
	Valid     bool
	Complex64 complex64
}

func NewComplex64(value complex64) Complex64 { _ = "STUB: not implemented"; return *new(Complex64) }

type Float32 struct {
	Valid   bool
	Float32 float32
}

func NewFloat32(value float32) Float32 { _ = "STUB: not implemented"; return *new(Float32) }

type Float64 struct {
	Valid   bool
	Float64 float64
}

func NewFloat64(value float64) Float64 { _ = "STUB: not implemented"; return *new(Float64) }

type Int struct {
	Valid bool
	Int   int
}

func NewInt(value int) Int { _ = "STUB: not implemented"; return *new(Int) }

type Int16 struct {
	Valid bool
	Int16 int16
}

func NewInt16(value int16) Int16 { _ = "STUB: not implemented"; return *new(Int16) }

type Int32 struct {
	Valid bool
	Int32 int32
}

func NewInt32(value int32) Int32 { _ = "STUB: not implemented"; return *new(Int32) }

type Int64 struct {
	Valid bool
	Int64 int64
}

func NewInt64(value int64) Int64 { _ = "STUB: not implemented"; return *new(Int64) }

type Int8 struct {
	Valid bool
	Int8  int8
}

func NewInt8(value int8) Int8 { _ = "STUB: not implemented"; return *new(Int8) }

type Rune struct {
	Valid bool
	Rune  rune
}

func NewRune(value rune) Rune { _ = "STUB: not implemented"; return *new(Rune) }

type String struct {
	Valid  bool
	String string
}

func NewString(value string) String { _ = "STUB: not implemented"; return *new(String) }

type Uint struct {
	Valid bool
	Uint  uint
}

func NewUint(value uint) Uint { _ = "STUB: not implemented"; return *new(Uint) }

type Uint16 struct {
	Valid  bool
	Uint16 uint16
}

func NewUint16(value uint16) Uint16 { _ = "STUB: not implemented"; return *new(Uint16) }

type Uint32 struct {
	Valid  bool
	Uint32 uint32
}

func NewUint32(value uint32) Uint32 { _ = "STUB: not implemented"; return *new(Uint32) }

type Uint64 struct {
	Valid  bool
	Uint64 uint64
}

func NewUint64(value uint64) Uint64 { _ = "STUB: not implemented"; return *new(Uint64) }

type Uint8 struct {
	Valid bool
	Uint8 uint8
}

func NewUint8(value uint8) Uint8 { _ = "STUB: not implemented"; return *new(Uint8) }
