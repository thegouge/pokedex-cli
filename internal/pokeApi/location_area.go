package pokeApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(pageURL *string) (AreaResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	areaData := AreaResponse{}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return areaData, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return areaData, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return areaData, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return areaData, err
	}

	jsonError := json.Unmarshal(body, &areaData)

	if jsonError != nil {
		return AreaResponse{}, err
	}

	return areaData, nil
}
