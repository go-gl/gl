// +build !egl,!noauto

package auto

import "github.com/go-gl/glow/procaddr/glx"

func init() {
	GetProcAddress = glx.GetProcAddress
}
