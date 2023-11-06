package pokeApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAreaDetails(areaName string) (SingleArea, error) {
	fullURL := baseURL + "/location-area/" + areaName

	areaDetails := SingleArea{}

	requestCache, ok := c.cache.Get(fullURL)

	if ok {
		jsonError := json.Unmarshal(requestCache, &areaDetails)

		if jsonError != nil {
			return SingleArea{}, jsonError
		}

		return areaDetails, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return SingleArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return SingleArea{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return SingleArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return SingleArea{}, err
	}

	c.cache.Add(fullURL, body)

	jsonError := json.Unmarshal(body, &areaDetails)

	if jsonError != nil {
		return SingleArea{}, err
	}

	return areaDetails, nil
}
