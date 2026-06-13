package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 || args[0] == "" {
		return fmt.Errorf("location name is required")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", args[0])
	pokemonEncounters, err := cfg.client.PokemonEncounters(url)
	if err != nil {
		return err
	}
	if len(pokemonEncounters) == 0 {
		fmt.Printf("No Pokemon encounters found for location: %s\n", args[0])
		return nil
	}
	fmt.Printf("Exploring %s...\nFound Pokemon: \n", args[0])
	for _, encounter := range pokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
