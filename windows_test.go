//go:build windows
// +build windows

package command

import (
	"fmt"
	"os"
	"testing"

	"github.com/Tom5521/CmdRunTools/command"
)

func WriteToLog(data ...any) {
	Strdata := fmt.Sprint(data...)
	os.WriteFile("out.log", []byte(Strdata), os.ModePerm)
}

func Test_PS(t *testing.T) {
	cmd := command.InitCmd("ls")
	cmd.RunWithPS(true)
	err := cmd.Run()
	t.Log(cmd.GetExec().String())
	WriteToLog(cmd.GetExec())
	if err != nil {
		WriteToLog(cmd.GetExec(), "\n", err)
		t.Fail()
	}
}
