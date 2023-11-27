//go:build windows
// +build windows

package command

import "fmt"

// It sets the customized powershell flags, its syntax when executed would be something like this "powershell.exe [flags] /c [command]".
func (sh *Cmd) CustomPSFlags(flags string) {
	sh.Powershell.PSFlags = flags
}

func (sh *Cmd) CustomPSFlagsf(flags string, args ...any) {
	sh.Powershell.PSFlags = fmt.Sprintf(flags, args...)
}

// It sets the customized cmd flags, its syntax when executed would be something like this "cmd.exe [flags] /c [command]".
func (sh *Cmd) CustomCmdFlags(flags string) {
	sh.Cmd.CmdFlags = flags
}

func (sh *Cmd) CustomCmdFlagsf(flags string, args ...any) {
	sh.Cmd.CmdFlags = fmt.Sprintf(flags, args...)
}
