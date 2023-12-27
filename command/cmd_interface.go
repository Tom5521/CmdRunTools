package command

import "os/exec"

type BaseCmd interface {
	GetExec() *exec.Cmd

	// Configuration

	SetInput(string)
	SetInputf(string, ...any)

	Std(bool, bool, bool)
	Stdout(bool)
	Stdin(bool)
	Stderr(bool)

	// Running

	Run() error
	Out() (string, error)
	Start() error
	CombinedOut() (string, error)

	// Set andf...

	SetAndStartf(string, ...any) error
	SetAndCombinedOutf(string, ...any) (string, error)
	SetAndOutf(string, ...any) (string, error)
	SetAndRunf(string, ...any) error

	// Set andf...

	SetAndStart(string) error
	SetAndCombinedOut(string) (string, error)
	SetAndOut(string) (string, error)
	SetAndRun(string) error
}

type WindowsCmd interface {
	BaseCmd

	// Configuration

	RunWithPS(bool)
	RunWithoutCmd(bool)
	HideCmdWindow(bool)

	CustomPSFlags(string)
	CustomPSFlagsf(string, ...any)

	CustomCmdFlags(string)
	CustomCmdFlagsf(string, ...any)
}

type UnixCmd interface {
	BaseCmd

	// Configuration

	RunWithShell(bool)
	UseBashShell(bool)
	CustomShell(string, string)
	GetChroot() (string, bool)
	SetChroot(string)
	SetChrootf(string, ...any)
}
