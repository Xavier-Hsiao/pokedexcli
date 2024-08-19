package main

import "fmt"

func callPokedex(cfg *config, arg ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("pokedex empty: your have not caught any pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
