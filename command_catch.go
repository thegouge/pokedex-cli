package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const CATCH_THRESHOLD = 30

func commandCatch(config *config, args []string) error {
	if len(args) == 0 {
		return errors.New("You threw your Pokeball at the air... nothing happened")
	}

	pokemonToCatch := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonToCatch)

	pokemonData, err := config.pokeApiClient.GetPokemonDetails(pokemonToCatch)

	if err != nil {
		return err
	}

	catchRoll := rand.Intn(pokemonData.BaseExperience)

	if catchRoll <= CATCH_THRESHOLD {
		fmt.Printf("%s was caught!\n", pokemonToCatch)
		config.Pokedex[pokemonToCatch] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonToCatch)
	}

	return nil
}
