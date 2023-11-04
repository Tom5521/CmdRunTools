//go:build unix
// +build unix

package command

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/Tom5521/CmdRunTools/command"
)

var conf = getTestConf()

type jsondata struct {
	Passwd        string `json:"password"`
	ChrootDir     string `json:"chroot-dir"`
	ChrootCommand string `json:"chroot-command"`
}

func WriteLog(data ...any) {
	Strdata := fmt.Sprint(data...)
	os.WriteFile("out.log", []byte(Strdata), os.ModePerm)
}

func getTestConf() jsondata {
	data := jsondata{}
	var confFile = "Testconf.json"
	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		file, _ := os.Create(confFile)
		out, _ := json.Marshal(data)
		file.Write(out)
	}
	file, _ := os.ReadFile(confFile)
	json.Unmarshal(file, &data)
	return data
}

func Test_Sudo(t *testing.T) {
	if conf.Passwd == "" {
		t.Log("Passwd is <nil>")
		t.Fatal()
	}
	var LogOut string
	defer WriteLog(LogOut)
	t.Log(conf.Passwd)
	ls := func() {
		cmd := command.InitCmd("ls /")
		out, _ := cmd.Out()
		t.Log(out)
		LogOut = fmt.Sprintf("%v\n%v", LogOut, out)
	}
	file := "/asdadass"
	cmd := command.Sudo_Cmd("mkdir "+file, conf.Passwd)
	err := cmd.Run()
	ls()
	_, checkfile := os.Stat(file)
	if os.IsNotExist(checkfile) {
		t.Fail()
	}
	cmd.SetInput("rm -r " + file)
	err1 := cmd.Run()
	ls()

	if err1 != nil || err != nil {
		LogOut += "\n" + err1.Error()
		t.Fail()
	}
}
func Test_CmdLib(t *testing.T) {
	cmd := command.InitCmd("ls /")
	cmd.UseBashShell(true)
	out, err := cmd.Out()
	t.Log(string(out))
	if err != nil {
		t.Fail()
	}
	//cmd.Stdout(true)
	out, err = cmd.SetAndCombinedOut("ls")
	WriteLog(out)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func Test_SetAnd(t *testing.T) {
	cmd := command.Cmd{}
	_, err := cmd.SetAndOut("mkdir test")
	out, err := cmd.SetAndCombinedOut("ls")
	WriteLog(out)

	if _, err := os.Stat("test"); os.IsNotExist(err) {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
	cmd.SetAndRun("rmdir test")
}

// I test this in a virtual machine
func Test_Chroot(t *testing.T) {
	cmd := command.InitCmd(conf.ChrootCommand)
	cmd.SetChroot(conf.ChrootDir)
	t.Log(cmd.Chroot)
	t.Log(cmd.GetExec())
	//cmd.CustomStd(true, true, true)
	out, err := cmd.Out()
	WriteLog(out)
	if err != nil {
		t.Log(err.Error())
		outLog, _ := os.ReadFile("out.log")
		t.Log(string(outLog))
		t.Fail()
		return
	}
	outLog, _ := os.ReadFile("out.log")
	t.Log(string(outLog))

}
