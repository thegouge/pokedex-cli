package main

import (
	"bufio"
	"fmt"
	"github.com/thegouge/pokedex-cli/internal/pokeApi"
	"os"
	"strings"
)

type config struct {
	pokeApiClient    pokeApi.Client
	NextLocation     *string
	PreviousLocation *string
	Pokedex          map[string]pokeApi.PokemonData
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	config := config{
		pokeApiClient: pokeApi.NewClient(),
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		userCommand := cleaned[0]

		if len(cleaned) == 0 {
			continue
		}
		if command, exists := commands[userCommand]; exists {
			err := command.callback(&config, cleaned[1:])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("I don't understand that command")
		}
	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	output := strings.Fields(lowercase)
	return output
}
