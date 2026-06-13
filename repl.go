package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ivanorribo/PokeAPI_bootdev/internal/pokeapi"
)

type config struct {
	Nexturl       *string
	Previousurl   *string
	client        *pokeapi.Client
	caughtPokemon map[string]*pokeapi.Pokemon
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() // this stops and waits for the user to enter something
		input := scanner.Text()
		inputwords := cleanInput(input)
		// Check if the user entered a command. if not, ask them to enter one and continue the loop
		if len(inputwords) == 0 {
			fmt.Printf("Please enter a command.\n")
			continue
		}
		if cmd, ok := getCommands()[inputwords[0]]; ok {
			err := cmd.callback(cfg, inputwords[1:]...)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", inputwords[0])
		}
	}
}
