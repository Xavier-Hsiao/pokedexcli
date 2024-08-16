package main

import (
	"fmt"
	"log"
)

func callMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}
