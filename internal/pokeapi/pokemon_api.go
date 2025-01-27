package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	fmt.Println("URL: ", url)
	fmt.Println("Proceed to get the location list in cache")

	cacheVal, ok := c.cache.Get(url)
	if ok {
		var pokemonResponse PokemonResponse
		err := json.Unmarshal(cacheVal, &pokemonResponse)
		if err != nil {
			return PokemonResponse{}, err
		}
		return pokemonResponse, nil
	}

	fmt.Println("Proceed to make request to get pokemon")
	res, err := c.httpClient.Get(url)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return PokemonResponse{}, nil
	}

	var pokemonResponse PokemonResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	fmt.Println("Add to cache")
	jsonData, err := json.Marshal(pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}
	c.cache.Add(url, jsonData)
	return pokemonResponse, nil
}
