package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"errors"

	"github.com/vallesda/pokedexcli/internal/pokedex"
)

func (c *Client) GetPokemon(url string) (pokedex.Pokemon, error) {
	pokemons := pokedex.Pokemon{}

	dat, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(dat, pokemons)
		if err != nil {
			return pokemons, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemons, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokemons, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return pokemons,  errors.New(res.Status)
	}

	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return pokemons, errors.New("Error unmarshaling json")
	}
	
	return pokemons, nil
}