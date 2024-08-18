package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// Check cache hit
	data, ok := c.cache.Get(fullURL)
	if ok {
		// Cache hit
		// Unmarshel data into struct
		fmt.Println("Cache hit!")
		var pokemon Pokemon
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, fmt.Errorf("failed to unmarshel json: %v", err)
		}
		return pokemon, nil
	}
	fmt.Println("Cache miss!")

	// Cache miss
	// Make the http request
	resp, err := http.Get(fullURL)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("failed to unmarshel json: %v", err)
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
