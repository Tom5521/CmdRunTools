package cmd

import (
	"testing"

	"github.com/Tom5521/CmdRunTools/unix"
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
	ls()
	cmd.SetInput("rm -r " + file)
	err1 := cmd.Run()
	ls()

	if err1 != nil || err != nil {
		t.Fail()
	}

}
