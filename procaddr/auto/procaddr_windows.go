// +build !egl,!noauto

package auto

import "github.com/go-gl/glow/procaddr/wgl"

func init() {
	GetProcAddress = wgl.GetProcAddress
}
