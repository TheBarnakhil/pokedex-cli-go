package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInArea(areaName string) (PokemonInArea, error) {
	url := baseURL + locationEndpoint + areaName

	var finalData []byte

	if data, exists := c.cache.GetFromCache(url); exists {
		finalData = data
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonInArea{}, fmt.Errorf("error while creating request: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonInArea{}, fmt.Errorf("error while fetching response: %w", err)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return PokemonInArea{}, fmt.Errorf("error while reading response body: %w", err)
		}

		c.cache.AddToCache(url, data)
		finalData = data
	}

	var pokemon PokemonInArea
	if err := json.Unmarshal(finalData, &pokemon); err != nil {
		return PokemonInArea{}, fmt.Errorf("error while unmarshaling data: %w", err)
	}

	return pokemon, nil
}
