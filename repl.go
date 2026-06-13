package main

import (
	"strings"

	"github.com/ivanorribo/PokeAPI_bootdev/internal/pokeapi"
)

type config struct {
	Nexturl     *string
	Previousurl *string
	client      *pokeapi.Client
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
