package main

import (
	"fmt"
	"os"

	"github.com/TooCuteToo/repl-pokedex/internal/pokeapi"
	"github.com/TooCuteToo/repl-pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callBack    func(config *config) error
}

type config struct {
	pokeApiClient pokeapi.Client
	cache         pokecache.Cache
	nextUrl       *string
	prevUrl       *string
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callBack:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callBack:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display a current list of 20 locations",
			callBack:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Display a previous list of 20 locations",
			callBack:    commandMapBack,
		},
	}
	return commands
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for k, v := range getCommands() {
		fmt.Printf("%v: %v\n", k, v.description)
	}
	return nil
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *config) error {
	areasRepsone, err := config.pokeApiClient.GetAreas(config.nextUrl)
	if err != nil {
		return err
	}

	for _, v := range areasRepsone.Results {
		fmt.Println(v.Name)
	}

	config.nextUrl = areasRepsone.Next
	config.prevUrl = areasRepsone.Prev
	return nil
}

func commandMapBack(config *config) error {
	areasRepsone, err := config.pokeApiClient.GetAreas(config.prevUrl)
	if err != nil {
		return err
	}

	for _, v := range areasRepsone.Results {
		fmt.Println(v.Name)
	}

	config.nextUrl = areasRepsone.Next
	config.prevUrl = areasRepsone.Prev
	return nil
}
