package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/GLobyNew/pokedex/internal/pokecache"
)

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *configStruct) error {
	// TODO: implement cache usage
	res, err := http.Get(config.next)
	if err != nil {
		return errors.New("failed Get request in commandMap func")
	}
	var locAreas locationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locAreas); err != nil {
		return errors.New("failed to decode response")
	}

	config.next = locAreas.Next
	config.previous = locAreas.Previous

	for _, area := range locAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(config *configStruct) error {
	if config.previous == "" {
		fmt.Println("Already at first page")
		return nil
	}
	res, err := http.Get(config.previous)
	if err != nil {
		return errors.New("failed Get request in commandMapb func")
	}
	var locAreas locationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locAreas); err != nil {
		return errors.New("failed to decode response")
	}

	config.next = locAreas.Next
	config.previous = locAreas.Previous

	for _, area := range locAreas.Results {
		fmt.Println(area.Name)
	}

	return nil
}
