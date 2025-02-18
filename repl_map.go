package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/GLobyNew/pokedex/internal/pokecache"
)

type direction = int

const (
	next direction = iota
	previous
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

func MakeRequest(config *configStruct, URL string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return []byte(resBody), nil
}

func bytesTolocationAreas(data []byte) (locationAreas, error) {
	var locAreas locationAreas
	if err := json.Unmarshal(data, &locAreas); err != nil {
		return locationAreas{}, errors.New("failed to decode response")
	}
	return locAreas, nil
}

func printResults(data []byte) error {
	locAreas, err := bytesTolocationAreas(data)
	if err != nil {
		return err
	}
	for _, area := range locAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func setConfigPages(config *configStruct, data []byte) error {
	locAreas, err := bytesTolocationAreas(data)
	if err != nil {
		return err
	}
	config.next = locAreas.Next
	config.previous = locAreas.Previous
	fmt.Println(config.next)
	fmt.Println(config.previous)
	return nil
}

func newPageResult(config *configStruct, cache *pokecache.Cache, d direction) error {
	// TODO Fix cache bug
	var URL string
	switch d {
	case next:
		URL = config.next
	case previous:
		URL = config.previous
	default:
		return errors.New("wrong direction in func newPageResult")
	}
	if URL == "" {
		fmt.Println("You already on a first page")
		return nil
	}

	if data, exist := cache.Get(URL); exist {
		defer setConfigPages(config, data)
		printResults(data)
		return nil
	}
	jsonData, err := MakeRequest(config, URL)
	if err != nil {
		return err
	}
	cache.Add(config.next, jsonData)
	printResults(jsonData)
	defer setConfigPages(config, jsonData)

	return nil
}

func commandMap(config *configStruct, cache *pokecache.Cache) error {
	err := newPageResult(config, cache, next)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(config *configStruct, cache *pokecache.Cache) error {
	err := newPageResult(config, cache, previous)
	if err != nil {
		return err
	}
	return nil
}
