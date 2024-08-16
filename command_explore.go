package main

import (
	"errors"
	"fmt"
)

func callExplore(cfg *config, arg ...string) error {
	if len(arg) == 0 {
		return errors.New("area name is required")
	}

	areaName := arg[0]

	resp, err := cfg.pokeapiClient.GetArea(areaName)
	if err != nil {
		return errors.New("failed to fetch data")
	}

	fmt.Printf("Pokemon in %s area\n", resp.Name)

	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
