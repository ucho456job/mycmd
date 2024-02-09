package util

import (
	"errors"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func InputPrompt(label string) (input string, err error) {
	prompt := promptui.Prompt{
		Label: BlueFont(label),
	}

	input, err = prompt.Run()
	if err != nil {
		return "", err
	}

	return input, nil
}

func SelectGroupPrompt(cmdsMap CmdsMap) (group string, err error) {
	groups := make([]string, 0, len(cmdsMap))
	for g := range cmdsMap {
		groups = append(groups, g)
	}

	prompt := promptui.Select{
		Label: BlueFont("Please select group"),
		Items: groups,
	}

	_, group, err = prompt.Run()
	if err != nil {
		return "", err
	}

	return group, nil
}

func SelectTaskPrompt(groupMap GroupMap) (task string, err error) {
	tasks := make([]string, 0, len(groupMap))
	for task := range groupMap {
		tasks = append(tasks, task)
	}

	prompt := promptui.Select{
		Label: BlueFont("Please select task"),
		Items: tasks,
	}

	_, task, err = prompt.Run()
	if err != nil {
		return "", err
	}

	return task, nil
}

func ConfirmExecCmdsPrompt(cmds []string) (shouldProceed bool, err error) {
	msg := "Selected commands are below:"
	for idx, cmd := range cmds {
		msg += fmt.Sprintf("\nStep %d: %s", idx+1, cmd)
	}
	fmt.Println(msg)

	prompt := promptui.Prompt{
		Label:   BlueFont("Execute these commands?") + " [Y/n]",
		Default: "",
		Validate: func(input string) error {
			if input == "" || strings.ToUpper(input) == "Y" || strings.ToUpper(input) == "N" {
				return nil
			}
			return errors.New("input must be 'y' or 'n'")
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	if result == "" {
		result = "Y"
	}

	shouldProceed = strings.ToUpper(result) == "Y"
	return shouldProceed, nil
}
