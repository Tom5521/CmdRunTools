//go:build windows
// +build windows

package command

import (
	"fmt"
	"os"
	"testing"

	"github.com/Tom5521/CmdRunTools/command"
	"github.com/stretchr/testify/assert"
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

func Test_Cmd(t *testing.T) {
	cmd := command.InitCmd("dir")
	out, err := cmd.CombinedOut()
	WriteToLog(out)
	if err != nil {
		WriteToLog(err.Error(), out)
		t.Fail()
	}
	cmd.SetInput("ls")
	cmd.RunWithPS(true)
	//cmd.RunWithoutCmd(true)
	err = cmd.Run()
	if err != nil {
		WriteToLog(err.Error())
		t.Fail()
	}
}

func Test_ChangeInput(t *testing.T) {
	assert := assert.New(t)
	cmd := command.InitCmd("dir")
	original := cmd.Input
	fmt.Println(cmd.Input)
	cmd.SetInput("Hola mundo")
	n := cmd.Input
	assert.NotEqual(n, original)
}
