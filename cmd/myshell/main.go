package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

// Function to check if word is in slice
func contains(slice []string, word string) bool {
	for _, a := range slice {
		if a == word {
			return true
		}
	}
	return false
}


func main() {

	// Wrap the input reader in a bufio.Reader
	reader := bufio.NewReader(os.Stdin)
	builtinCommands := []string{"echo", "type", "exit"}
	
	// REPL loop
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := reader.ReadString('\n')

		input = strings.TrimSuffix(input, "\n")
		parts := strings.Fields(input)

		cmd := parts[0]

		// Handle builtin commands
		switch cmd{
			case "exit" :
				flg, err := strconv.Atoi(parts[1])
				if err!= nil {
					flg = 1
				}
				os.Exit(flg)
			case "echo":
				fmt.Println(strings.Join(parts[1:], " "))

			case "type":
				if contains(builtinCommands, parts[1]) {
					fmt.Println(parts[1] + " is a shell builtin")
				} else {
					fmt.Printf("%v: not found\n", parts[1])
				}

			default:
				fmt.Printf("%v: not found\n", input)
		}
	}

}
