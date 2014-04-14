package egl

// #cgo linux LDFLAGS: -lEGL
// #include <stdlib.h>
// #include <EGL/egl.h>
import "C"
import "unsafe"

func GetProcAddress(name string) unsafe.Pointer {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return unsafe.Pointer(C.eglGetProcAddress(cname))
}
