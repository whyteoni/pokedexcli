package main

import (
	"encoding/json"
	"fmt"
)

func getLocationArea(config *pokedexConfig, url string) (locationAreaLookup, error) {
	var data locationAreaLookup

	// Get caching lookup results
	content, err := cachedLookup(config.cache, url)
	if err != nil { return locationAreaLookup{}, err }

	// Convert []byte contents into locationAreaLookup{}
	err = json.Unmarshal(content, &data)
	if err != nil { return locationAreaLookup{}, err }

	// Update global config
	if data.Next != nil { config.next = *data.Next }
	if data.Prev != nil { config.prev = *data.Prev }

	return data, nil
}

func commandMap(config *pokedexConfig, args []string) (err error) {
	data, err := getLocationArea(config, config.next)
	if err != nil { return }
	for _, location := range data.Results {	fmt.Println(location.Name) }
	return
}

func commandMapB(config *pokedexConfig, args []string) (err error) {
	data, err := getLocationArea(config, config.prev)
	if err != nil { return }
	for _, location := range data.Results {	fmt.Println(location.Name) }
	return
}
