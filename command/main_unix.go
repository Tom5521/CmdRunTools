//go:build unix
// +build unix

package command

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/Tom5521/CmdRunTools/internal"
)

// global struct
type Cmd struct {
	internal.Shared
	Shell struct {
		Enabled  bool
		bash     bool // Default unix shell is sh
		CustomSh struct {
			Enabled bool
			ShName  string
			ShArg   string // Shell execution cmd
		}
	}
	Chroot struct {
		Route   string
		Enabled bool
	}
}

// Init functions

// Creates a Cmd structure
func InitCmd(Command string) Cmd {
	sh := Cmd{}
	sh.Input = Command
	return sh
}

// It's the same as InitCmd(fmt.Sprintf(command,args...))
func InitCmdf(Command string, args ...any) Cmd {
	return InitCmd(fmt.Sprintf(Command, args...))
}

// If the value is true use exec.Command([shell],[arg],input) instead of exec.Command(input[0],input[1:]...)
func (sh *Cmd) RunWithShell(set bool) {
	sh.Shell.Enabled = set
}

// Set a custom stdin,stdout or stderr. Default std is all in false
func (sh *Cmd) CustomStd(Stdin, Stdout, Stderr bool) {
	sh.CStd.Enabled = true
	sh.CStd.Stderr = Stderr
	sh.CStd.Stdout = Stdout
	sh.CStd.Stdin = Stdin
}
func (sh *Cmd) Stdin(set bool) {
	sh.CStd.Enabled = true
	sh.CStd.Stdin = set
}
func (sh *Cmd) Stderr(set bool) {
	sh.CStd.Enabled = true
	sh.CStd.Stderr = set
}
func (sh *Cmd) Stdout(set bool) {
	sh.CStd.Enabled = true
	sh.CStd.Stdout = set
}

// Set a custom shell to exec the command
func (sh *Cmd) CustomShell(Shell_Name, Exec_Arg string) {
	sh.RunWithShell(true)
	sh.Shell.CustomSh.Enabled = true
	sh.Shell.CustomSh.ShArg = Exec_Arg
	sh.Shell.CustomSh.ShName = Shell_Name
}

func (sh *Cmd) UseBashShell(set bool) {
	sh.RunWithShell(true)
	sh.Shell.bash = true
}

// Internal funcions

func (sh Cmd) getExec() *exec.Cmd {
	var cmd *exec.Cmd

	if sh.Chroot.Enabled {
		if sh.Shell.CustomSh.Enabled {
			csh := sh.Shell.CustomSh
			cmd = exec.Command("bash", "-c", csh.ShName, csh.ShArg, sh.Input)
		} else {
			cmd = exec.Command("bash", "-c", sh.Input)
		}
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Chroot: sh.Chroot.Route,
		}
		return cmd

	}

	if sh.Shell.Enabled {
		if sh.Shell.bash {
			cmd = exec.Command("bash", "-c", sh.Input)
		}
		if sh.Shell.CustomSh.Enabled {
			cshell := sh.Shell.CustomSh
			cmd = exec.Command(cshell.ShName, cshell.ShArg, sh.Input)
		}
	} else {
		command := strings.Fields(sh.Input)
		cmd = exec.Command(command[0], command[1:]...)
	}

	return cmd
}

// normal running funcions

// Executes normally the command with the parameters set, with the classic exec.Command(<command>).Run()
func (sh *Cmd) Run() error {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	return cmd.Run()
}

// It is the same as run, but skips the Std setting and returns an error value and the output as a string, i.e. exec.Command(<command>).Output()
func (sh Cmd) Out() (string, error) {
	cmd := sh.getExec()
	out, err := cmd.Output()
	return string(out), err
}

// It is the same as run, but returns one more string value (the output of the command)
func (sh Cmd) CombinedOut() (string, error) {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// Run the command in a new goroutine, just like cmd.Run(), but using exec.Command(<cmd>).Start()
func (sh Cmd) Start() error {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	return cmd.Start()
}

// Returns the exec.Cmd structure with all parameters already configured
func (sh Cmd) GetExec() *exec.Cmd {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	return cmd
}
