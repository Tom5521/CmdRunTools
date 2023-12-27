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

// Creates a UnixCmd interface but with specific functions for executing sudo commands.
func NewUnixSudoCmd(command string, optionalPassword ...string) UnixSudoCmd {
	return NewSudoCmd(command, optionalPassword...)
}

// NewUnixSudoCmd but with a fmt.Sprintf in it.
func NewUnixSudoCmdf(passwd, command string, args ...any) UnixSudoCmd {
	return NewSudoCmdf(passwd, command, args...)
}

// NewUnixCmd but with a fmt.Sprintf in it.
func NewUnixCmdf(command string, args ...any) UnixCmd {
	return NewUnixCmd(fmt.Sprintf(command, args...))
}

// Creates a UnixCmd interface that is convertible to:
// UnixCmd -> UnixSudoCmd.
func NewUnixCmd(command string) UnixCmd {
	return NewCmd(command)
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

// Create a sudo command.
func NewSudoCmd(command string, optionalPassword ...string) *SudoCmd {
	sudoSh := &SudoCmd{}
	sudoSh.Input = command
	if len(optionalPassword) > 0 {
		sudoSh.SetPasswd(optionalPassword[0])
	}
	return sudoSh
}

// NewSudoCmd but with a fmt.Sprintf in it.
func NewSudoCmdf(passwd, command string, args ...any) *SudoCmd {
	return NewSudoCmd(fmt.Sprintf(command, args...), passwd)
}
