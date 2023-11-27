//go:build unix
// +build unix

package command

import (
	"fmt"
)

// General parameter funcions

func (sh *Cmd) SetChroot(mountPoint string) {
	sh.Chroot.Enabled = true
	sh.Chroot.Route = mountPoint
}

func (sh *Cmd) SetChrootf(mountPoint string, args ...any) {
	sh.Chroot.Enabled = true
	sh.Chroot.Route = fmt.Sprintf(mountPoint, args...)
}
