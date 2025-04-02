package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name string) (PokemonInfo, error) {
	url := baseURL + pokemonEndpoint + name

	var finalData []byte

	if data, exists := c.cache.GetFromCache(url); exists {
		finalData = data
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonInfo{}, fmt.Errorf("error when creating request: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonInfo{}, fmt.Errorf("error when fetching response: %w", err)

		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return PokemonInfo{}, fmt.Errorf("error when reading response body: %w", err)
		}

		finalData = data

		c.cache.AddToCache(url, finalData)
	}

	var pokemon PokemonInfo
	if err := json.Unmarshal(finalData, &pokemon); err != nil {
		return PokemonInfo{}, fmt.Errorf("error when unmarshaling data: %w", err)
	}

	return pokemon, nil
}
