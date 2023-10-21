package cmd

import (
	"testing"

	"github.com/Tom5521/CmdRunTools/unix"
)

func Test_Run(t *testing.T) {
	var cmd unix.UnixCmd
	out, err := cmd.SetAndOut("mkdir test")
	t.Log(out, err)
	if err != nil {
		t.Fail()
	}
	cmd.SetAndRun("rmdir test")
}
