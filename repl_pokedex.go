package main

import (
	"fmt"

	"github.com/GLobyNew/pokedex/internal/argumentbuffer"
	"github.com/GLobyNew/pokedex/internal/pokecache"
)

func commandPokedex(config *configStruct, cache *pokecache.Cache, args *argumentbuffer.ArgumentBuff) error {
	if len(config.catchedPokemons) == 0 {
		fmt.Println("There is no caught Pokemons!")
		return nil
	}

	for _, pokemon := range config.catchedPokemons {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}
