package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeApi "github.com/TheBarnakhil/pokedex-cli-go/internal/pokeApi"
)

type (
	cliCommand struct {
		name        string
		description string
		callback    func(*config, ...string) error
	}

	config struct {
		pokeapiClient    *pokeApi.Client
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
		args := []string{}
		command, ok := registry[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if len(words) >= 2 {
			args = words[1:]
		}

		if err := command.callback(cfg, args...); err != nil {
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
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Get all the pokemon available in the area provided",
			callback:    commandExplore,
		},
	}
}
