package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/parser"
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

		parts := parser.ParseCommand(command)

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
