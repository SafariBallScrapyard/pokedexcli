package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	if bytes, ok := c.cache.Get(url); ok {
		location := Location{}
		if err := json.Unmarshal(bytes, &location); err != nil {
			return Location{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)
	return locationResp, nil
}
