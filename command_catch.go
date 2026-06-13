package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error { //args will be the name of the pokemon to catch
	if len(args) == 0 || args[0] == "" {
		return fmt.Errorf("pokemon name is required")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemon, err := cfg.client.GetPokemon(args[0])
	if err != nil {
		return fmt.Errorf("error fetching pokemon: %w", err)
	}
	roll := rand.Intn(pokemon.BaseExperience)
	if roll < 60 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
