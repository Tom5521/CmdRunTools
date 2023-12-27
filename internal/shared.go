package internal

import (
	"fmt"
	"os"
	"os/exec"
)

type Shared struct {
	Input    string
	Path     string
	PathConf struct {
		Enabled bool
		Path    string // It's the same as Cmd.Path
	}

	CStd struct {
		Enabled bool
		Stdin   bool
		Stderr  bool
		Stdout  bool
	}
}

func (sh *Shared) SetInput(input string) {
	sh.Input = input
}

func (sh *Shared) SetInputf(input string, args ...any) {
	sh.Input = fmt.Sprintf(input, args...)
}

func (sh *Shared) SetPath(path string) {
	sh.PathConf.Enabled = true
	sh.PathConf.Path = path
	sh.Path = path
}

func (sh *Shared) SetPathf(path string, args ...any) {
	pathset := fmt.Sprintf(path, args...)
	sh.PathConf.Enabled = true
	sh.PathConf.Path = pathset
	sh.Path = pathset
}

// Set custom Stdin,Stdout,Stderr in one function.
func (sh *Shared) Std(stdin, stdout, stderr bool) {
	sh.CStd.Enabled = true
	sh.CStd.Stderr = stderr
	sh.CStd.Stdin = stdin
	sh.CStd.Stdout = stdout
}

// Set the individual Stdin.
func (sh *Shared) Stdin(set bool) {
	sh.CStd.Enabled = true
	sh.CStd.Stdin = set
}

// Set the individual Stdout.
func (sh *Shared) Stdout(set bool) {
	sh.CStd.Enabled = true
	sh.CStd.Stdout = set
}

// Set the individual Stderr.
func (sh *Shared) Stderr(set bool) {
	sh.CStd.Enabled = true
	sh.CStd.Stderr = set
}

func SetStd(sh Shared, cmd *exec.Cmd) {
	if !sh.CStd.Enabled {
		return
	}
	cStd := sh.CStd
	if cStd.Stdin {
		cmd.Stdin = os.Stdin
	}
	if cStd.Stdout {
		cmd.Stdout = os.Stdout
	}
	if cStd.Stderr {
		cmd.Stderr = os.Stderr
	}
}
