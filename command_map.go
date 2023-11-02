package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config) error {
	response, err := config.pokeApiClient.GetLocationArea(config.NextLocation)
	if err != nil {
		return err
	}

	config.PreviousLocation = response.Previous
	config.NextLocation = response.Next

	for i, area := range response.Results {
		fmt.Printf("%v: %s\n", i+1, area.Name)
	}

	return nil
}

func commandMapB(config *config) error {
	if config.PreviousLocation == nil {
		return errors.New("Cannot move further back")
	}

	response, err := config.pokeApiClient.GetLocationArea(config.PreviousLocation)

	if err != nil {
		return err
	}

	config.PreviousLocation = response.Previous
	config.NextLocation = response.Next

	for i, area := range response.Results {
		fmt.Printf("%v: %s\n", i+1, area.Name)
	}

	return nil
}
