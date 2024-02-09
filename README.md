# mycmd
mycmd is a command-line tool that allows you to manage frequently used commands. It's sudden, but you love terminal software, don't you? You think you can't be apart from it for a moment, right? I understand that feeling very well. However, in reality, that is very difficult. This is because it is very hard to remember all the commands. It is very tedious to copy and execute many commands from Notepad. But rest assured. I developed this tool for people like you. Surely, your problem will be solved. Welcome to the world of dreams.

## Preparation

### Linux or Mac:

1. Download
    ```sh
    # Linux
    curl -L -o ~/mycmd "https://github.com/ucho456job/mycmd/releases/download/v1.0.0/mycmd_linux"

    # Mac
    curl -L -o ~/mycmd "https://github.com/ucho456job/mycmd/releases/download/v1.0.0/mycmd_mac"
    ```

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
    Invoke-WebRequest -Uri "https://github.com/ucho456job/mycmd/releases/download/v1.0.0/mycmd-windows_amd64.exe" -OutFile "C:\Program Files\mycmd\mycmd.exe"
    ```

3. Add to PATH environment variable
    ```sh
    $env:Path += ";C:\Program Files\mycmd"
    ```

4. Make changes to system environment variables permanent
    ```sh
    [Environment]::SetEnvironmentVariable("Path", $env:Path, [EnvironmentVariableTarget]::Machine)
    ```

## Commands

- `clip`: Copy commands to the clipboard. '-g' and '-t' options can be used.
- `edit`: Edit cmd.json. '-e' option can be used.
- `exec`: Execute commands. '-g' and '-t' options can be used.
- `help`: Show help.
- `read`: Show commands. '-g' and '-t' options can be used.
- `view`: Show cmd.json.

## Options

- `-g`: Specify group.
- `-t`: Specify task.
- `-e`: Specify editor. path must exist.

## How to use

1. Run `mycmd edit` to open $HOME/.mycmd/cmd.json
2. Please update cmd.json referring to ["About cmd.json configuration structure"](#about-cmdjson-configuration-structure)
3. You can run registered commands by running `mycmd exec` and following the prompts.

## About cmd.json configuration structure

To customize the execution of tasks within our tool, you need to modify the cmd.json configuration file according to the following structure.

### Sample cmd.json:

```json
{
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

## Contribution

If you encounter any bugs or have feature requests, please create an issue on the [issue tracker](https://github.com/ucho456job/mycmd/issues).

## License

See the [LICENSE](LICENSE) file for details.

## Disclaimer

The author is not responsible for any loss or damage caused in connection with the use of this project.
