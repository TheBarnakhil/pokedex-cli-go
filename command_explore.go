package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("error: you must provide a location name")
	}
	data, err := cfg.pokeapiClient.GetPokemonInArea(args[0])
	if err != nil {
		return fmt.Errorf("error in pokeApi package: %w", err)
	}

	results := data.PokemonEncounters
	for _, result := range results {
		fmt.Println(result.Pokemon.Name)
	}

	return nil
}
