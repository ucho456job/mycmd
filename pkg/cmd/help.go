package cmd

import (
	"fmt"

	"github.com/ucho456job/mycmd/pkg/util"
)

func Help() {
	msg := ""
	msg += blue_("Usage:")
	msg += white("  mycmd [Command] [Option]")
	msg += "\n"
	msg += blue_("Commands:")
	msg += white("  clip  Copy commands to the clipboard. '-g' and '-t' options can be used.")
	msg += white("  edit  Edit cmd.json. '-e' option can be used.")
	msg += white("  exec  Execute commands. '-g' and '-t' options can be used.")
	msg += white("  help  Show help.")
	msg += white("  read  Show commands. '-g' and '-t' options can be used.")
	msg += white("  view  Show cmd.json.")
	msg += "\n"
	msg += blue_("Options:")
	msg += white("  -g  Specify group.")
	msg += white("  -t  Specify task.")
	msg += white("  -e  Specify editor. path must exist.")
	msg += "\n"
	msg += blue_("How to use:")
	msg += white("  1. Run `mycmd edit` to open $HOME/.mycmd/cmd.json")
	msg += white("  2. Please update cmd.json referring to \"About cmd.json configuration structure\"")
	msg += white("  3. You can run registered commands by running `mycmd exec` and following the prompts.")
	msg += "\n"
	msg += blue_("About cmd.json configuration structure:")
	msg += white("  To customize the execution of tasks within our tool,")
	msg += white("  you need to modify the cmd.json configuration file according to the following structure.")
	msg += "\n"
	msg += green("  Sample cmd.json:")
	msg += "\n"
	msg += `    {
      "group1": {
        "task1": [
          "command1",
          "command2"
        ],
        "task2": [
          "command1"
        ]
      },
      "group2": {
        "task1": [
          "command1",
          "command2"
        ]
      }
    }`
	msg += "\n\n"
	msg += green("  Structure Explanation:")
	msg += white("    group   : Represents a collection of tasks.")
	msg += white("              Each group should have a unique name and contain one or more tasks.")
	msg += white("    task    : Identifies an individual task within a group.")
	msg += white("              Each task is represented by a key and is associated with an array of commands to be executed.")
	msg += white("    commands: A list of string values, each representing a command to be executed as part of the task.")
	msg += white("              Commands within the task are executed in the order they are listed.")
	msg += "\n"
	msg += green("  Editing Guidelines:")
	msg += white("    Defining a New Group    : To define a new group of tasks, add a new key at the root level of cmd.json.")
	msg += white("                              Use a descriptive name for your group to easily identify the set of tasks it contains.")
	msg += white("    Adding a New Task       : Within a group, add a new key to represent the task.")
	msg += white("                              The key should be descriptive of the task's purpose.")
	msg += white("                              Associate this key with an array of commands that defines the task's execution logic.")
	msg += white("    Specifying Commands     : For each task, specify the commands to be executed in an array format.")
	msg += white("                              Ensure each command is a string and listed in the order you wish them to be executed.")
	msg += white("    Removing Tasks or Groups: To remove a task, delete its key from the group.")
	msg += white("                              If a group no longer contains any tasks, consider removing the group as well to keep the configuration clean and concise.")
	msg += "\n"
	msg += blue_("About dynamic argument:")
	msg += white("  The tool supports dynamic argument input for commands specified in cmd.json.")
	msg += white("  This feature allows for interactive execution of commands by enabling the user to input arguments at runtime.")
	msg += white("  However, only one argument can be dynamically changed.")
	msg += "\n"
	msg += green("  Sample cmd.json:")
	msg += "\n"
	msg += `    {
      "Linux": {
        "list": [
          "ls __arg__"
        ]
      }
    }`
	msg += "\n\n"
	msg += green("  How It Works:")
	msg += white("    You can specify a placeholder for runtime arguments using '__arg__'.")
	msg += white("    This placeholder will be replaced with the user's input right before the command is executed.")
	msg += "\n"
	msg += green("  Example:")
	msg += white("    For the sample cmd.json, run 'mycmd exec -g Linux -t list'.")
	msg += white("    This placeholder will be replaced with the user's input right before the command is executed.")

	fmt.Println(msg)
}

func white(msg string) string {
	return msg + "\n"
}

func blue_(msg string) string {
	return util.BlueFont(msg) + "\n"
}

func green(msg string) string {
	return util.GreenFont(msg) + "\n"
}
