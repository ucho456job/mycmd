package cmd

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/ucho456job/mycmd/pkg/util"
)

func Edit(editor string) {
	jsonFilePath, err := util.GetJsonFilePath()
	if err != nil {
		util.PrintErrMsg("failed to get cmd.json file path", err)
		return
	}

	if editor == "" {
		if runtime.GOOS == "windows" {
			editor = "notepad"
		} else {
			editor = "vi"
		}
	}

	execCmd := exec.Command(editor, jsonFilePath)
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	if err = execCmd.Run(); err != nil {
		util.PrintErrMsg("failed to execute editor", err)
	}
}
