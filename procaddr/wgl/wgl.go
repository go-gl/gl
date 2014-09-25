package wgl

// #cgo windows LDFLAGS: -lopengl32
// #define WIN32_LEAN_AND_MEAN 1
// #include <windows.h>
// #include <stdlib.h>
// static HMODULE ogl32dll = NULL;
// void* GlowGetProcAddress(const char* name) {
// 	void* pf = wglGetProcAddress((LPCSTR)name);
// 	if(pf) {
// 		return pf;
// 	}
// 	if(ogl32dll == NULL) {
// 		ogl32dll = LoadLibraryA("opengl32.dll");
// 	}
// 	return GetProcAddress(ogl32dll, (LPCSTR)name);
// }
import "C"
import "unsafe"

func GetProcAddress(name string) unsafe.Pointer {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return unsafe.Pointer(C.GlowGetProcAddress(cname))
}
