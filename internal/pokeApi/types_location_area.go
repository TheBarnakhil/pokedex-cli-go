package pokeapi

type (
	Result struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	LocationAreaResponse struct {
		Count    int      `json:"count"`
		Next     *string  `json:"next"`
		Previous *string  `json:"previous"`
		Results  []Result `json:"results"`
	}
)
