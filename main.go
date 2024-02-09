package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ucho456job/mycmd/pkg/cmd"
	"github.com/ucho456job/mycmd/pkg/util"
)

func init() {
	util.InitJson()
}

func main() {
	flagSet := flag.NewFlagSet("mycmd", flag.ExitOnError)
	groupFlag := flagSet.String("g", "", "Specify the group")
	taskFlag := flagSet.String("t", "", "Specify the task")

	if len(os.Args) < 2 {
		err := fmt.Errorf("insufficient arguments: expected 1, but got %d", len(os.Args)-1)
		util.PrintErrMsg("failed to execute mycmd", err)
		return
	}

	arg := os.Args[1]
	args := os.Args[2:]
	flagSet.Parse(args)
	switch arg {
	case "exec":
		cmd.Exec(*groupFlag, *taskFlag)
	default:
		err := errors.New("unknown arg")
		util.PrintErrMsg("failed to execute mycmd", err)
		return
	}
}
