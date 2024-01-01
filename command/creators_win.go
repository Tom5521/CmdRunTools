//go:build windows
// +build windows

package command

func NewWindowsCmd(command string) WindowsCmd {
	return NewCmd(command)
}

func NewWindowsCmdf(command string, args ...any) WindowsCmd {
	return NewCmdf(command, args...)
}
