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
	sudo_pars struct {
		getted bool
		Passwd string
	}
}

// Runs a command as sudo.
func NewSudoCmd(command string, optional_password ...string) UnixCmd {
	sudoSh := &SudoCmd{}
	sudoSh.Input = command
	if len(optional_password) > 0 {
		sudoSh.SetSudoPasswd(optional_password[0])
	}
	return sudoSh
}
func NewSudoCmdf(passwd, command string, args ...any) UnixCmd {
	return NewSudoCmd(fmt.Sprintf(command, args...), passwd)
}

// Sudo parameters.
func (sh *SudoCmd) SetSudoPasswd(password string) {
	sh.sudo_pars.getted = true
	sh.sudo_pars.Passwd = password
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
		_, err := io.WriteString(stdin, sh.sudo_pars.Passwd)
		if err != nil {
			reterr = err
		}
	}()
	return reterr
}

// sudo running funcions

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
