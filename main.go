package main

import (
	"github.com/ivanorribo/PokeAPI_bootdev/internal/pokeapi"
)

func main() {
	cfg := &config{
		client: pokeapi.NewClient(),
	}
	startRepl(cfg)
}
