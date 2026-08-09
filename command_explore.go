package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	locationAreaDetail, err := cfg.pokeapiClient.GetLocationAreaByName(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")

	pokemonEncounters := locationAreaDetail.PokemonEncounters
	for _, pokemonEncounter := range pokemonEncounters {
		name := pokemonEncounter.Pokemon.Name
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
