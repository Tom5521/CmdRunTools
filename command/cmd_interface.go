package command

import "os/exec"

// The universal base interface for the others,
// can be converted unidirectionally to all others.
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

// The exclusive command interface for windows,
// you can convert BaseCmd -> WindowsCmd with specific functions for that OS.
type WindowsCmd interface {
	BaseCmd

	// Configuration

	RunWithPS(bool)

	SetRawExec(bool)
	RawExec() bool
	SetHideCmd(bool)
	HideCmd() bool

	UsingPS() bool
	UsingCmd() bool

	SetPSFlags(string)
	SetPSFlagsf(string, ...any)
	PSFlags() string

	SetCmdFlags(string)
	SetCmdFlagsf(string, ...any)
	CmdFlags() string
}

// An exclusive interface for unix,
// which has the specific functions for that OS.
type UnixCmd interface {
	BaseCmd

	// Configuration

	RunWithShell(bool)
	UseBashShell(bool)
	CustomShell(string, string)
	GetChroot() (string, bool)
	SetChroot(string)
	SetChrootf(string, ...any)
	EnableChroot()
	DisableChroot()
}

// Unix interface with extra functions to cover sudo execution.
type UnixSudoCmd interface {
	UnixCmd

	SetPasswd(string)
}
