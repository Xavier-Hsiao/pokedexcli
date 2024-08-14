package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRpl() {
	// Get user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")

		scanner.Scan()
		input := scanner.Text()

		fmt.Println("You've entered:", input)
	}
}
