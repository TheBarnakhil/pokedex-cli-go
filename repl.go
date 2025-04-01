package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeApi "github.com/TheBarnakhil/pokedex-cli-go/internal/pokeApi"
	"github.com/TheBarnakhil/pokedex-cli-go/internal/pokecache"
)

type (
	cliCommand struct {
		name        string
		description string
		callback    func(*config) error
	}

	config struct {
		pokeapiClient    pokeApi.Client
		cache            *pokecache.Cache
		nextLocationsURL *string
		prevLocationsURL *string
	}
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	registry := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := registry[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}

	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    getHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}
