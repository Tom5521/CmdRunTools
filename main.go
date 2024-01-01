package main

import (
	"fmt"

	"github.com/Tom5521/CmdRunTools/command"
)

func main() {
	cmd := command.NewCmd("ls")
	cmd.RunWithPS(true)
	cmd.Std(true, true, true)
	err := cmd.Run()
	fmt.Println(err)
	fmt.Println(cmd.GetExec())
}
