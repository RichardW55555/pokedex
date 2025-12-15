package main

import (
    "time"
    "github.com/richardw55555/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
        client: pokeapi.NewClient(
            5*time.Second,  // HTTP timeout
            5*time.Second,  // cache interval
        ),
    }

	startRepl(cfg)
}
