package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	locationAreasURL = "https://pokeapi.co/api/v2/location-area/"
)

type configStruct struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configStruct) error
}

// The commands itself:
var registry map[string]cliCommand

func setInitConfig() *configStruct {
	return &configStruct{
		next:     locationAreasURL + "?offset=0",
		previous: locationAreasURL + "?offset=0",
	}
}

func init() {
	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Print help",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show 20 new location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show 20 previous location areas",
			callback:    commandMapb,
		},
	}

}

func commandExit(config *configStruct) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *configStruct) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, key := range registry {
		fmt.Printf("%v: %v\n", key.name, key.description)
	}
	return nil
}

func repl() {
	config := setInitConfig()
	sc := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		curCom := strings.Fields(strings.ToLower(sc.Text()))
		if val, ok := registry[curCom[0]]; ok {
			val.callback(config)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
