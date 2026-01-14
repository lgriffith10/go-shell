package main

import (
	"fmt"
)

func main() {
	var command string

	fmt.Print("$ ")
	fmt.Scanln(&command)

	fmt.Printf("%s: command not found\n", command)
	fmt.Print("$ ")
}
