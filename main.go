package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	config := config{offset: 0}

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
