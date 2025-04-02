package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("error: you must provide a pokemon name")
	}
	pokemonName := args[0]
	pokemon, exists := cfg.caughtPokemon[pokemonName]
	if !exists {
		return errors.New("you can only inspect a pokemon you have already caught")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Println("  -", pokeType.Type.Name)
	}
	return nil
}
