//go:build windows
// +build windows

package command

import win "github.com/Tom5521/CmdRunTools/internal/windows"

func Cmd(command string) *win.WinCmd {
	ret := win.Cmd(command)
	return &ret
}

func GetStruct() *win.WinCmd {
	return GetStruct()
}
