package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	data, err := cfg.pokeapiClient.GetLocationArea(cfg.nextLocationsURL)
	if err != nil {
		return fmt.Errorf("error in pokeApi package: %w", err)
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	results := data.Results
	for _, result := range results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	data, err := cfg.pokeapiClient.GetLocationArea(cfg.prevLocationsURL)
	if err != nil {
		return fmt.Errorf("error in pokeApi package: %w", err)
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	results := data.Results
	for _, result := range results {
		fmt.Println(result.Name)
	}

	return nil
}
