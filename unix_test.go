//go:build unix
// +build unix

package command

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Tom5521/CmdRunTools/command"
)

var conf = getTestConf()

type jsondata struct {
	Passwd string `json:"password"`
}

func WriteLog(toAdd string) {
	file, _ := os.Create("out.log")
	file.WriteString(toAdd)
}

func getTestConf() jsondata {
	data := jsondata{Passwd: "-"}
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
	var LogOut string
	defer WriteLog(LogOut)
	t.Log(conf.Passwd)
	ls := func() {
		cmd := command.InitCmd("ls /")
		out, _ := cmd.Out()
		t.Log(out)
		LogOut += "\n" + out
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
	_, err = cmd.SetAndCombinedOut("ls")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func Test_SetAnd(t *testing.T) {
	cmd := command.Cmd{}
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
