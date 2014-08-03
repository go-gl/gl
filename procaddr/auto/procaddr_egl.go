// +build egl,!noauto

package auto

import "github.com/go-gl/glow/procaddr/egl"

func init() {
	GetProcAddress = egl.GetProcAddress
}
