package command

import "fmt"

// It is the same as cmd := command.Cmd("<command>"); cmd.Run() but in a single argument,
// what it does is to put an internal input (the one provided)
// and execute it directly without affecting the main structure.
func (sh Cmd) SetAndRun(command string) error {
	sh.SetInput(command)
	return sh.Run()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.Out() but in a single argument,
// what it does is to put an internal input (the one provided)
// and execute it directly without affecting the main structure.
func (sh Cmd) SetAndOut(command string) (string, error) {
	sh.SetInput(command)
	return sh.Out()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.CombinedOut() but in a single argument,
// what it does is to put an internal input (the one provided)
// and execute it directly without affecting the main structure.
func (sh Cmd) SetAndCombinedOut(command string) (string, error) {
	sh.SetInput(command)
	return sh.CombinedOut()
}

// It is the same as cmd := command.Cmd("<command>"); cmd.Start() but in a single argument,
// what it does is to put an internal input (the one provided)
// and execute it directly without affecting the main structure.
func (sh Cmd) SetAndStart(command string) error {
	sh.SetInput(command)
	return sh.Start()
}

// Set andf... functions

// Is the same as a SetAndRun(fmt.Sprintf(command,args)) but shortened for better handling.
func (sh Cmd) SetAndRunf(command string, args ...any) error {
	return sh.SetAndRun(fmt.Sprintf(command, args...))
}

// Is the same as a SetAndOut(fmt.Sprintf(command,args)) but shortened for better handling.
func (sh Cmd) SetAndOutf(command string, args ...any) (string, error) {
	return sh.SetAndOut(fmt.Sprintf(command, args...))
}

// It is the same as a SetAndCombinedOut(fmt.Sprintf(command,args)) but shortened for better handling.
func (sh Cmd) SetAndCombinedOutf(command string, args ...any) (string, error) {
	return sh.SetAndCombinedOut(fmt.Sprintf(command, args...))
}

// It is the same as a SetAndStart(fmt.Sprintf(command,args)) but shortened for better handling.
func (sh Cmd) SetAndStartf(command string, args ...any) error {
	return sh.SetAndStart(fmt.Sprintf(command, args...))
}
