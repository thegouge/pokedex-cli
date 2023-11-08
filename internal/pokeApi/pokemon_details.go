package pokeApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(pokemonName string) (PokemonData, error) {
	fullUrl := baseURL + "/pokemon/" + pokemonName

	pokemon := PokemonData{}

	requestCache, ok := c.cache.Get(fullUrl)

	if ok {
		jsonError := json.Unmarshal(requestCache, &pokemon)

		if jsonError != nil {
			return PokemonData{}, jsonError
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return PokemonData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return PokemonData{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonData{}, err
	}

	c.cache.Add(fullUrl, body)

	jsonError := json.Unmarshal(body, &pokemon)

	if jsonError != nil {
		return PokemonData{}, jsonError
	}

	return pokemon, nil
}
