package cmd

import (
	"fmt"

	"github.com/ucho456job/mycmd/pkg/util"
)

func Read(group string, task string) {
	cmds, err := util.PrepareCmds(&group, &task)
	if err != nil {
		util.PrintErrMsg("failed to prepare commands", err)
		return
	}

	msg := "It is registered with the following contents:\n"
	msg += "Group is `" + group + "`\n"
	msg += "Task  is `" + task + "`\n"
	for idx, cmd := range cmds {
		msg += fmt.Sprintf("Command for Step %d is `%s`\n", idx+1, cmd)
	}
	util.PrintSuccessMsg(msg)
}
