package auto

import "github.com/errcw/glow/procaddr"
import "github.com/errcw/glow/procaddr/wgl"

func init() {
	procaddr.GetProcAddress = wgl.GetProcAddress
}
