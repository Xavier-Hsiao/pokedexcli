package main

import "github.com/Xavier-Hsiao/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.CreateNewClient(),
	}
	startRpl(&cfg)
}
