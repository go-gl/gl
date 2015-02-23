// +build !egl,!noauto

package auto

import "github.com/go-gl/gl/procaddr/darwin"

func init() {
	GetProcAddress = darwin.GetProcAddress
}
