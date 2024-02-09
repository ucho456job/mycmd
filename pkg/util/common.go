package util

import (
	"fmt"
)

func PrepareCmds(group *string, task *string) (cmds Cmds, err error) {
	cmdsMap, err := GetCmdsMap()
	if err != nil {
		return nil, fmt.Errorf("failed to load cmd.json: %w", err)
	}
	if len(cmdsMap) == 0 {
		return nil, fmt.Errorf("cmd.json is empty: please run `mycmd edit` and update cmd.json")
	}

	groupMap, exists := cmdsMap[*group]
	if !exists {
		if *group != "" {
			PrintWarnMsg(fmt.Sprintf("%s isn't defined", *group))
		}
		selectedGroup, err := SelectGroupPrompt(cmdsMap)
		if err != nil {
			return nil, fmt.Errorf("failed to select group: %w", err)
		}
		*group = selectedGroup
		groupMap = cmdsMap[*group]
	}

	cmds, exists = groupMap[*task]
	if !exists {
		if *task != "" {
			PrintWarnMsg(fmt.Sprintf("%s isn't in the %s", *task, *group))
		}
		selectedTask, err := SelectTaskPrompt(groupMap)
		if err != nil {
			return nil, fmt.Errorf("failed to select task: %w", err)
		}
		*task = selectedTask
		cmds = groupMap[*task]
	}

	return cmds, nil
}
