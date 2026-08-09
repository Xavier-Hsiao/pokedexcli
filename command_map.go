package main

import (
	"fmt"
)

func commandMap(cfg *config, args []string) error {
	areasResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = areasResp.Next
	cfg.prevURL = areasResp.Previous

	for _, area := range areasResp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *config, args []string) error {
	// if the user is already on the first page
	if cfg.prevURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	areaResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevURL)
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
