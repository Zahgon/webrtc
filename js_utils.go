//go:build js && wasm
// +build js,wasm

package webrtc

import (
	"syscall/js"
)

func awaitPromise(promise js.Value) (js.Value, error) {
	_ = "STUB: not implemented"
	return *new(js.Value), nil
}

func valueToUint16Pointer(val js.Value) *uint16 { _ = "STUB: not implemented"; return nil }

func valueToStringPointer(val js.Value) *string { _ = "STUB: not implemented"; return nil }

func stringToValueOrUndefined(val string) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func uint8ToValueOrUndefined(val uint8) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func interfaceToValueOrUndefined(val any) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func valueToStringOrZero(val js.Value) string { _ = "STUB: not implemented"; return "" }

func valueToUint8OrZero(val js.Value) uint8 { _ = "STUB: not implemented"; return 0 }

func valueToUint16OrZero(val js.Value) uint16 { _ = "STUB: not implemented"; return 0 }

func valueToUint32OrZero(val js.Value) uint32 { _ = "STUB: not implemented"; return 0 }

func valueToStrings(val js.Value) []string { _ = "STUB: not implemented"; return nil }

func valueToBoolOrFalse(val js.Value) bool { _ = "STUB: not implemented"; return false }

func valueToBoolPointer(val js.Value) *bool { _ = "STUB: not implemented"; return nil }

func stringPointerToValue(val *string) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func uint16PointerToValue(val *uint16) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func boolToValueOrUndefined(val bool) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func boolPointerToValue(val *bool) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func stringsToValue(strings []string) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

func stringEnumToValueOrUndefined(s string) js.Value {
	_ = "STUB: not implemented"
	return *new(js.Value)
}

func recoveryToError(e any) error { _ = "STUB: not implemented"; return nil }

func uint8ArrayValueToBytes(val js.Value) []byte { _ = "STUB: not implemented"; return nil }
