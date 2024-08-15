package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

// Create custom Client to set timeout
type Client struct {
	httpClient http.Client
}

func CreateNewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
