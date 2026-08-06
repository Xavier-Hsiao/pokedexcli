package main

import (
	"fmt"

	"github.com/xavier-hsiao/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	areasResp, err := pokeapi.GetAreas(cfg.nextURL)
	if err != nil {
		return err
	}

	// todo: write next and prev back to the config, print areas on the console
	cfg.nextURL = areasResp.Next
	cfg.prevURL = areasResp.Previous

	for _, area := range areasResp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	// if the user is already on the first page
	if cfg.prevURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	areaResp, err := pokeapi.GetAreas(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = areaResp.Next
	cfg.prevURL = areaResp.Previous

	for _, area := range areaResp.Results {
		fmt.Println(area.Name)
	}

	return nil
}
