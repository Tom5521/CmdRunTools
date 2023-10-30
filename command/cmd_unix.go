//go:build unix
// +build unix

package command

import (
	"github.com/Tom5521/CmdRunTools/internal/unix"
)

func GetStruct() *unix.UnixCmd {
	return unix.GetStruct()
}

func GetSudoStruct() *unix.UnixSudoCmd {
	return unix.GetSudoStruct()
}

func Cmd(command string) *unix.UnixCmd {
	ret := unix.Cmd(command)
	return &ret
}

func SudoCmd(command string, password ...string) *unix.UnixSudoCmd {
	ret := unix.Sudo_Cmd(command)
	if len(password) > 0 {
		ret.SetSudoPasswd(password[0])
	}
	return &ret

}
