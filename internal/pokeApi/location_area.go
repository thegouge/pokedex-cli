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

	requestCache, ok := c.cache.Get(fullURL)

	if ok {
		jsonError := json.Unmarshal(requestCache, &areaData)

		if jsonError != nil {
			return AreaResponse{}, jsonError
		}

		return areaData, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return AreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return AreaResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return AreaResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return AreaResponse{}, err
	}

	c.cache.Add(fullURL, body)

	jsonError := json.Unmarshal(body, &areaData)

	if jsonError != nil {
		return AreaResponse{}, jsonError
	}

	return areaData, nil
}
