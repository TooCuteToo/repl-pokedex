package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TooCuteToo/repl-pokedex/internal/pokeapi"
)

const (
	timeout       = 30 * time.Second
	cacheInterval = 5 * time.Minute
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	config := config{
		pokeApiClient: pokeapi.NewClient(timeout, cacheInterval),
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		line := scanner.Text()

		words := strings.Fields(line)
		loweredWords := []string{}

		for _, w := range words {
			loweredWords = append(loweredWords, strings.ToLower(w))
		}

		v, ok := commands[loweredWords[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		sv := ""
		if len(loweredWords) > 1 {
			sv = loweredWords[1]
		}

		err := v.callBack(&config, sv)
		if err != nil {
			fmt.Printf("there was an error: %v", err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Split(text, " ")
}
