package pokeapi

import (
	"time"

	"github.com/xavier-hsiao/pokedexcli/internal/pokecache"
)

type Client struct {
	cache *pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{cache: pokecache.NewCache(interval)}
}
