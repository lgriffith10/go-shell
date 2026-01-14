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
	parts := strings.Split(command, " ")

	instruction := parts[0]

	switch instruction {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(strings.Join(parts[1:], " "))
	default:
		fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
	}

	main()
}
