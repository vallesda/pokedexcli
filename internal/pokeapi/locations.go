package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"errors"
)

type LocationsResp struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(url string) (LocationsResp, error) {
	locations := LocationsResp{}

	dat, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(dat, locations)
		if err != nil {
			return locations, err
		}
	}

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