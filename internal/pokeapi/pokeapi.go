package pokeapi

import (
	"net/http"
	"time"
	"strings"

	"github.com/vallesda/pokedexcli/internal/pokecache"
)

const endpoint = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{httpClient: http.Client{Timeout: time.Minute}, cache: pokecache.NewCache(cacheInterval)}
}

func (c *Client) BuildUrl(paths ...string) string {
	path := strings.Join(paths, "/")
	return endpoint + path
}