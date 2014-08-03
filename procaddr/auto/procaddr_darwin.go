// +build !egl,!noauto

package auto

import "github.com/go-gl/glow/procaddr/darwin"

func init() {
	GetProcAddress = darwin.GetProcAddress
}
