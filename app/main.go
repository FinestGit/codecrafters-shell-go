package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	echoCommand = "echo"
	exitCommand = "exit"
	typeCommand = "type"
)

var availableCommands = map[string]struct{}{exitCommand: {}, echoCommand: {}, typeCommand: {}}

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
	default:
		fmt.Println(command + ": command not found")
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
		fmt.Printf("%s: not found\n", command)
	}
}
