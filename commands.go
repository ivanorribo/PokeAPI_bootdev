package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Allows to move forward in the map of the Pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Allows to move backward in the map of the Pokemon locations",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error { //command map to display the map locations and advance forward
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Nexturl != nil {
		url = *cfg.Nexturl
	}
	resp, err := cfg.client.GetLocation(url)
	if err != nil {
		return err
	}
	for _, location := range resp.Results {
		fmt.Printf("%s\n", location.Name)
	}
	cfg.Nexturl = resp.Next
	cfg.Previousurl = resp.Previous
	return nil
}

func commandMapb(cfg *config) error { //command map to display the map locations and advance backward
	if cfg.Previousurl == nil {
		fmt.Println("You are on the first page")
		return nil
	}
	url := *cfg.Previousurl
	resp, err := cfg.client.GetLocation(url)
	if err != nil {
		return err
	}
	for _, location := range resp.Results {
		fmt.Printf("%s\n", location.Name)
	}
	cfg.Nexturl = resp.Next
	cfg.Previousurl = resp.Previous
	return nil
}
