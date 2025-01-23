package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetAreas(pageUrl *string) (LocationAreasResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	fmt.Println("URL: ", url)
	fmt.Println("Proceed to get the location list in cache")
	cacheVal, ok := c.cache.Get(url)
	if ok {
		var areasResponse LocationAreasResponse
		err := json.Unmarshal(cacheVal, &areasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return areasResponse, nil
	}

	fmt.Println("Proceed to make request to get location list")
	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	var areasResponse LocationAreasResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	fmt.Println("Add to cache")
	jsonData, err := json.Marshal(areasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	c.cache.Add(url, jsonData)
	return areasResponse, nil
}
