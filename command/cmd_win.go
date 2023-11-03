//go:build windows
// +build windows

package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type Cmd struct {
	powershell struct {
		psFlags           string
		runWithPowershell bool
	}
	cmd struct {
		cmdFlags      string
		runWithoutCmd bool
		hideCmdWindow bool
	}
	input string
	path  struct {
		enabled bool
		path    string
	}
	customStd struct {
		enabled bool
		stdin   bool
		stderr  bool
		stdout  bool
	}
}

// Init

// Initializes a new instance of the command, already setting the command to run
func InitCmd(input string) Cmd {
	sh := Cmd{}
	sh.SetInput(input)
	return sh
}

// Global config parameters

// Run the command using "powershell.exe [parameters] /c [command]" instead of "cmd.exe [parameters] /c [command]"
func (sh *Cmd) RunWithPS(set bool) {
	sh.powershell.runWithPowershell = set
}

// Set the running path of the command
func (sh *Cmd) SetPath(path string) {
	sh.path.enabled = true
	sh.path.path = path
}

// Execute the command directly, it is useful if you want to execute a binary, this mode does not have access to the path so you will have to put the full path of the binary or use something relative to execute it.
func (sh *Cmd) RunWithoutCmd(set bool) {
	sh.cmd.runWithoutCmd = set
}

// Set the command to be executed
func (sh *Cmd) SetInput(input string) {
	sh.input = input
}

// Hides the cmd/powershell window that appears when executing a command in go.
func (sh *Cmd) HideCmdWindow(set bool) {
	sh.cmd.hideCmdWindow = set
}

// It sets the customized powershell flags, its syntax when executed would be something like this "powershell.exe [flags] /c [command]".
func (sh *Cmd) CustomPSFlags(flags string) {
	sh.powershell.psFlags = flags
}

// It sets the customized cmd flags, its syntax when executed would be something like this "cmd.exe [flags] /c [command]".
func (sh *Cmd) CustomCmdFlags(flags string) {
	sh.cmd.cmdFlags = flags
}

// Set custom Stdin,Stdout,Stderr in one function
func (sh *Cmd) CustomStd(Stdin, Stdout, Stderr bool) {
	sh.customStd.enabled = true
	sh.customStd.stderr = Stderr
	sh.customStd.stdin = Stdin
	sh.customStd.stdout = Stdout
}

// Set the individual Stdin
func (sh *Cmd) Stdin(set bool) {
	sh.customStd.enabled = true
	sh.customStd.stdin = set
}

// Set the individual Stdout
func (sh *Cmd) Stdout(set bool) {
	sh.customStd.enabled = true
	sh.customStd.stdout = set
}

// Set the individual Stderr
func (sh *Cmd) Stderr(set bool) {
	sh.customStd.enabled = true
	sh.customStd.stderr = set
}

// Internal functions

func (sh Cmd) getFinal() *exec.Cmd {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd
}

func (sh Cmd) setStd(cmd *exec.Cmd) {
	if !sh.customStd.enabled {
		return
	}
	cStd := sh.customStd
	if cStd.stdin {
		cmd.Stdin = os.Stdin
	}
	if cStd.stdout {
		cmd.Stdout = os.Stdout
	}
	if cStd.stderr {
		cmd.Stderr = os.Stderr
	}
}

func (sh Cmd) formatcmd() string {
	var cmd string
	if sh.powershell.runWithPowershell {
		cmd = fmt.Sprintf("powershell.exe %v /c %v", sh.powershell.psFlags, sh.input)
	}
	if !sh.cmd.runWithoutCmd {
		cmd = fmt.Sprintf("cmd.exe %v /c %v", sh.cmd.cmdFlags, sh.input)
	} else {
		cmd = sh.input
	}
	return cmd
}

//Functions that break down the linter for quick commenting

func (sh Cmd) getExec() *exec.Cmd {
	command := strings.Fields(sh.formatcmd())
	cmd := exec.Command(command[0], command[1:]...)
	if sh.cmd.hideCmdWindow {
		cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
	}
	return cmd
}

// Set and ... functions

// It is the same as cmd := command.Cmd("<command>"); cmd.Run() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndRun(command string) error {
	sh.SetInput(command)
	return sh.Run()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.Out() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndOut(command string) (string, error) {
	sh.SetInput(command)
	return sh.Out()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.Start() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndStart(command string) error {
	sh.SetInput(command)
	return sh.Start()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.CombinedOut() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndCombinedOut(command string) (string, error) {
	sh.SetInput(command)
	return sh.CombinedOut()
}

// Running functions

// Execute the command with all the parameters already set, something like "exec.Command([formatted command]).Run()" and return its error output
func (sh Cmd) Run() error {
	return sh.getFinal().Run()
}

// Execute the command with all the parameters already set, something like "exec.Command([formatted command]).Start()" and return its error output
func (sh Cmd) Start() error {
	return sh.getFinal().Run()
}

// Execute the command with all the parameters already set, something like "exec.Command([formatted command]).Output()" and return its string and error output
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

// Returns the exec.Cmd structure with all parameters already configured
func (sh Cmd) GetExec() *exec.Cmd {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd
}
