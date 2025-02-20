package main

import (
	"fmt"
	"strconv"

	"github.com/GLobyNew/pokedex/internal/argumentbuffer"
	"github.com/GLobyNew/pokedex/internal/pokecache"
)

func listPokemonTypes(pokemon Pokemon) string {
	pokemonTypes := ""
	for _, pkType := range pokemon.Types {
		pokemonTypes += " - " + pkType.Type.Name + "\n"
	}
	return pokemonTypes
}

func listPokemonStats(pokemon Pokemon) string {
	pokemonStats := ""
	for _, pkStat := range pokemon.Stats {
		pokemonStats += " -" + pkStat.Stat.Name + ": " + strconv.Itoa(pkStat.BaseStat) + "\n"
	}
	return pokemonStats
}

func commandInspect(config *configStruct, cache *pokecache.Cache, args *argumentbuffer.ArgumentBuff) error {
	pokemonName := args.GetArgs()[0]
	if _, catched := config.catchedPokemons[pokemonName]; !catched {
		fmt.Printf("%v is not in your inventory!\n", pokemonName)
		return nil
	}
	pokemon := config.catchedPokemons[pokemonName]
	fmt.Printf(`Name: %v
Height: %v
Weight: %v
Stats:
%v
Types:
%v`, pokemon.Name, pokemon.Height, pokemon.Weight, listPokemonStats(pokemon), listPokemonTypes(pokemon) )
	return nil
}