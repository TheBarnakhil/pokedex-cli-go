package main

import "fmt"

func getHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	registry := getCommands()
	for _, command := range registry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
