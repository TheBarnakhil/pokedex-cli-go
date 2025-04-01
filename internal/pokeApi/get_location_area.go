package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"

	var finalData []byte

	if pageURL != nil {
		url = *pageURL
	}

	if data, exists := c.cache.GetFromCache(url); exists {
		finalData = data
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("error creating request: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("error fetching response: %w", err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("error reading data from response: %w", err)
		}

		finalData = data

		c.cache.AddToCache(url, data)
	}

	var response LocationAreaResponse
	if err := json.Unmarshal(finalData, &response); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error marshaling data from response: %w", err)
	}

	return response, nil
}
