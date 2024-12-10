package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"os/exec"
)


// Function to check if word is in slice
func contains(slice []string, word string) bool {
	for _, a := range slice {
		if a == word {
			return true
		}
	}
	return false
}


// Function to check if command is in path
func commandInPath(command string) (string, bool) {
	pathEnv := os.Getenv("PATH")
	path := strings.Split(pathEnv, ":")

	for _, dir := range path {
		fullPath := dir + "/" + command

		if fileInfo, err := os.Stat(fullPath); (err == nil) && (fileInfo.Mode().IsRegular()) {
			return fullPath, true
		}
	}
	return "", false	
}


func main() {

	// Wrap the input reader in a bufio.Reader
	reader := bufio.NewReader(os.Stdin)
	builtinCommands := []string{"echo", "type", "exit", "pwd"}
	
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
				if err != nil {
					flg = 1
				}
				os.Exit(flg)
			case "echo":
				fmt.Println(strings.Join(parts[1:], " "))

			case "type":
				if contains(builtinCommands, parts[1]) {
					fmt.Println(parts[1] + " is a shell builtin")
				} else if filePath, isExists := commandInPath(parts[1]); isExists {
					fmt.Println(parts[1] + " is " + filePath)
				} else {
					fmt.Printf("%v: not found\n", parts[1])
				}
			case "pwd":
				dir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error getting current directory:", err)
				} else {
					fmt.Println(dir)

				}
			default:
				_, isExists := commandInPath(cmd)

				if !isExists {
					fmt.Printf("%v: not found\n", input)
					continue
				}

				command := exec.Command(cmd, parts[1:]...)

				command.Stdout = os.Stdout
				command.Stderr = os.Stderr

				command.Run()
		}
	}
}
