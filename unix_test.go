package cmd_test

import (
	"os"
	"testing"

	"github.com/Tom5521/CmdRunTools/command"
	"github.com/Tom5521/CmdRunTools/internal/unix"
)

func Test_Sudo(t *testing.T) {
	ls := func() {
		cmd := unix.Cmd("ls /")
		cmd.Stdout(true)
		cmd.Run()
	}
	file := "/asdadass"
	cmd := unix.Sudo_Cmd("mkdir "+file, "4142")
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
	var cmd unix.UnixCmd
	out, err := cmd.SetAndOut("mkdir test")
	t.Log(out, err)
	if err != nil {
		t.Fail()
	}
	cmd.SetAndRun("rmdir test")
}

func Test_Unix_Struct(t *testing.T) {
	t.Log(unix.GetStruct())
}
