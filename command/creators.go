package command

import "fmt"

// Creates a new BaseCmd interface,
// which is completely universal and cross-platform.
func NewBaseCmd(command string) BaseCmd {
	return NewCmd(command)
}

// NewBaseCmd but with a fmt.Sprintf in it.
func NewBaseCmdf(command string, args ...any) BaseCmd {
	return NewBaseCmd(fmt.Sprintf(command, args...))
}

// Create a cmd structure convertible to the default interface for your platform.
// Ex:
// Windows -> WindowsCmd.
// Unix -> UnixCmd.
func NewCmd(command string) *Cmd {
	sh := &Cmd{}
	sh.Input = command
	return sh
}

// It's the same as NewCmd(fmt.Sprintf(command,args...)).
func NewCmdf(command string, args ...any) *Cmd {
	return NewCmd(fmt.Sprintf(command, args...))
}
