GL
==

This repository holds Go bindings to various OpenGL versions. They are auto-generated using [Glow](github.com/go-gl/glow).

Features:
- Go functions that mirror the C specification using Go types.
- Support for multiple OpenGL APIs (GL/GLES/EGL/WGL/GLX/EGL), versions, and profiles.
- Support for extensions (including debug callbacks).

Usage
-----

Use `go get` to download and install one of the prebuilt packages. The prebuilt packages support OpenGL versions 3.2, 3.3, 4.1, 4.4 and 4.5 across both the core and compatibility profiles and include all extensions.

    go get github.com/go-gl/gl/v{3.2,3.3,4.1,4.4,4.5}-{core,compatibility}/gl
    go get github.com/go-gl/gl/v3.3-core/gl

Once the bindings are installed you can use them with the appropriate import statements.

```Go
import "github.com/go-gl/gl/v3.3-core/gl"

func main() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
}
```

The `gl` package contains the OpenGL functions and enumeration values for the imported version. It also contains helper functions for working with the API. Of note is `gl.Ptr` which takes a Go array or slice or pointer and returns a corresponding `uintptr` to use with functions expecting data pointers. Also of note is `gl.Str` which takes a null-terminated Go string and returns a corresponding `*int8` to use with functions expecting character pointers.

A note about threading and goroutines. The bindings do not expose a mechanism to make an OpenGL context current on a different thread so you must restrict your usage to the thread on which you called `gl.Init()`. To do so you should use [LockOSThread](https://code.google.com/p/go-wiki/wiki/LockOSThread).

Examples
--------

Examples illustrating how to use the bindings are available in the [examples](https://github.com/go-gl/examples) repo. There are examples for [OpenGL 4.1 core](https://github.com/go-gl/examples/blob/master/glfw31-gl41core-cube) and [OpenGL 2.1](https://github.com/go-gl/examples/tree/master/glfw31-gl21-cube).

Function Loading
----------------

The `procaddr` package contains platform-specific functions for [loading OpenGL functions](https://www.opengl.org/wiki/Load_OpenGL_Functions). Calling `gl.Init()` uses the `auto` subpackage to automatically select an appropriate implementation based on the build environment. If you want to select a specific implementation you can use the `noauto` build tag and the `gl.InitWithProcAddrFunc` initialization function.

Generating
----------

These gl bindings are generated using `glow` generator.

```bash
go generate -tags=gen github.com/go-gl/gl
```

More information about these bindings can be found in the [Glow repository](https://github.com/go-gl/glow).
