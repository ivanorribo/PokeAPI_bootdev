package main

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		println("You have not caught any Pokemon yet.")
		return nil
	}
	println("Caught Pokemon:")
	for name := range cfg.caughtPokemon {
		println(" - " + name)
	}
	return nil
}
