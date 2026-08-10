package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client Client) GetPokemon(name string) (Pokemon, error) {
	targetURL := pokemonBaseURL + "/" + name

	if body, ok := client.cache.Get(targetURL); ok {
		var pokemonResp Pokemon
		if err := json.Unmarshal(body, &pokemonResp); err != nil {
			return Pokemon{}, err
		}

		fmt.Println("🚀 cache hit")
		return pokemonResp, nil
	}

	fmt.Println("📩 cache miss -> send http request")

	res, err := http.Get(targetURL)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("pokemon api request failed with status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	client.cache.Add(targetURL, body)

	var pokemonResp Pokemon

	if err := json.Unmarshal(body, &pokemonResp); err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}
