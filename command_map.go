package main

import (
	"fmt"
	"log"

	"github.com/Xavier-Hsiao/pokedexcli/internal/pokeapi"
)

func callMap() error {
	pokeapiClient := pokeapi.CreateNewClient()
	resp, err := pokeapiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}
