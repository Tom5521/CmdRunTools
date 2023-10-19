//go:build windows
// +build windows

package cmd

import win "github.com/Tom5521/CmdRunTools/windows"

func Cmd(input string) win.WinCmd {
	cmd := win.Cmd(input)
	return cmd
}
