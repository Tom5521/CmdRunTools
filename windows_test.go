//go:build windows
// +build windows

package command

import (
	"testing"

	"github.com/Tom5521/CmdRunTools/command"
)

func Test_PS(t *testing.T) {
	cmd := command.InitCmd("ls")
	cmd.RunWithPS(true)
	cmd.Run()
}
