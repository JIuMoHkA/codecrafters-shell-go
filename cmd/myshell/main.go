package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	// Wrap the input reader in a bufio.Reader
	reader := bufio.NewReader(os.Stdin)
	
	// REPL loop
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := reader.ReadString('\n')

		// Remove trailing newline
		input = strings.TrimSuffix(input, "\n")

		// Handle exit command
		switch input {
			case "exit 0":
				os.Exit(0)
			default:
				fmt.Printf("%v: not found\n", input)
		}
	}

}
