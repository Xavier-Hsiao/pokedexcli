package main

import (
	"errors"
	"fmt"
)

func callInsepect(cfg *config, arg ...string) error {
	if len(arg) == 0 {
		return errors.New("pokemon name is required")
	}

	pokemonName := arg[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught this pokemon")
	}

	// Get required pokemon info
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, tpy := range pokemon.Types {
		fmt.Printf(" - %s\n", tpy.Type.Name)
	}

	return nil
}
