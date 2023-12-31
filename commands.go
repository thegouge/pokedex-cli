package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays the help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists 20 pokemon world locations, each subsequent call will give you the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 pokemon world locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <area name>",
			description: "List out all the pokemon that appear in a particular area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Throw a pokeball out to catch the specified pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Get detailed information about a pokemon you've caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all of the Pokemon you've caught",
			callback:    commandPokedex,
		},
	}
}
