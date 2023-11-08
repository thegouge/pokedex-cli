package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, args []string) error {

	if len(args) == 0 {
		return errors.New("You need to specify an area!")
	}

	specifiedArea := args[0]
	fmt.Printf("exploring %s...\n", specifiedArea)

	response, err := config.pokeApiClient.GetAreaDetails(specifiedArea)

	if err != nil {
		return err
	}

	for _, result := range response.PokemonEncounters {
		fmt.Printf("- %s\n", result.Pokemon.Name)
	}

	return nil
}
