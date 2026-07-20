package rtcerr

type UnknownError struct {
	Err error
}

func (e *UnknownError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *UnknownError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type InvalidStateError struct {
	Err error
}

func (e *InvalidStateError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *InvalidStateError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type InvalidAccessError struct {
	Err error
}

func (e *InvalidAccessError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *InvalidAccessError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type NotSupportedError struct {
	Err error
}

func (e *NotSupportedError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *NotSupportedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type InvalidModificationError struct {
	Err error
}

func (e *InvalidModificationError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *InvalidModificationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type SyntaxError struct {
	Err error
}

func (e *SyntaxError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *SyntaxError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type TypeError struct {
	Err error
}

func (e *TypeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *TypeError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type OperationError struct {
	Err error
}

func (e *OperationError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *OperationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type NotReadableError struct {
	Err error
}

func (e *NotReadableError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *NotReadableError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type RangeError struct {
	Err error
}

func (e *RangeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *RangeError) Unwrap() error { _ = "STUB: not implemented"; return nil }
