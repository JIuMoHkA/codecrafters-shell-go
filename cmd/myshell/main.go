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

		input = strings.TrimSuffix(input, "\n")
		parts := strings.Fields(input)

		cmd := parts[0]


		

		// Handle exit command
		switch cmd{
			case "exit" :
				if len(parts) > 1 && parts[1] == "0" {
					os.Exit(0)
				} else {
					fmt.Printf("%v: not found\n", input)
				}
			case "echo":
				fmt.Println(strings.Join(parts[1:], " "))

			default:
				fmt.Printf("%v: not found\n", input)
		}
	}

}
