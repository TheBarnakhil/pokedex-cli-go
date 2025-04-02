package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("error: you must provide a pokemon name")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokeInfo, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	chance := rand.Intn(100) + 30
	if chance > pokeInfo.BaseExperience {
		fmt.Printf("%s was caught!", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokeInfo
	} else {
		fmt.Printf("%s escaped!", pokemonName)
	}

	return nil
}
