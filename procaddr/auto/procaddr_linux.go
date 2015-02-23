// +build !egl,!noauto

package auto

import "github.com/go-gl/gl/procaddr/glx"

func init() {
	GetProcAddress = glx.GetProcAddress
}
