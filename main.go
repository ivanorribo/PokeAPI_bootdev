package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", inputwords[0])
		}
	}
}
