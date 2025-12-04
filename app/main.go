package main

import (
	"fmt"
	"log"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	fmt.Print("$ ")
	var command string
	_, err := fmt.Scanln(&command)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v: command not found\n", command)
}
