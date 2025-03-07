package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/vallesda/pokedexcli/internal/pokeapi"
	"github.com/vallesda/pokedexcli/internal/pokedex"
)

type config struct {
	PokeClient pokeapi.Client
	Pokedex map[string]pokedex.Pokemon
	Next *string
	Prev *string
}

func startReading() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{pokeapi.NewClient(5 * time.Minute), pokedex.NewPokedex(), nil, nil}
	for {
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)
		command := words[0]
		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}

		switch command {
		case "map":
			err := commandMap(&conf, args...)
			if err != nil {
				break
			}
		case "bmap":
			err := commandBMap(&conf, args...)
			if err != nil {
				break
			}
		case "exit":
			err := commandExit(&conf, args...)
			if err != nil {
				break
			}
		case "help":
			err := commandHelp(&conf, args...)
			if err != nil {
				break
			}
		case "explore":
			err := commandExplore(&conf, args...)
			if err != nil {
				break
			}
		case "catch":
			err := commandCatch(&conf, args...)
			if err != nil {
				break
			}
		case "inspect":
			err := commandInspect(&conf, args...)
			if err != nil {
				break
			}
		case "pokedex":
			err := commandPokedex(&conf, args...)
			if err != nil {
				break
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}

func getCommandsMap() map[string]cliCommand{
	return map[string]cliCommand{
		"map": {
			name: "map",
			description: "Gets next 20 locations",
			callback: commandMap,
		},
		"bmap": {
			name: "bmap",
			description: "Gets prev 20 locations",
			callback: commandBMap,
		},
		"exit": {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"explore": {
			name: "explore",
			description: "Displays availablePokemon",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "cathes pokemon and adds it to pokedex",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "inspects pokemon in pokedex",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "displays pokemon in pokedex",
			callback: commandPokedex,
		},
	}
}