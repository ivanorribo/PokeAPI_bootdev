package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error { //command map to display the map locations and advance forward
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

func commandMapb(cfg *config, args ...string) error { //command map to display the map locations and advance backward
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
