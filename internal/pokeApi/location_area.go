package pokeApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type AreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type JSONParsing struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []AreaResult `json:"results"`
}

func GetLocationArea(offset int) (next string, previous string, results []AreaResult, error error) {
	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset))
	if err != nil {
		return "", "", nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return "", "", nil, errors.New(fmt.Sprintf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body))
	}
	if err != nil {
		return "", "", nil, err
	}

	areaData := JSONParsing{}
	jsonError := json.Unmarshal(body, &areaData)
	if jsonError != nil {
		return "", "", nil, jsonError
	}

	return areaData.Next, areaData.Previous, areaData.Results, nil
}
