package pokeapi

import (
	"net/http"
	"time"

	"github.com/Xavier-Hsiao/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// Create custom Client to set timeout
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func CreateNewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.CreateNewCache(cacheInterval),
	}
}
