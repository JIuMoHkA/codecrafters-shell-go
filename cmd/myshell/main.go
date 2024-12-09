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
	// Uncomment this block to pass the first stage
	
	// Wait for user input
	for {
		// Print '$' to emulate shell prompt
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		// Remove trailing newline
		input = strings.TrimSuffix(input, "\n")

		// Print the input back to the user
		fmt.Printf("%v: not found\n", input)
	}

}
