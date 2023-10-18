//go:build windows
// +build windows

package cmd

import (
	"testing"

	win "github.com/Tom5521/CmdRunTools/windows"
)

func TestPS1(t *testing.T) {
	cmd := win.PSCmd("mkdir test")
	cmd.Stdout(true)
	err := cmd.Run()
	t.Log("TTEst")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
