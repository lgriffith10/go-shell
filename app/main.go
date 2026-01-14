package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")

	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	command = strings.TrimSpace(command)

	fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
}
