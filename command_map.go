package main

import (
	"errors"
	"fmt"
	"internal/pokeApi"
)

func commandMap(config *config) error {
	_, _, results, err := pokeApi.GetLocationArea(config.offset)

	config.offset += 20

	if err != nil {
		return err
	}

	for _, result := range results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapB(config *config) error {
	if config.offset == 0 {
		return errors.New("unable to go back any further")
	}
	config.offset -= 20

	_, _, results, err := pokeApi.GetLocationArea(config.offset)

	if err != nil {
		return err
	}

	for _, result := range results {
		fmt.Println(result.Name)
	}

	return nil
}
