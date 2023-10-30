package cmd_test

import (
	"testing"

	win "github.com/Tom5521/CmdRunTools/internal/windows"
)

func Test_PS(t *testing.T) {
	cmd := win.Cmd("ls")
	cmd.RunWithPS(true)
	cmd.Run()
}
