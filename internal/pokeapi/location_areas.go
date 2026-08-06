package pokeapi

import (
	"encoding/json"
	"fmt"
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

func GetAreas(pageURL *string) (AreasResp, error) {
	targetURL := baseURL
	if pageURL != nil {
		targetURL = *pageURL
	}

	res, err := http.Get(targetURL)
	if err != nil {
		return AreasResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return AreasResp{}, fmt.Errorf("pokedex api request failed with status: %d", res.StatusCode)
	}

	var areaResp AreasResp

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&areaResp); err != nil {
		return AreasResp{}, err
	}

	return areaResp, nil
}
