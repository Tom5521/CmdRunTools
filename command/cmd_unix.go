//go:build unix
// +build unix

package command

import (
	"os"
	"os/exec"
	"strings"
)

// global struct
type Cmd struct {
	Input     string
	Path      string
	path_conf struct {
		enabled bool
		path    string
	}
	Shell struct {
		enabled  bool
		bash     bool // Default linux shell is sh
		CustomSh struct {
			enable bool
			ShName string
			ShArg  string // Shell execution cmd
		}
	}
	CStd struct {
		enable bool
		Stdin  bool
		Stdout bool
		Stderr bool
	}
}

// Init functions

// Runs a normal command (without sudo)
func InitCmd(Command string) Cmd {
	sh := Cmd{Input: Command}
	return sh
}

// General parameter funcions
func (sh *Cmd) SetInput(input string) {
	sh.Input = input
}
func (sh *Cmd) SetPath(path string) {
	sh.path_conf.enabled = true
	sh.path_conf.path = path
	sh.Path = path
}

// If the value is true use exec.Command([shell],[arg],input) instead of exec.Command(input[0],input[1:]...)
func (sh *Cmd) RunWithShell(set bool) {
	sh.Shell.enabled = set
}

// Set a custom stdin,stdout or stderr. Default std is all in false
func (sh *Cmd) CustomStd(Stdin, Stdout, Stderr bool) {
	sh.CStd.enable = true
	sh.CStd.Stderr = Stderr
	sh.CStd.Stdout = Stdout
	sh.CStd.Stdin = Stdin
}
func (sh *Cmd) Stdin(set bool) {
	sh.CStd.enable = true
	sh.CStd.Stdin = set
}
func (sh *Cmd) Stderr(set bool) {
	sh.CStd.enable = true
	sh.CStd.Stderr = set
}
func (sh *Cmd) Stdout(set bool) {
	sh.CStd.enable = true
	sh.CStd.Stdout = set
}

// Set a custom shell to exec the command
func (sh *Cmd) CustomShell(Shell_Name, Exec_Arg string) {
	sh.RunWithShell(true)
	sh.Shell.CustomSh.enable = true
	sh.Shell.CustomSh.ShArg = Exec_Arg
	sh.Shell.CustomSh.ShName = Shell_Name
}

func (sh *Cmd) UseBashShell(set bool) {
	sh.RunWithShell(true)
	sh.Shell.bash = true
}

// Internal funcions

func (sh Cmd) setStd(cmd *exec.Cmd) {
	if sh.CStd.enable {
		std := sh.CStd
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
func (sh Cmd) getExec() *exec.Cmd {
	var cmd *exec.Cmd
	if sh.Shell.enabled {
		if sh.Shell.bash {
			cmd = exec.Command("bash", "-c", sh.Input)
		}
		if sh.Shell.CustomSh.enable {
			cshell := sh.Shell.CustomSh
			cmd = exec.Command(cshell.ShName, cshell.ShArg, sh.Input)
		}
	} else {
		command := strings.Fields(sh.Input)
		cmd = exec.Command(command[0], command[1:]...)
	}
	return cmd
}

// Set and... functions

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

// It is the same as cmd := command.Cmd("<command>"); cmd.CombinedOut() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndCombinedOut(command string) (string, error) {
	sh.SetInput(command)
	return sh.CombinedOut()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.Start() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndStart(command string) error {
	sh.SetInput(command)
	return sh.Start()
}

// normal running funcions

// Executes normally the command with the parameters set, with the classic exec.Command(<command>).Run()
func (sh *Cmd) Run() error {
	cmd := sh.getExec()
	sh.setStd(cmd)
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
	sh.setStd(cmd)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// Run the command in a new goroutine, just like cmd.Run(), but using exec.Command(<cmd>).Start()
func (sh Cmd) Start() error {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd.Start()
}

// Returns the exec.Cmd structure with all parameters already configured
func (sh Cmd) GetExec() *exec.Cmd {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd
}
