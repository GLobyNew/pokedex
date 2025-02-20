package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/GLobyNew/pokedex/internal/argumentbuffer"
	"github.com/GLobyNew/pokedex/internal/pokecache"
	"github.com/GLobyNew/pokedex/internal/requests"
)

const (
	pokemonEndpointURL = "https://pokeapi.co/api/v2/pokemon/"
)

func bytesToPokemon(data []byte) (Pokemon, error) {
	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, errors.New("failed to decode response")
	}
	return pokemon, nil
}

func commandCatch(config *configStruct, cache *pokecache.Cache, args *argumentbuffer.ArgumentBuff) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pokemonName := args.GetArgs()[0]
	URL := pokemonEndpointURL + pokemonName
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	pokemonData, err := requests.MakeGETRequest(URL)
	if err != nil {
		return err
	}
	pokemon, err := bytesToPokemon(pokemonData)
	if err != nil {
		return err
	}

	chanceNumber := r.Intn(1000)
	if chanceNumber > pokemon.BaseExperience {
		config.catchedPokemons[pokemonName] = pokemon
		fmt.Printf("%v was caught!\n", pokemonName)
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil
}
