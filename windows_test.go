package cmd_test

import (
	"testing"

	"github.com/Tom5521/CmdRunTools/command"
)

func Test_PS(t *testing.T) {
	cmd := command.Cmd("ls")
	cmd.RunWithPS(true)
	cmd.Run()
}
