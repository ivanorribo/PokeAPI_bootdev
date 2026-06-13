package main

import (
	"testing"

	"github.com/ivanorribo/PokeAPI_bootdev/internal/pokeapi"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Go is Great",
			expected: []string{"go", "is", "great"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length %d, got %d", len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word %q, got %q", expectedWord, word)
			}
		}
	}
}

func TestGetCommands(t *testing.T) {
	commands := getCommands()
	expectedCommands := []string{"exit", "help", "map", "mapb"}

	for _, cmd := range expectedCommands {
		if _, exists := commands[cmd]; !exists {
			t.Errorf("Expected command %q to exist, but it does not.", cmd)
		}
	}
}

func TestCommandHelp(t *testing.T) {
	cfg := &config{}
	err := commandHelp(cfg)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestCommandMap(t *testing.T) {
	cfg := &config{
		client: pokeapi.NewClient(),
	}
	err := commandMap(cfg)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
