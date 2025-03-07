package main

import (
	"fmt"
	"os"
	"errors"
	"math/rand"
	"strings"
)

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

func commandCatch(cg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Malformed location")
	}

	name := args[0]
	path := "pokemon/" + name
	fullUrl := cg.PokeClient.BuildUrl(path)
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemonData, err := cg.PokeClient.GetPokemon(fullUrl)
	if err != nil {
		return err
	}
	r := rand.Intn(pokemonData.BaseExperience)
	threshold := 40
	if r > threshold {
		fmt.Printf("%s escaped!\n", name)
	} else {
		cg.Pokedex[name] = pokemonData
		fmt.Printf("%s was caught!\n", name)
	}
	return nil
}

func commandInspect(cg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Malformed location")
	}

	name := args[0]
	
	p, ok := cg.Pokedex[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	var stats, types strings.Builder
	for _, stat := range p.Stats {
		fmt.Fprintf(&stats, "  - %s: %2d\n", stat.Stat.Name, stat.BaseStat)
	}
	for _, typ := range p.Types {
		fmt.Fprintf(&types, "  - %s\n", typ.Type.Name)
	}
	fmt.Printf("Name: %s\nHeight: %2d\nWeight: %2d\nStats:\n%sTypes:\n%s", p.Name, p.Height, p.Weight, stats.String(), types.String())
	return nil
}

func commandPokedex(cg *config, args ...string) error {
	n := len(cg.Pokedex)
	if n == 0 {
		return fmt.Errorf("you have not caught pokemons")
	}

	for _, pokemon := range cg.Pokedex {
		fmt.Println(pokemon.Name)
	}

	return nil
}