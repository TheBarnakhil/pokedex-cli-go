module github.com/TheBarnakhil/pokedex-cli-go

go 1.24.1

require github.com/TheBarnakhil/pokedex-cli-go/internal/pokeApi v0.0.0

replace github.com/TheBarnakhil/pokedex-cli-go/internal/pokeApi => ./internal/pokeApi

require github.com/TheBarnakhil/pokedex-cli-go/internal/pokecache v0.0.0 // indirect

replace github.com/TheBarnakhil/pokedex-cli-go/internal/pokecache => ./internal/pokecache
