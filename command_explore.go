package main

import (
	"encoding/json"
	"fmt"
)

type locationDetails struct {
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Name  string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}


func commandExplore(config *pokedexConfig, args []string) (err error) {
	var locationName string = args[0]
	var url string = "https://pokeapi.co/api/v2/location-area/" + locationName
	var data locationDetails

	// Get caching lookup results and unmarshal to any{}
	content, err := cachedLookup(config.cache, url)
	if err != nil { return err }
	err = json.Unmarshal(content, &data)
	if err != nil { return err }

	fmt.Printf("Exploring %s...\n", locationName)

	if len(data.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon Found!!")
		return
	}
	
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return
}
