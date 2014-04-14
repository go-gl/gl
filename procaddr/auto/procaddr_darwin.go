package auto

import "github.com/errcw/glow/procaddr"
import "github.com/errcw/glow/procaddr/darwin"

func init() {
	procaddr.GetProcAddress = darwin.GetProcAddress
}
