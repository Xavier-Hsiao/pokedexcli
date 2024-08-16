package main

import (
	"errors"
	"fmt"
)

func callMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page already")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return errors.New("failed to fetch data")
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	// Update next page url for evey map command call
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous
	return nil
}
