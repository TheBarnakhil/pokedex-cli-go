package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for k := range cfg.caughtPokemon {
		fmt.Println("  -", k)
	}
	return nil
}
