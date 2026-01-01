package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	location := ""
	if len(args) > 0 {
		location = args[0]
	}

	if location == "" {
		return fmt.Errorf("No location given")
	}

	loc, err := cfg.client.GetLocation(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", loc.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range(loc.PokemonEncounters) {
		fmt.Printf("%s\n", enc.Pokemon.Name)
	}

	return nil
}