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

		// Exit the shell if the user types "exit"
		if input == "exit" {
			os.Exit(0)
		}

		// Prepare the Command to be run, capturing the input
		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
