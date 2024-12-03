package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		// Display the prompt
		fmt.Print("ccsh> ")

		// Read a line of input from the user
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// If the user didn't type anything, loop again
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		var command = parts[0]
		var args = parts[1:]

		switch command {
		case "exit":
			os.Exit(0)

		case "cd":
			var path string
			var err error

			if len(args) > 0 {
				path = args[0]
			} else {
				path, err = os.UserHomeDir()
			}
			err = os.Chdir(path)
			if err != nil {
				fmt.Printf("%v\n", err)
			}

		case "pwd":
			dir, _ := os.Getwd()
			fmt.Println(dir)

		default:
			cmd := exec.Command(command, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
