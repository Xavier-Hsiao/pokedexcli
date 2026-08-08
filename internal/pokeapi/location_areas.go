package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AreasResp struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

const baseURL = "https://pokeapi.co/api/v2/location-area"

func (client Client) GetAreas(pageURL *string) (AreasResp, error) {
	targetURL := baseURL
	if pageURL != nil {
		targetURL = *pageURL
	}

	// Look up cache
	if body, ok := client.cache.Get(targetURL); ok {
		var areaResp AreasResp
		if err := json.Unmarshal(body, &areaResp); err != nil {
			return AreasResp{}, err
		}
		// cache hit
		fmt.Println("🚀 cache hit")
		return areaResp, nil
	}

	// cache miss -> send http request
	fmt.Println("📩 cache miss -> send http request")
	res, err := http.Get(targetURL)
	if err != nil {
		return AreasResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return AreasResp{}, fmt.Errorf("pokedex api request failed with status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AreasResp{}, err
	}

	client.cache.Add(targetURL, body)

	var areaResp AreasResp

	if err := json.Unmarshal(body, &areaResp); err != nil {
		return AreasResp{}, err
	}

	return areaResp, nil
}
