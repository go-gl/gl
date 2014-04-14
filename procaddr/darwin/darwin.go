package darwin

// #cgo darwin LDFLAGS: -framework OpenGL
// #include <stdlib.h>
// #include <dlfcn.h>
// void* GetProcAddress(const char* name) {
// 	return dlsym(RTLD_DEFAULT, name);
// }
import "C"
import "unsafe"

func GetProcAddress(name string) unsafe.Pointer {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return unsafe.Pointer(C.GetProcAddress(cname))
}
