// +build !egl,!noauto

package auto

import "github.com/errcw/glow/procaddr/wgl"

func init() {
	GetProcAddress = wgl.GetProcAddress
}
