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
	callback    func() error
}

func startRepl() {
	commands := getCommands()
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
			err := command.callback()
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
	}
}
