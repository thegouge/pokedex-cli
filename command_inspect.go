package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, args []string) error {
	if len(args) == 0 {
		return errors.New("You can't inspect nothing!")
	}

	pokemon, ok := config.Pokedex[args[0]]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %vdm\n", pokemon.Height)
	fmt.Printf("Weight: %vhg\n", pokemon.Weight)
	fmt.Println("Base Stats:")
	for _, statStruct := range pokemon.Stats {
		fmt.Printf("    - %s: %v\n", statStruct.Stat.Name, statStruct.BaseStat)
	}
	fmt.Println("Types: ")
	for _, typeStruct := range pokemon.Types {
		fmt.Printf("    - %s\n", typeStruct.Type.Name)
	}

	return nil
}
