package main

import (
	"bufio"
	"fmt"
	"io"
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

		commands := strings.Split(input, "|")
		var cmds []*exec.Cmd
		var output io.ReadCloser

		for _, command := range commands {
			command = strings.TrimSpace(command)
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
					path, _ = os.UserHomeDir()
				}
				err = os.Chdir(path)
				if err != nil {
					fmt.Printf("%v\n", err)
				}

			case "pwd":
				pr, pw := io.Pipe()
				go pwd(pw)
				if len(commands) == 1 {
					io.Copy(os.Stdout, pr)
				} else {
					output = pr
				}

			default:
				cmd := exec.Command(command, args...)
				cmd.Stderr = os.Stderr

				cmds = append(cmds, cmd)

				if output != nil {
					cmd.Stdin = output
				}
				output, _ = cmd.StdoutPipe()

			}

			if len(cmds) > 0 {
				cmds[len(cmds)-1].Stdout = os.Stdout
			}

			for _, cmd := range cmds {
				cmd.Start()
			}

			for _, cmd := range cmds {
				err := cmd.Wait()
				if err != nil {
					if cmd.ProcessState.ExitCode() == -1 {
						fmt.Printf("Command not found: %s\n", cmd.Path)
					}
				}
			}
		}
	}
}
