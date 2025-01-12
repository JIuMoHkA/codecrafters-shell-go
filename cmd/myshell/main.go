package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"
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

// Function to expand environment variables
func expandVariables(input string) string {
	var result strings.Builder
	inVar := false
	varName := strings.Builder{}

	for _, r := range input {
		if inVar {
			// If inVar is true, we're inside a variable name
			if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
				varName.WriteRune(r)
			} else {
				// We've reached the end of the variable name, expand it
				envVar := os.Getenv(varName.String())
				result.WriteString(envVar)
				varName.Reset()
				inVar = false
				result.WriteRune(r)
			}
		} else if r == '$' {
			// Start capturing variable name
			inVar = true
		} else {
			result.WriteRune(r)
		}
	}

	// Handle any remaining variable
	if varName.Len() > 0 {
		envVar := os.Getenv(varName.String())
		result.WriteString(envVar)
	}

	return result.String()
}

func parseInput(input string) []string {
	var result []string
	var curResult strings.Builder
	var isSingleQuoted, isDoubleQuoted = false, false
	escapeFlg := false

	for _, r := range input {
		switch {
		case escapeFlg:
			if isDoubleQuoted && (r != '$' && r != '"' && r != '\\') {
				curResult.WriteRune('\\')
			}
			curResult.WriteRune(r)
			escapeFlg = false
		case r == '\'':
			if !isDoubleQuoted {
				isSingleQuoted = !isSingleQuoted
			} else {
				curResult.WriteRune(r)
			}
		case r == '"':
			if !isSingleQuoted {
				isDoubleQuoted = !isDoubleQuoted
			} else {
				curResult.WriteRune(r)
			}
		case r == '\\':
			if isSingleQuoted {
				curResult.WriteRune(r)
			} else {
				escapeFlg = true
			}
		case unicode.IsSpace(r):
			if !isSingleQuoted && !isDoubleQuoted {
				if curResult.Len() > 0 {
					part := curResult.String()
					if !isSingleQuoted {
						part = expandVariables(part)
					}
					result = append(result, part)
					curResult.Reset()
				}
			} else {
				curResult.WriteRune(r)
			}
		default:
			curResult.WriteRune(r)
		}
	}

	if curResult.Len() > 0 {
		part := curResult.String()
		if !isSingleQuoted {
			part = expandVariables(part)
		}
		result = append(result, part)
	}

	return result
}

func main() {

	// Wrap the input reader in a bufio.Reader
	reader := bufio.NewReader(os.Stdin)
	builtinCommands := []string{"echo", "type", "exit", "pwd", "cd"}

	// REPL loop
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := reader.ReadString('\n')
		if len(input) == 1 {
			continue
		}

		input = strings.TrimSuffix(input, "\n")
		parts := parseInput(input)

		cmd := parts[0]

		// Handle builtin commands
		switch cmd {
		case "exit":
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
		case "cd":
			path := parts[1]
			if path == "~" {
				path = os.Getenv("HOME")
			}

			err := os.Chdir(path)

			if err != nil {
				fmt.Printf("cd: %v: No such file or directory\n", parts[1])
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
