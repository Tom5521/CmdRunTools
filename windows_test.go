package cmd_test

import (
	"testing"

	"github.com/Tom5521/CmdRunTools/internal/win"
)

func Test_PS(t *testing.T) {
	cmd := win.Cmd("ls")
	cmd.RunWithPS(true)
	cmd.Run()
}
