package internal

import (
	"encoding/json"
)

func (c *Client) GetAreas(pageUrl *string) (LocationAreasResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

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
	return areasResponse, nil
}
