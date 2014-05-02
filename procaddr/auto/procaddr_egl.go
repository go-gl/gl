// +build egl,!noauto

package auto

import "github.com/errcw/glow/procaddr/egl"

func init() {
	GetProcAddress = egl.GetProcAddress
}
