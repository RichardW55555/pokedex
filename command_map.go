package main

import (
	"fmt"
	"github.com/richardw55555/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != "" {
		url = cfg.Next
	}

	resp, err := cfg.client.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}

func commandMapB(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Previous != "" {
		url = cfg.Previous
	} else {
		fmt.Println("You're on the first page")
		return nil
	}

	resp, err := cfg.client.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}

type config struct {
	Next          string
	Previous      string
	client        pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
}