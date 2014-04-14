package procaddr

import "unsafe"

type GetProcAddressFunc func(name string) unsafe.Pointer

// Global function for loading OpenGL procedure addresses.
var GetProcAddress GetProcAddressFunc
