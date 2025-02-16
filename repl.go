package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// The commands itself:
var registry map[string]cliCommand

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
}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, key := range registry {
		fmt.Printf("%v: %v\n", key.name, key.description)
	}
	return nil
}

func repl() {
	sc := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		curCom := strings.Fields(strings.ToLower(sc.Text()))
		if val, ok := registry[curCom[0]]; ok {
			val.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}
