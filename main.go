package main

import (
	"time"

	pokeApi "github.com/TheBarnakhil/pokedex-cli-go/internal/pokeApi"
)

func main() {
	pokeClient := pokeApi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
