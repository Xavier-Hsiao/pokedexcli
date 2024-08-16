package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRpl(cfg *config) {
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

		commandName := cleanedInput[0]
		availableCommands := getCommands()

		cmd, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// We have two commands:
// help: prints a help message describing how to use the REPL
// exit: exits the program
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// Get all available commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Summon pokedex help menu",
			callback:    callHelp,
		},
		"map": {
			name:        "map",
			description: "List out location areas",
			callback:    callMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back to the previous page of location areas",
			callback:    callMapb,
		},
		"exit": {
			name:        "exit",
			description: "Turn off pokedex",
			callback:    callExit,
		},
	}
}

// Clean user input and return a slice of tokens
func cleanInput(str string) []string {
	cleanedStr := strings.ToLower(str)
	return strings.Fields(cleanedStr)
}
