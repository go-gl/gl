// Package procaddr defines platform-specific mechanisms for loading OpenGL
// function addresses.
package procaddr

import "unsafe"

// A function that returns a C function pointer to the specified function.
type GetProcAddressFunc func(name string) unsafe.Pointer
