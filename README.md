# gl [![Build Status](https://github.com/go-gl/gl/actions/workflows/main/badge.svg)](https://travis-ci.org/go-gl/gl) [![Go Reference](https://pkg.go.dev/badge/github.com/go-gl/gl.svg)](https://pkg.go.dev/github.com/go-gl/gl)

This repository holds Go bindings to various OpenGL versions. They are auto-generated using [Glow](https://github.com/go-gl/glow).

Features:
- Go functions that mirror the C specification using Go types.
- Support for multiple OpenGL APIs (GL/GLES/EGL/WGL/GLX/EGL), versions, and profiles.
- Support for extensions (including debug callbacks).

Requirements:
- A cgo compiler (typically gcc).
- On Ubuntu/Debian-based systems, the `libgl1-mesa-dev` package.

Usage
-----

Use `go get -u` to download and install the prebuilt packages. The prebuilt packages support OpenGL versions 2.1, 3.1, 3.2, 3.3, 4.1, 4.2, 4.3, 4.4, 4.5, 4.6 across both the core and compatibility profiles and include all extensions. Pick whichever one(s) you need:

    go get -u github.com/go-gl/gl/v{3.2,3.3,4.1,4.2,4.3,4.4,4.5,4.6}-{core,compatibility}/gl
    go get -u github.com/go-gl/gl/v3.1/gles2
    go get -u github.com/go-gl/gl/v2.1/gl

Once the bindings are installed you can use them with the appropriate import statements.

```Go
import "github.com/go-gl/gl/v3.3-core/gl"

func main() {
	window := ... // Open a window.
	window.MakeContextCurrent()

	// Important! Call gl.Init only under the presence of an active OpenGL context,
	// i.e., after MakeContextCurrent.
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
}
```

The `gl` package contains the OpenGL functions and enumeration values for the imported version. It also contains helper functions for working with the API. Of note is `gl.Ptr` which takes a Go array or slice or pointer and returns a corresponding `uintptr` to use with functions expecting data pointers. Also of note is `gl.Str` which takes a null-terminated Go string and returns a corresponding `*int8` to use with functions expecting character pointers.

A note about threading and goroutines. The bindings do not expose a mechanism to make an OpenGL context current on a different thread so you must restrict your usage to the thread on which you called `gl.Init()`. To do so you should use [LockOSThread](https://code.google.com/p/go-wiki/wiki/LockOSThread).

Examples
--------

Examples illustrating how to use the bindings are available in the [example](https://github.com/go-gl/example) repo. There are examples for [OpenGL 4.1 core](https://github.com/go-gl/example/tree/master/gl41core-cube) and [OpenGL 2.1](https://github.com/go-gl/example/tree/master/gl21-cube).

Function Loading
----------------

The `procaddr` package contains platform-specific functions for [loading OpenGL functions](https://www.opengl.org/wiki/Load_OpenGL_Functions). Calling `gl.Init()` uses the `auto` subpackage to automatically select an appropriate implementation based on the build environment. If you want to select a specific implementation you can use the `noauto` build tag and the `gl.InitWithProcAddrFunc` initialization function.

Generating
----------

These gl bindings are generated using the [Glow](https://github.com/go-gl/glow) generator. Only developers of this repository need to do this step.

It is required to have `glow` source in a sibling directory to `go-gl/gl` since relative paths are used for generation (see `generate.go`).
For non-module-aware cases, this means `glow` needs to be in the same Go workspace as `go-gl/gl`.
For module-aware cases, `go-gl/glow` needs to be checked out parallel to `go-gl/gl`. 

In either case, the `glow` binary must be in your `$PATH`. Doable with `go get -u github.com/go-gl/glow` if your `$GOPATH/bin` is in your `$PATH`.

Perform generation with the following:

```bash
cd path/to/go-gl/gl
go generate -tags=gen .
```

More information about these bindings can be found in the [Glow repository](https://github.com/go-gl/glow).
