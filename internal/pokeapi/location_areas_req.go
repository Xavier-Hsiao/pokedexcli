package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	// Assign nextLocationsURL to `fullURL`
	if pageURL != nil {
		fullURL = *pageURL
	}

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

	// Returns the data as a slice of bytes
	// and any error encountered during the read.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// Unmarshel data into struct
	var locationResponse LocationAreasResp
	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return LocationAreasResp{}, fmt.Errorf("failed to unmarshel json: %v", err)
	}

	return locationResponse, nil

}
