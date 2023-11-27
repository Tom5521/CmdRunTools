//go:build windows
// +build windows

package command

import "fmt"

// Set and ... functions

// It is the same as cmd := command.Cmd("<command>"); cmd.Run() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndRun(command string) error {
	sh.SetInput(command)
	return sh.Run()
}

func (sh Cmd) SetAndRunf(command string, args ...any) error {
	return sh.SetAndRun(fmt.Sprintf(command, args...))
}

// It is the same as cmd := command.Cmd("<command>"); cmd.Out() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndOut(command string) (string, error) {
	sh.SetInput(command)
	return sh.Out()
}

func (sh Cmd) SetAndOutf(command string, args ...any) (string, error) {
	return sh.SetAndOut(fmt.Sprintf(command, args...))
}

// It is the same as cmd := command.Cmd("<command>"); cmd.Start() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndStart(command string) error {
	sh.SetInput(command)
	return sh.Start()
}

func (sh Cmd) SetAndStartf(command string, args ...any) error {
	return sh.SetAndStart(fmt.Sprintf(command, args...))
}

// It is the same as cmd := command.Cmd("<command>"); cmd.CombinedOut() but in a single argument, what it does is to put an internal input (the one provided) and execute it directly without affecting the main structure.
func (sh Cmd) SetAndCombinedOut(command string) (string, error) {
	sh.SetInput(command)
	return sh.CombinedOut()
}

func (sh Cmd) SetAndCombinedOutf(command string, args ...any) (string, error) {
	return sh.SetAndCombinedOut(fmt.Sprintf(command, args...))
}
