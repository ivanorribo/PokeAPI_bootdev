package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 || args[0] == "" {
		return fmt.Errorf("pokemon name is required")
	}
	pokemonName := args[0]
	pokemon, found := cfg.caughtPokemon[pokemonName]
	if !found {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", s.Stat.Name, s.Base_stat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	return nil
}
