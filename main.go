package main

import (
	"github.com/ivanorribo/PokeAPI_bootdev/internal/pokeapi"
)

func main() {
	cfg := &config{
		client:        pokeapi.NewClient(),
		caughtPokemon: make(map[string]*pokeapi.Pokemon),
	}
	startRepl(cfg)
}
