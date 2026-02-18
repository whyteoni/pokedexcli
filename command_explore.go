package main

import (
	"encoding/json"
	"fmt"
)

func commandExplore(config *pokedexConfig, args []string) (err error) {
	var locationName string = args[0]
	var url string = "https://pokeapi.co/api/v2/location-area/" + locationName
	var data locationDetails

	// Get caching lookup results and unmarshal to any{}
	content, err := cachedLookup(config.cache, url)
	if err != nil { 
		errString := err.Error()
		threeFromTheEnd := len(errString) - 3
		if errString[threeFromTheEnd:] == "404" {
			fmt.Printf("Oh no! We did not find %s. Please check the spelling and try again.\n", locationName)
			return nil
		}
		return err 
	}
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
