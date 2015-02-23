// +build !egl,!noauto

package auto

import "github.com/go-gl/gl/procaddr/wgl"

func init() {
	GetProcAddress = wgl.GetProcAddress
}
