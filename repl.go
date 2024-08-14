package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRpl() {
	// Get user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")

		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}

		fmt.Println("Entered: ", cleanedInput)
	}
}

// Clean user input and return a slice of tokens
func cleanInput(str string) []string {
	cleanedStr := strings.ToLower(str)
	return strings.Fields(cleanedStr)
}
