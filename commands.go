package main

import (
	"fmt"
	"os"
	"errors"
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
	callback func(cg *config, args ...string) error
}

func commandExit(cg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	registry := getCommandsMap()
	for _, command := range registry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cg *config, args ...string) error {
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

func commandBMap(cg *config, args ...string) error {
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

func commandExplore(cg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Malformed location")
	}

	location := args[0]
	path := "location-area/" + location
	fullUrl := cg.PokeClient.BuildUrl(path)
	locationsDetails, err := cg.PokeClient.GetLocationDetails(fullUrl)
	if err != nil {
		return err
	}

	results := locationsDetails.PokemonEncounters
	for _, result := range results {
		fmt.Println(result.Pokemon.Name)
	}

	return nil
}
