package commands

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var builtins = []string{
	"echo",
	"type",
	"exit",
	"pwd",
	"cd",
}

func CommandExit(args []string) {
	os.Exit(0)
}

func CommandEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func CommandPwd(args []string) {
	if currentDirectory, _ := os.Getwd(); currentDirectory != "" {
		fmt.Printf("%s\n", currentDirectory)
	}
}

func CommandCd(args []string) {
	if args[0] == "" {
		fmt.Println("Please provide a valid path")
		return
	}

	path := args[0]

	if path == "~" {
		path = os.Getenv("HOME")
	}

	if err := os.Chdir(path); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
	}
}

func CommandType(args []string) {
	if len(args) != 1 {
		fmt.Println("Provide only one argument")
		return
	}

	if slices.Contains(builtins, args[0]) {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return
	}

	if execPath, err := exec.LookPath(args[0]); err == nil {
		fmt.Printf("%s is %s\n", args[0], execPath)
		return
	}

	fmt.Printf("%s: not found\n", args[0])
}
