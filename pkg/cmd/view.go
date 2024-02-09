package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ucho456job/mycmd/pkg/util"
)

func View() {
	cmdsMap, err := util.GetCmdsMap()
	if err != nil {
		util.PrintErrMsg("failed to get cmd.json file path", err)
		return
	}

	prettyJson, err := json.MarshalIndent(cmdsMap, "", "  ")
	if err != nil {
		util.PrintErrMsg("failed to format json", err)
	}

	fmt.Print(string(prettyJson), "\n")
}
