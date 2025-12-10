package main

import (
	"fmt"
	"github.com/richardw55555/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != "" {
		url = cfg.Next
	}

	resp, err := pokeapi.GetLocationAreas(url)
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

func commandMapB(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Previous != "" {
		url = cfg.Previous
	} else {
		fmt.Println("You're on the first page")
		return nil
	}

	resp, err := pokeapi.GetLocationAreas(url)
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
	Next     string
	Previous string
}