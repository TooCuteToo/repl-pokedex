package main

import (
	"fmt"
	"os"

	"github.com/TooCuteToo/repl-pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callBack    func(config *config, value string) error
}

type config struct {
	pokeApiClient pokeapi.Client
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
		"explore": {
			name:        "exlpre",
			description: "Display a list of pokemon encounters in the location",
			callBack:    commandExpore,
		},
	}
	return commands
}

func commandHelp(config *config, value string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for k, v := range getCommands() {
		fmt.Printf("%v: %v\n", k, v.description)
	}
	return nil
}

func commandExit(config *config, value string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *config, value string) error {
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

func commandMapBack(config *config, value string) error {
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

func commandExpore(config *config, value string) error {
	fmt.Printf("Exploring %v...\n", value)
	areaDetailResponse, err := config.pokeApiClient.ExploreArea(value)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, v := range areaDetailResponse.Pokemons {
		fmt.Printf("- %v\n", v.Pokemon.Name)
	}

	return nil
}
