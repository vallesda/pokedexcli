package pokeapi

import (
	"net/http"
	"time"
	"strings"
)

const endpoint = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{httpClient: http.Client{Timeout: time.Minute}}
}

func (c *Client) BuildUrl(paths ...string) string {
	path := strings.Join(paths, "/")
	return endpoint + path
}