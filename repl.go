package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		
		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

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
			description: "Prints a list of commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Prints the next 20 locations in pokemon",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints the previous 20 locations in pokemon",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Prints a list of pokemon for a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catches a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "prints all the pokemon caught",
			callback:    commandPokedex,
		},
	}
}