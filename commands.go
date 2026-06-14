package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore",
			description: "Allows to explore a specific location and see the Pokemon encounters",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Allows to catch a specific Pokemon by its name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Allows to inspect a specific Pokemon by its name if you have caught it",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the list of caught Pokemon",
			callback:    commandPokedex,
		},
	}
}
