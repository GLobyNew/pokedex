package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/GLobyNew/pokedex/internal/argumentbuffer"
	"github.com/GLobyNew/pokedex/internal/pokecache"
	"github.com/GLobyNew/pokedex/internal/requests"
)

type PokemonsLocsResp struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func bytesToPokemonLocsResp(data []byte) (PokemonsLocsResp, error) {
	var pokemonLocs PokemonsLocsResp
	if err := json.Unmarshal(data, &pokemonLocs); err != nil {
		return PokemonsLocsResp{}, errors.New("failed to decode response")
	}
	return pokemonLocs, nil
}

// func getPokemonNameByID(id int) (string, error) {
// 	URL := pokemonURL + strconv.Itoa(id)
// 	resp, err := requests.MakeGETRequest(URL)
// 	if err != nil {
// 		return "", err
// 	}
// 	data := make(map[string]interface{})
// 	if err := json.Unmarshal(resp, &data); err != nil {
// 		return "", errors.New("failed to decode response")
// 	}
// 	if name, ok := data["name"].(string); ok {
// 		return name, nil
// 	} else {
// 		return "", errors.New("no name found when parsing pokemon name")
// 	}
// }

func cacheList(cache *pokecache.Cache, URL string, list string) {
	cache.Add(URL, []byte(list))
}

func printFromCache(cache *pokecache.Cache, URL string) {
	list, _ := cache.Get(URL)
	fmt.Println(string(list))
}

func requestAndPrintPokemon(data []byte, cache *pokecache.Cache, URL string) error {
	pokemonLocs, err := bytesToPokemonLocsResp(data)
	pokemonList := ""
	if err != nil {
		return err
	}
	for _, PokemonEncounter := range pokemonLocs.PokemonEncounters {
		pokemonList += "- " + PokemonEncounter.Pokemon.Name + "\n"
	}
	fmt.Println("Found Pokemon:")
	fmt.Println(pokemonList)
	cacheList(cache, URL, pokemonList)
	return nil
}

func commandExplore(config *configStruct, cache *pokecache.Cache, args *argumentbuffer.ArgumentBuff) error {
	for _, arg := range args.GetArgs() {
		fmt.Printf("Exploring %v...\n", arg)
		URL := locationAreasURL + arg
		if _, exist := cache.Get(URL); exist {
			printFromCache(cache, URL)
			return nil
		}

		jsonData, err := requests.MakeGETRequest(URL)
		if err != nil {
			return err
		}
		requestAndPrintPokemon(jsonData, cache, URL)
	}
	return nil
}
