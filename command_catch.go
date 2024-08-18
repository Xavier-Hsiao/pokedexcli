package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

	fmt.Printf("Throwing a Pokeball to %s...\n", pokemonName)
	time.Sleep(1000 * time.Microsecond)

	// Use random number to determine the chance of cathing the pokemon
	// based upon the pokemon's `base experience`
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	cfg.caughtPokemon[pokemonName] = pokemon

	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
