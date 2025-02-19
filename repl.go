package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/GLobyNew/pokedex/internal/argumentbuffer"
	"github.com/GLobyNew/pokedex/internal/pokecache"
)

const (
	locationAreasURL = "https://pokeapi.co/api/v2/location-area/"
)

type configStruct struct {
	next     string
	previous string
}

type cliCommandNoArgs struct {
	name        string
	description string
	callback    func(*configStruct, *pokecache.Cache) error
}

type cliCommandWithArgs struct {
	name        string
	description string
	callback    func(*pokecache.Cache, *argumentbuffer.ArgumentBuff) error
}

// The commands itself:
var registryNoArgs map[string]cliCommandNoArgs
var registryWithArgs map[string]cliCommandWithArgs

func setInitConfig() *configStruct {
	return &configStruct{
		next:     locationAreasURL + "?offset=0",
		previous: locationAreasURL + "?offset=0",
	}
}

func init() {
	registryNoArgs = map[string]cliCommandNoArgs{
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
	registryWithArgs = map[string]cliCommandWithArgs{
		"explore": {
			name:        "explore",
			description: "List what Pokemons can be found in desired location",
			callback:    commandExplore,
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
	for _, key := range registryNoArgs {
		fmt.Printf("%v: %v\n", key.name, key.description)
	}
	return nil
}

func repl() {
	config := setInitConfig()
	cache := pokecache.NewCache(20 * time.Second)
	argBuf := argumentbuffer.NewArgumentBuff()
	sc := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		curCom := cleanInput(sc.Text())
		if len(curCom) > 1 {
			if val, ok := registryWithArgs[curCom[0]]; ok {
				argBuf.Set(curCom)
				val.callback(cache, argBuf)
			} else {
				fmt.Println("Unknown command")
			}
		} else if val, ok := registryNoArgs[curCom[0]]; ok {
			val.callback(config, cache)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
