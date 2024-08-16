package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area?offset=0&limit=20"
	fullURL := baseURL + endpoint

	// Assign nextLocationsURL to `fullURL`
	if pageURL != nil {
		fullURL = *pageURL
	}

	// Check cache hit
	data, ok := c.cache.Get(fullURL)
	if ok {
		// Cache hit
		// Unmarshel data into struct
		fmt.Println("Cache hit!")
		var locationResponse LocationAreasResp
		if err := json.Unmarshal(data, &locationResponse); err != nil {
			return LocationAreasResp{}, fmt.Errorf("failed to unmarshel json: %v", err)
		}
		return locationResponse, nil
	}
	fmt.Println("Cache miss!")

	// Cache miss
	// Make the http request
	resp, err := http.Get(fullURL)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return LocationAreasResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	var locationResponse LocationAreasResp
	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return LocationAreasResp{}, fmt.Errorf("failed to unmarshel json: %v", err)
	}

	c.cache.Add(fullURL, data)

	return locationResponse, nil

}
