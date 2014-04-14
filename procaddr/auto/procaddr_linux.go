// +build !egl

package auto

import "github.com/errcw/glow/procaddr"
import "github.com/errcw/glow/procaddr/glx"

func init() {
	procaddr.GetProcAddress = glx.GetProcAddress
}
