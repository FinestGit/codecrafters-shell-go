package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func handleCommand(command string) {
	switch command {
	case "exit":
		os.Exit(0)
	default:
		fmt.Println(command + ": command not found")
	}
}
