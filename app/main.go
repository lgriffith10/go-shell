package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("$ ")

	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error reading command:", err)
		return
	}

	fmt.Printf("%s: command not found\n", command)
	fmt.Print("$ ")
}
