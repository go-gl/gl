// +build !egl,!noauto

package auto

import "github.com/errcw/glow/procaddr/glx"

func init() {
	GetProcAddress = glx.GetProcAddress
}
