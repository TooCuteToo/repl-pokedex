package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

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

		err := v.callBack()
		if err != nil {
			fmt.Printf("there was an error: %v", err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Split(text, " ")
}
