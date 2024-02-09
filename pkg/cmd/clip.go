package cmd

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/ucho456job/mycmd/pkg/util"
)

func Clip(group string, task string) {
	cmds, err := util.PrepareCmds(&group, &task)
	if err != nil {
		util.PrintErrMsg("failed to prepare commands", err)
		return
	}

	cmdsStr := strings.Join(cmds, "\n")
	if err := clipboard.WriteAll(cmdsStr); err != nil {
		util.PrintErrMsg("failed to copy to the clipboard", err)
		return
	}

	util.PrintSuccessMsg(fmt.Sprintf("Copied commands to clipboard:\n%s", cmdsStr))
}
