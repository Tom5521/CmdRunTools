//go:build unix
// +build unix

package command

import "fmt"

// NewUnixCmd but with a fmt.Sprintf in it.
func NewUnixCmdf(command string, args ...any) UnixCmd {
	return NewUnixCmd(fmt.Sprintf(command, args...))
}

// Creates a UnixCmd interface that is convertible to:
// UnixCmd -> UnixSudoCmd.
func NewUnixCmd(command string) UnixCmd {
	return NewCmd(command)
}

// Creates a UnixCmd interface but with specific functions for executing sudo commands.
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
