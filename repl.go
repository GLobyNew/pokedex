package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GLobyNew/pokedex/internal/pokecache"
)

const (
	locationAreasURL = "https://pokeapi.co/api/v2/location-area/"
)

type configStruct struct {
	current  string
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configStruct, *pokecache.Cache) error
}

// The commands itself:
var registry map[string]cliCommand

func setInitConfig() *configStruct {
	return &configStruct{
		current:  locationAreasURL + "?offset=0",
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

func commandExit(config *configStruct, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *configStruct, cache *pokecache.Cache) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, key := range registry {
		fmt.Printf("%v: %v\n", key.name, key.description)
	}
	return nil
}

func repl() {
	config := setInitConfig()
	cache := pokecache.NewCache(20 * time.Second)
	sc := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		curCom := strings.Fields(strings.ToLower(sc.Text()))
		if val, ok := registry[curCom[0]]; ok {
			val.callback(config, cache)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
