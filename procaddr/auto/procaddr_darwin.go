// +build !egl,!noauto

package auto

import "github.com/errcw/glow/procaddr/darwin"

func init() {
	GetProcAddress = darwin.GetProcAddress
}
