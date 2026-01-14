package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type commandFunc func(args []string)

var builtins = []string{
	"echo",
	"type",
	"exit",
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	commands := map[string]commandFunc{
		"exit": func(args []string) {
			os.Exit(0)
		},
		"echo": func(args []string) {
			fmt.Println(strings.Join(args, " "))
		},
		"type": func(args []string) {
			if len(args) == 1 {
				if slices.Contains(builtins, args[0]) {
					fmt.Printf("%s is a shell builtin\n", args[0])
					return
				}

				fmt.Printf("%s: not found\n", args[0])
				return
			}
		},
	}

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		parts := strings.Fields(command)
		instruction := parts[0]

		handler, ok := commands[instruction]

		if !ok {
			fmt.Printf("%s: command not found\n", instruction)
			continue
		}

		handler(parts[1:])

	}
}
