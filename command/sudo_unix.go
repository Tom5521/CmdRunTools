//go:build unix
// +build unix

package command

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/Tom5521/CmdRunTools/internal"
)

type SudoCmd struct {
	Cmd
	sudoPars struct {
		getted bool
		Passwd string
	}
}

// Sudo parameters.
func (sh *SudoCmd) SetPasswd(password string) {
	sh.sudoPars.getted = true
	sh.sudoPars.Passwd = password
}

// Internal sudo functions

func (sh SudoCmd) getExec() *exec.Cmd {
	var cmd *exec.Cmd
	if sh.Shell.Enabled {
		if sh.Shell.bash {
			command := strings.Fields(fmt.Sprintf("bash -c \"sudo -S %v\"", sh.Input))
			cmd = exec.Command(command[0], command[1:]...)
		}
		if sh.Shell.CustomSh.Enabled {
			cshell := sh.Shell.CustomSh
			command := strings.Fields(fmt.Sprintf("sudo -S %v %v %v", cshell.ShName, cshell.ShArg, sh.Input))
			cmd = exec.Command(command[0], command[1:]...)
		}
	} else {
		command := strings.Fields("sudo -S " + sh.Input)
		cmd = exec.Command(command[0], command[1:]...)
	}
	return cmd
}

func (sh SudoCmd) writePasswd(cmd *exec.Cmd) error {
	stdin, _ := cmd.StdinPipe()
	var reterr error
	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, sh.sudoPars.Passwd)
		if err != nil {
			reterr = err
		}
	}()
	return reterr
}

// sudo running funcions

func (sh *SudoCmd) GetExec() *exec.Cmd {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	return cmd
}

func (sh *SudoCmd) Run() error {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	err := sh.writePasswd(cmd)
	if err != nil {
		return err
	}
	return cmd.Run()
}

func (sh SudoCmd) Out() (string, error) {
	cmd := sh.getExec()
	err := sh.writePasswd(cmd)
	if err != nil {
		return "", err
	}
	out, err := cmd.Output()
	return string(out), err
}

func (sh SudoCmd) CombinedOut() (string, error) {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	err := sh.writePasswd(cmd)
	if err != nil {
		return "", err
	}
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func (sh SudoCmd) Start() error {
	cmd := sh.getExec()
	internal.SetStd(sh.Shared, cmd)
	err := sh.writePasswd(cmd)
	if err != nil {
		return err
	}
	return cmd.Start()
}
