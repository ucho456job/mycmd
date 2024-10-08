# mycmd
mycmd is a tool for streamlining command line operations. This tool allows you to easily manage frequently used commands and quickly perform the tasks you need. By registering complex and long commands, you can automatically execute them simply by selecting them, or copy them to the clipboard for reuse. This makes tedious command line tasks easier and faster.

## Preparation

### Linux or Mac:

1. Download
    ```sh
    # Linux
    curl -L -o ~/mycmd "https://github.com/ucho456job/mycmd/releases/download/v1.0.1/mycmd_linux"

    # Mac
    curl -L -o ~/mycmd "https://github.com/ucho456job/mycmd/releases/download/v1.0.1/mycmd_mac"
    ```
    *If you are unable to download using commands due to security reasons, please download directly from [here](https://github.com/ucho456job/mycmd/releases/tag/v1.0.0). Next, rename the downloaded binary her file to `mycmd`.

2. Granting execution privileges
    ```sh
    chmod +x ~/mycmd
    ```

3. Move binary files to /usr/local/bin
    ```sh
    sudo mv ~/mycmd /usr/local/bin/
    ```

### Windows:

Before running the steps below, please open PowerShell with administrator privileges.

1. Create a directory to download the binary
    ```powershell
    New-Item -Path "C:\Program Files\mycmd" -ItemType Directory
    ```

2. Download
    ```powershell
    Invoke-WebRequest -Uri "https://github.com/ucho456job/mycmd/releases/download/v1.0.1/mycmd-windows_amd64.exe" -OutFile "C:\Program Files\mycmd\mycmd.exe"
    ```
    *If you are unable to download using commands due to security reasons, please download directly from [here](https://github.com/ucho456job/mycmd/releases/tag/v1.0.0). Next, rename the downloaded exe file to `mycmd` and place it at `C:\Program Files\mycmd\mycmd.exe`.

3. Add to PATH environment variable
    ```sh
    $env:Path += ";C:\Program Files\mycmd"
    ```

4. Make changes to system environment variables permanent
    ```sh
    [Environment]::SetEnvironmentVariable("Path", $env:Path, [EnvironmentVariableTarget]::Machine)
    ```

## How to use

1. Run `mycmd edit` to open $HOME/.mycmd/cmd.json
2. Please update cmd.json referring to ["About cmd.json configuration structure"](#about-cmdjson-configuration-structure)
3. You can run registered commands by running `mycmd exec` and following the prompts.

## Commands

- `clip`: Copy commands to the clipboard. '-g' and '-t' options can be used.
- `edit`: Edit cmd.json. '-e' option can be used.
- `exec`: Execute commands. '-g', '-t', '-f' and '-a' options can be used.
- `help`: Show help.
- `read`: Show commands. '-g' and '-t' options can be used.
- `view`: Show cmd.json.

## Options

- `-g`: Specify group.
- `-t`: Specify task.
- `-f`: Run without confirmation.
- `-a`: Specifies the argument to replace with `__arg__`.
- `-e`: Specify editor. path must exist.

## About cmd.json configuration structure

To customize the execution of tasks within our tool, you need to modify the cmd.json configuration file according to the following structure.

### Sample cmd.json:

```json
{
  "001_group1": {
    "001_task1": [
      "command1",
      "command2"
    ],
    "002_task2": [
      "command1"
    ]
  },
  "002_group2": {
    "001_task1": [
      "command1",
      "command2"
    ]
  }
}
```

### Structure Explanation:

- **group**: Represents a collection of tasks. Each group should have a unique name and contain one or more tasks.
- **task**: Identifies an individual task within a group. Each task is represented by a key and is associated with an array of commands to be executed.
- **commands**: A list of string values, each representing a command to be executed as part of the task. Commands within the task are executed in the order they are listed.

### Editing Guidelines:

- **Defining a New Group**: To define a new group of tasks, add a new key at the root level of cmd.json. Use a descriptive name for your group to easily identify the set of tasks it contains.
- **Adding a New Task**: Within a group, add a new key to represent the task. The key should be descriptive of the task's purpose. Associate this key with an array of commands that defines the task's execution logic.
- **Specifying Commands**: For each task, specify the commands to be executed in an array format. Ensure each command is a string and listed in the order you wish them to be executed.
- **Removing Tasks or Groups**: To remove a task, delete its key from the group. If a group no longer contains any tasks, consider removing the group as well to keep the configuration clean and concise.

## About dynamic argument

The tool supports dynamic argument input for commands specified in cmd.json. This feature allows for interactive execution of commands by enabling the user to input arguments at runtime. However, only one argument can be dynamically changed.

### Sample cmd.json:
```json
{
  "Linux": {
    "list": [
      "ls __arg__"
    ]
  }
}
```

### How It Works:

You can specify a placeholder for runtime arguments using `__arg__`. This placeholder will be replaced with the user's input right before the command is executed.

### Example:

```sh
$ mycmd exec -g Linux -t list
Selected commands are below:
Step 1: ls __arg__
Execute these commands? [Y/n]: Y

Step 1: ls __arg__
Enter a value to replace __arg__: -a
.
..
.git
.vscode
Makefile
README.md
go.mod
go.sum
main.go
pkg

Completed step 1 command: ls -a
```

## Important Note on Command Execution
Note that this tool uses bash to run commands in non-Windows environments. This choice was made due to bash's wide adoption and compatibility with a variety of Linux distributions, including popular choices such as Ubuntu and Mac in WSL2 environments.

If you are planning to run commands that are specific to bash or utilize features unique to bash, ensure that bash is available on your system. This is especially relevant for WSL2 users or those on Unix-like systems where bash is not the default shell.

For Windows users, commands are executed using PowerShell. Ensure that your commands are compatible with the environment you're working in.

## Contribution

If you encounter any bugs or have feature requests, please create an issue on the [issue tracker](https://github.com/ucho456job/mycmd/issues).

## License

See the [LICENSE](LICENSE) file for details.

## Disclaimer

The author is not responsible for any loss or damage caused in connection with the use of this project.
