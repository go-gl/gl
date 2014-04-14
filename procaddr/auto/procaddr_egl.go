// +build linux,egl

package auto

import "github.com/errcw/glow/procaddr"
import "github.com/errcw/glow/procaddr/egl"

func init() {
	procaddr.GetProcAddress = egl.GetProcAddress
}
