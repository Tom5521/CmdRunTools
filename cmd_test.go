package cmd

import (
	"testing"

	win "github.com/Tom5521/CmdRunTools/windows"
)

func TestCmd1(t *testing.T) {
	cmd := win.Cmd("dir")
	cmd.Stdout(true)
	cmd.Run()
	cmd.SetInput("ls")
	cmd.RunWithPS(true)
	cmd.Run()
}
