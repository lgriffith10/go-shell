package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

type commandFunc func(args []string)

var builtinCommands = map[string]commandFunc{
	"exit": commands.CommandExit,
	"echo": commands.CommandEcho,
	"type": commands.CommandType,
	"pwd":  commands.CommandPwd,
	"cd":   commands.CommandCd,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

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

		parts := parseCommand(command)

		instruction := parts[0]
		args := parts[1:]

		runCommand(instruction, args)
	}
}

func runCommand(instruction string, args []string) {
	handler, ok := builtinCommands[instruction]

	if ok {
		handler(args)
		return
	}

	_, err := exec.LookPath(instruction)
	if err == nil {
		cmd := exec.Command(instruction, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()
		return
	}

	fmt.Printf("%s: command not found\n", instruction)
}

func parseCommand(input string) []string {
	var args []string
	var current strings.Builder

	isInQuotes := false

	for i := range len(input) {
		c := input[i]

		switch c {
		case '\'', '"':
			isInQuotes = !isInQuotes
		case ' ':
			if isInQuotes {
				current.WriteByte(c)
			} else if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteByte(c)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}
