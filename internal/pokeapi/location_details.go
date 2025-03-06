package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"errors"
	"fmt"
)


type LocationDetails struct {
	ID       int `json:"id"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationDetails(url string) (LocationDetails, error) {
	locations := LocationDetails{}
	dat, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(dat, locations)
		if err != nil {
			return locations, err
		}
	}
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locations, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return locations,  errors.New(res.Status)
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return locations, errors.New("Error unmarshaling json")
	}

	return locations, nil
}
