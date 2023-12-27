//go:build windows
// +build windows

package command

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/Tom5521/CmdRunTools/internal"
)

type Cmd struct {
	internal.Shared
	Powershell struct {
		PSFlags string
		Enabled bool
	}
	Cmd struct {
		CmdFlags      string
		RunWithoutCmd bool
		HideCmdWindow bool
	}
}

// Global config parameters

// Run the command using "powershell.exe [parameters] /c [command]" instead of "cmd.exe [parameters] /c [command]".
func (sh *Cmd) RunWithPS(set bool) {
	sh.Powershell.Enabled = set
}

// Execute the command directly, it is useful if you want to execute a binary, this mode does not have access to the path so you will have to put the full path of the binary or use something relative to execute it.
func (sh *Cmd) RunWithoutCmd(set bool) {
	sh.Cmd.RunWithoutCmd = set
}

// Hides the cmd/powershell window that appears when executing a command in go.
func (sh *Cmd) HideCmdWindow(set bool) {
	sh.Cmd.HideCmdWindow = set
}

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

// Internal functions

func (sh Cmd) getFinal() *exec.Cmd {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	return cmd
}

func (sh Cmd) formatcmd() string {
	var cmd string
	if sh.Powershell.Enabled {
		cmd = fmt.Sprintf("powershell.exe %v /c %v", sh.Powershell.PSFlags, sh.Input)
	} else if !sh.Cmd.RunWithoutCmd {
		cmd = fmt.Sprintf("cmd.exe %v /c %v", sh.Cmd.CmdFlags, sh.Input)
	} else {
		cmd = sh.Input
	}
	return cmd
}

//Functions that break down the linter for quick commenting

func (sh Cmd) getExec() *exec.Cmd {
	command := strings.Fields(sh.formatcmd())
	cmd := exec.Command(command[0], command[1:]...)
	if sh.Cmd.HideCmdWindow {
		cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
	}
	return cmd
}

// Running functions

// Execute the command with all the parameters already set, something like "exec.Command([formatted command]).Run()" and return its error output.
func (sh Cmd) Run() error {
	return sh.getFinal().Run()
}

// Execute the command with all the parameters already set, something like "exec.Command([formatted command]).Start()" and return its error output.
func (sh Cmd) Start() error {
	return sh.getFinal().Run()
}

// Execute the command with all the parameters already set, something like "exec.Command([formatted command]).Output()" and return its string and error output.
func (sh Cmd) Out() (string, error) {
	cmd := sh.getExec()
	out, err := cmd.Output()
	return string(out), err
}

// Execute the command with all parameters already set, something like "exec.Command([formatted command]).CombinedOutput()" and return its error and string output, as well as executing with output to stdin,stdout and stderr.
func (sh Cmd) CombinedOut() (string, error) {
	out, err := sh.getFinal().CombinedOutput()
	return string(out), err
}

// Returns the exec.Cmd structure with all parameters already configured.
func (sh Cmd) GetExec() *exec.Cmd {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	return cmd
}
