package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// Use pointers to distinguish the API's null (no page) from an empty string.
// nil means no next page (API returns null) and pure string cannot accept nil.
type config struct {
	nextURL *string
	prevURL *string
}

func startRepl() {
	commands := getCommands()
	cfg := config{
		nextURL: nil,
		prevURL: nil,
	}
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

		if command, ok := commands[input[0]]; ok {
			err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "List out the available commands and their usage",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the name of 20 areas in Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the name of areas on the previous page",
			callback:    commandMapb,
		},
	}
}
