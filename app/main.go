package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	echoCommand = "echo"
	exitCommand = "exit"
	typeCommand = "type"
	pwdCommand  = "pwd"
	cdCommmand  = "cd"
)

var availableCommands = map[string]struct{}{exitCommand: {}, echoCommand: {}, typeCommand: {}, pwdCommand: {}, cdCommmand: {}}

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	repl()
}

func repl() {
	for {
		fmt.Print("$ ")
		var command string
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		handleCommand(command[:len(command)-1])
	}
}

func handleCommand(commandString string) {
	commandArray := strings.Split(commandString, " ")
	command := commandArray[0]
	additionalArgs := commandArray[1:]
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		handleEcho(additionalArgs)
	case "type":
		handleType(additionalArgs)
	case "pwd":
		handlePwd()
	case "cd":
		if len(additionalArgs) == 0 {
			return
		}
		handleCd(additionalArgs[0])
	default:
		attemptExecutable(command, additionalArgs)
	}
}

func isCommandBuiltin(command string) bool {
	_, exists := availableCommands[command]
	return exists
}

func handleEcho(args []string) {
	if len(args) > 0 {
		fmt.Println(strings.Join(args, " "))
	} else {
		fmt.Println()
	}
}

func handleType(args []string) {
	command := args[0]
	if isCommandBuiltin(command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else {
		found, path := searchPath(command)
		if !found {
			fmt.Printf("%s: not found\n", command)
			return
		}
		fmt.Printf("%s is %s\n", command, path)
	}
}

func searchPath(commandToFind string) (bool, string) {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	for _, dir := range paths {
		filepath := filepath.Join(dir, commandToFind)

		info, err := os.Stat(filepath)

		if err != nil {
			continue
		}

		if info.IsDir() {
			continue
		}

		if info.Mode()&0111 == 0 {
			continue
		}

		return true, filepath
	}
	return false, ""
}

func attemptExecutable(executable string, args []string) {
	found, path := searchPath(executable)
	if !found {
		fmt.Println(executable + ": command not found")
		return
	}
	cmd := exec.Command(path, args...)
	cmd.Args[0] = executable
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing command:", err)
		return
	}

	fmt.Print(string(stdout))
}

func handlePwd() {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(dir)
}

func handleCd(targetDir string) {
	if targetDir == "~" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return
		}
		targetDir = homeDir
	}
	err := os.Chdir(targetDir)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", targetDir)
	}
}
