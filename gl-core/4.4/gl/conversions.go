// Glow automatically generated OpenGL binding: http://github.com/errcw/glow
package gl
import (
	"C"
	"fmt"
	"log"
	"reflect"
	"strings"
	"unsafe"
)
// Ptr takes a pointer, slice, or array and returns its GL-compatible address.
func Ptr(data interface{}) uintptr {
	if data == nil {
		return uintptr(0)
	}
	v := reflect.ValueOf(data)
	switch v.Type().Kind() {
	case reflect.Ptr:
		e := v.Elem()
		switch e.Kind() {
		case
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return e.UnsafeAddr()
		}
	case reflect.Uintptr:
		return v.Pointer()
	case reflect.Slice:
		return v.Index(0).UnsafeAddr()
	case reflect.Array:
		return v.UnsafeAddr()
	}
	panic(fmt.Sprintf("Unsupproted type %s; must be a pointer, slice, or array", v.Type()))
}
// Str takes a null-terminated Go string and returns its GL-compatible address.
// This function reaches into Go string storage in an unsafe way so the caller
// must ensure the string is not garbage collected.
func Str(str string) *int8 {
	if !strings.HasSuffix(str, "\x00") {
		log.Fatal("str argument missing null terminator", str)
	}
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*int8)(unsafe.Pointer(header.Data))
}
// GoStr takes a null-terminated string returned by OpenGL and constructs a
// corresponding Go string.
func GoStr(cstr *uint8) string {
	return C.GoString((*C.char)(unsafe.Pointer(cstr)))
}
