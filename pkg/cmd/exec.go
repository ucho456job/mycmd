package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ucho456job/mycmd/pkg/util"
)

func Exec(group string, task string) {
	cmds, err := util.PrepareCmds(&group, &task)
	if err != nil {
		util.PrintErrMsg("failed to prepare commands", err)
		return
	}

	shouldProceed, err := util.ConfirmExecCmdsPrompt(cmds)
	if err != nil {
		util.PrintErrMsg("failed to confirm commands should be executed", err)
		return
	}
	if !shouldProceed {
		fmt.Println("Exited without executing any commands")
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		util.PrintErrMsg("failed to get current directory", err)
	}

	for idx, cmd := range cmds {
		fmt.Printf(util.BlueFont("\nStep %d: %s\n"), idx+1, cmd)

		// replace argument
		if strings.Contains(cmd, "__arg__") {
			arg, err := util.InputPrompt("Enter a value to replace __arg__")
			if err != nil {
				util.PrintErrMsg("failed to enter arg", err)
				return
			}
			cmd = strings.Replace(cmd, "__arg__", arg, -1)
		}

		// If using a cd, change the current directory
		if strings.HasPrefix(cmd, "cd ") {
			targetDir := strings.TrimSpace(cmd[3:])
			expandDir, err := expandHomeDir(targetDir)
			if err != nil {
				util.PrintErrMsg("failed to expand home directory", err)
				return
			}

			if !filepath.IsAbs(expandDir) {
				expandDir = filepath.Join(currentDir, expandDir)
			}
			currentDir = expandDir
		} else {
			var execCmd *exec.Cmd

			if runtime.GOOS == "windows" {
				execCmd = exec.Command("powershell.exe", "-Command", cmd)
			} else {
				execCmd = exec.Command("bash", "-c", cmd)
			}

			// Fix execution directory
			execCmd.Dir = currentDir

			// Configure standard output and standard error pipes for real-time output
			stdoutPipe, _ := execCmd.StdoutPipe()
			stderrPipe, _ := execCmd.StderrPipe()

			// Start command
			if err := execCmd.Start(); err != nil {
				util.PrintErrMsg(fmt.Sprintf("failed to start command \"%s\"", cmd), err)
				break
			}

			// Reading standard output and standard error output
			go printOutput(stdoutPipe)
			go printOutput(stderrPipe)

			// Wait for command to finish
			err := execCmd.Wait()
			if err != nil {
				util.PrintErrMsg(fmt.Sprintf("failed to command \"%s\"", cmd), err)
				break
			}
		}

		fmt.Printf(util.GreenFont("Completed step %d command: %s\n"), idx+1, cmd)
	}
}

func printOutput(pipe io.ReadCloser) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func expandHomeDir(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return strings.Replace(path, "~", homeDir, 1), nil
	}
	return path, nil
}
