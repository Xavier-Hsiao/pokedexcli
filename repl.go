package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := cleanInput(scanner.Text())
		if len(input) < 1 {
			continue
		}
		fmt.Printf("Your command was: %s\n", input[0])
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func cleanInput(text string) []string {
	// split the user's input into words based on whitespace
	// lowercase the input and trim any leading or trailing whitespaces
	words := strings.Fields(strings.ToLower(text))
	return words
}
