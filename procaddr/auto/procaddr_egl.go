// +build egl,!noauto

package auto

import "github.com/go-gl/gl/procaddr/egl"

func init() {
	GetProcAddress = egl.GetProcAddress
}
