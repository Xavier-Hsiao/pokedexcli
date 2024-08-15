package main

import (
	"fmt"
	"log"

	"github.com/Xavier-Hsiao/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.CreateNewClient()
	resp, err := pokeapiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	// startRpl()
}
