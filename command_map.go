package main

import (
	"fmt"
	"log"
)

func callMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationsURL)

	if err != nil {
		log.Fatal(err)
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
