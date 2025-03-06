package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/vallesda/pokedexcli/internal/pokeapi"
)

type config struct {
	PokeClient pokeapi.Client
	Next *string
	Prev *string
}

type cliCommand struct {
	name string
	description string
	callback func(cg *config) error
}

func commandExit(cg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	registry := getCommandsMap()
	for _, command := range registry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cg *config) error {
	path := "location-area"
	fullUrl := cg.PokeClient.BuildUrl(path)
	if cg.Next != nil {
		fullUrl = *cg.Next
	}

	locationsResult, err := cg.PokeClient.GetLocations(fullUrl)
	if err != nil {
		return err
	}

	results := locationsResult.Results
	cg.Next = locationsResult.Next
	cg.Prev = locationsResult.Previous
	for _, area := range results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandBMap(cg *config) error {
	path := "location-area"
	fullUrl := cg.PokeClient.BuildUrl(path)
	if cg.Prev != nil {
		fullUrl = *cg.Prev
	}

	locationsResult, err := cg.PokeClient.GetLocations(fullUrl)
	if err != nil {
		return err
	}
	results := locationsResult.Results
	cg.Next = locationsResult.Next
	cg.Prev = locationsResult.Previous
	for _, area := range results {
		fmt.Println(area.Name)
	}
	return nil
}

func startReading() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{pokeapi.NewClient(5 * time.Minute), nil, nil}
	for {
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)
		command := words[0]

		switch command {
		case "map":
			err := commandMap(&conf)
			if err != nil {
				break
			}
		case "bmap":
			err := commandBMap(&conf)
			if err != nil {
				break
			}
		case "exit":
			err := commandExit(&conf)
			if err != nil {
				break
			}
		case "help":
			err := commandHelp(&conf)
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
	}
}