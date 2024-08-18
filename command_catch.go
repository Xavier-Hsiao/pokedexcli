package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callCatch(cfg *config, arg ...string) error {
	if len(arg) == 0 {
		return errors.New("pokemon name is required")
	}

	pokemonName := arg[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return errors.New("failed to fetch data")
	}

	// Use random number to determine the chance of cathing the pokemon
	// based upon the pokemon's `base experience`
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	fmt.Printf("You've catched %s", pokemonName)

	return nil
}
