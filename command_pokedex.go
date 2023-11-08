package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *config, args []string) error {
	if len(config.Pokedex) == 0 {
		return errors.New("You haven't caught any Pokemon yet!")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.Pokedex {
		fmt.Printf("  - %v\n", pokemon.Name)
	}

	return nil
}
