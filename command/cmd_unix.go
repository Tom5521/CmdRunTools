//go:build unix
// +build unix

package command

import (
	"os"
	"os/exec"
	"strings"
)

// global struct
type UnixCmd struct {
	input string
	path  struct {
		Enabled bool
		Path    string
	}
	runWithShell struct {
		Enabled  bool
		bash     bool // Default linux shell is sh
		customSh struct {
			Enable bool
			ShName string
			ShArg  string // Shell execution cmd
		}
	}
	customStd struct {
		Enable bool
		Stdin  bool
		Stdout bool
		Stderr bool
	}
}

// Init functions

// Runs a normal command (without sudo)
func Cmd(Command string) UnixCmd {
	sh := UnixCmd{input: Command}
	return sh
}

// General parameter funcions
func (sh *UnixCmd) SetInput(input string) {
	sh.input = input
}
func (sh *UnixCmd) SetPath(path string) {
	sh.path.Enabled = true
	sh.path.Path = path
}

// If the value is true use exec.Command([shell],[arg],input) instead of exec.Command(input[0],input[1:]...)
func (sh *UnixCmd) RunWithShell(set bool) {
	sh.runWithShell.Enabled = set
}

// Set a custom stdin,stdout or stderr. Default std is all in false
func (sh *UnixCmd) CustomStd(Stdin, Stdout, Stderr bool) {
	sh.customStd.Enable = true
	sh.customStd.Stderr = Stderr
	sh.customStd.Stdout = Stdout
	sh.customStd.Stdin = Stdin
}
func (sh *UnixCmd) Stdin(set bool) {
	sh.customStd.Enable = true
	sh.customStd.Stdin = set
}
func (sh *UnixCmd) Stderr(set bool) {
	sh.customStd.Enable = true
	sh.customStd.Stderr = set
}
func (sh *UnixCmd) Stdout(set bool) {
	sh.customStd.Enable = true
	sh.customStd.Stdout = set
}

// Set a custom shell to exec the command
func (sh *UnixCmd) CustomShell(Shell_Name, Exec_Arg string) {
	sh.RunWithShell(true)
	sh.runWithShell.customSh.Enable = true
	sh.runWithShell.customSh.ShArg = Exec_Arg
	sh.runWithShell.customSh.ShName = Shell_Name
}

func (sh *UnixCmd) UseBashShell(set bool) {
	sh.RunWithShell(true)
	sh.runWithShell.bash = true
}

// Internal funcions

func (sh UnixCmd) setStd(cmd *exec.Cmd) {
	if sh.customStd.Enable {
		std := sh.customStd
		if std.Stderr {
			cmd.Stderr = os.Stderr
		}
		if std.Stdout {
			cmd.Stdout = os.Stdout
		}
		if std.Stdin {
			cmd.Stdin = os.Stdin
		}
	}
}
func (sh UnixCmd) getExec() *exec.Cmd {
	var cmd *exec.Cmd
	if sh.runWithShell.Enabled {
		if sh.runWithShell.bash {
			cmd = exec.Command("bash", "-c", sh.input)
		}
		if sh.runWithShell.customSh.Enable {
			cshell := sh.runWithShell.customSh
			cmd = exec.Command(cshell.ShName, cshell.ShArg, sh.input)
		}
	} else {
		command := strings.Fields(sh.input)
		cmd = exec.Command(command[0], command[1:]...)
	}
	return cmd
}

// Set and... functions

// It is the same as cmd := command.Cmd("<command>"); cmd.Run() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh UnixCmd) SetAndRun(command string) error {
	sh.SetInput(command)
	return sh.Run()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.Out() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh UnixCmd) SetAndOut(command string) (string, error) {
	sh.SetInput(command)
	return sh.Out()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.CombinedOut() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh UnixCmd) SetAndCombinedOut(command string) (string, error) {
	sh.SetInput(command)
	return sh.CombinedOut()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.Start() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh UnixCmd) SetAndStart(command string) error {
	sh.SetInput(command)
	return sh.Start()
}

// normal running funcions

// Executes normally the command with the parameters set, with the classic exec.Command(<command>).Run()
func (sh *UnixCmd) Run() error {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd.Run()
}

// It is the same as run, but skips the Std setting and returns an error value and the output as a string, i.e. exec.Command(<command>).Output()
func (sh UnixCmd) Out() (string, error) {
	cmd := sh.getExec()
	out, err := cmd.Output()
	return string(out), err
}

// It is the same as run, but returns one more string value (the output of the command)
func (sh UnixCmd) CombinedOut() (string, error) {
	cmd := sh.getExec()
	sh.setStd(cmd)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// Run the command in a new goroutine, just like cmd.Run(), but using exec.Command(<cmd>).Start()
func (sh UnixCmd) Start() error {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd.Start()
}

// Returns the exec.Cmd structure with all parameters already configured
func (sh UnixCmd) GetExec() *exec.Cmd {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd
}