package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		line := scanner.Text()

		words := strings.Fields(line)
		loweredWords := []string{}

		for _, w := range words {
			loweredWords = append(loweredWords, strings.ToLower(w))
		}

		fmt.Printf("Your command was: %v\n", loweredWords[0])
	}
}

func cleanInput(text string) []string {
	return strings.Split(text, " ")
}
