package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client Client) GetLocationAreaByName(name string) (LocationAreaDetail, error) {
	targetURL := baseURL + "/" + name

	// Look up cache
	if body, ok := client.cache.Get(targetURL); ok {
		var locationAreaDetail LocationAreaDetail
		if err := json.Unmarshal(body, &locationAreaDetail); err != nil {
			return LocationAreaDetail{}, err
		}
		// cache hit
		fmt.Println("🚀 cache hit")
		return locationAreaDetail, nil
	}

	// cache miss -> send http request
	fmt.Println("📩 cache miss -> send http request")
	res, err := http.Get(targetURL)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaDetail{}, fmt.Errorf("pokedex api request failed with status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	client.cache.Add(targetURL, body)

	var locationAreaDetail LocationAreaDetail

	if err := json.Unmarshal(body, &locationAreaDetail); err != nil {
		return LocationAreaDetail{}, err
	}

	return locationAreaDetail, nil
}
