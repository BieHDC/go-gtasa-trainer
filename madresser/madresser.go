package madresser

import (
	"unsafe"
)

// Generic Pointer to unintptr
func AddressOf[Type any](t *Type) uintptr {
	return uintptr(unsafe.Pointer(t))
}

// The only way to make a type of another type
func DupeType[Type any](t Type) Type {
	var zero Type
	return zero
}

// Sometimes we have addresses that directly store the target struct
// but since it was not allocated by go, we have to do a little trickery
// by making a local reference to it, so we can properly access it
func TypeAtAbsolute[Type comparable](p uintptr) *Type {
	// take our absolute pointer
	ptr := unsafe.Pointer(p)
	// make a reference to it
	localstorage := &ptr
	// then return the address
	return (*Type)(*localstorage)
}
