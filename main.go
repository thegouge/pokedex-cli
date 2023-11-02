package main

import (
	"bufio"
	"fmt"
	"internal/pokeApi"
	"os"
	"strings"
)

type config struct {
	pokeApiClient    pokeApi.Client
	NextLocation     *string
	PreviousLocation *string
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
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
				return
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
