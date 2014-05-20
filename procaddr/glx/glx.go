package glx

// #cgo linux LDFLAGS: -lGL
// #include <stdlib.h>
// #include <GL/glx.h>
import "C"
import "unsafe"

func GetProcAddress(name string) unsafe.Pointer {
	var cname *C.GLubyte = (*C.GLubyte)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(cname))
	return unsafe.Pointer(C.glXGetProcAddress(cname))
}
