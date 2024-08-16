package main

import (
	"errors"
	"fmt"
)

func callMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationsURL)
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
