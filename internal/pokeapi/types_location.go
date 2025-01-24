package pokeapi

type LocationAreasResponse struct {
	Count   int            `json:"count"`
	Next    *string        `json:"next"`
	Prev    *string        `json:"previous"`
	Results []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
}

type LocationAreaDetailResponse struct {
	Name     string `json:"name"`
	Index    int    `json:"game_index"`
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
		}
	} `json:"pokemon_encounters"`
}
