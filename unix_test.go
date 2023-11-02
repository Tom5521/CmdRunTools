package cmd_test

import (
	"os"
	"testing"

	"github.com/Tom5521/CmdRunTools/command"
)

func Test_Sudo(t *testing.T) {
	ls := func() {
		cmd := command.Cmd("ls /")
		cmd.Stdout(true)
		cmd.Run()
	}
	file := "/asdadass"
	cmd := command.Sudo_Cmd("mkdir "+file, "4142")
	err := cmd.Run()
	_, checkfile := os.Stat(file)
	if os.IsNotExist(checkfile) {
		t.Fail()
	}
	ls()
	cmd.SetInput("rm -r " + file)
	err1 := cmd.Run()
	ls()

	if err1 != nil || err != nil {
		t.Fail()
	}
}
func Test_CmdLib(t *testing.T) {
	cmd := command.Cmd("ls /")
	cmd.UseBashShell(true)
	out, err := cmd.Out()
	t.Log(string(out))
	if err != nil {
		t.Fail()
	}
	cmd.Stdout(true)
	_, err = cmd.SetAndCombinedOut("ls SD")
	if err != nil {
		t.Fail()
	}
}

func Test_SetAnd(t *testing.T) {
	var cmd command.UnixCmd
	out, err := cmd.SetAndOut("mkdir test")
	if _, err := os.Stat("test"); os.IsNotExist(err) {
		t.Fail()
	}
	t.Log(out, err)
	if err != nil {
		t.Fail()
	}
	cmd.SetAndRun("rmdir test")
}
