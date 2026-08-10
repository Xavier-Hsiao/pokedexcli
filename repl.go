package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/xavier-hsiao/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

// Use pointers to distinguish the API's null (no page) from an empty string.
// nil means no next page (API returns null) and pure string cannot accept nil.
type config struct {
	nextURL       *string
	prevURL       *string
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
}

func startRepl() {
	commands := getCommands()
	cfg := config{
		nextURL:       nil,
		prevURL:       nil,
		pokeapiClient: pokeapi.NewClient(5 * time.Minute),
		pokedex:       map[string]pokeapi.Pokemon{},
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
			err := command.callback(&cfg, input[1:])
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
		"explore": {
			name:        "explore",
			description: "See a list of all Pokémon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokémon and add it to the Pokédex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display the caught Pokémon in the Pokédex by name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List out all caught pokemon in the Pokedex",
			callback:    commandPokedex,
		},
	}
}
