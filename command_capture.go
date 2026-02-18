package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)



func commandCatch(config *pokedexConfig, args []string) (err error) {
	var targetPokemon string
	var scaleTipper int = 0

	if len(args) == 0 { return fmt.Errorf("missing Pokemon to target") }
	if len(args) >= 1 { targetPokemon = args[0] }
	if len(args) >= 2 { 
		if i, err := strconv.Atoi(args[1]); err == nil {
			scaleTipper = i
		}
	}
 
	var url string = "https://pokeapi.co/api/v2/pokemon/" + targetPokemon
	var data Pokemon
	
	// Get caching lookup results and unmarshal to any{}
	content, err := cachedLookup(config.cache, url)
	if err != nil { 
		errString := err.Error()
		threeFromTheEnd := len(errString) - 3
		if errString[threeFromTheEnd:] == "404" {
			fmt.Printf("Oh no! We did not find %s. Please check the spelling and try again.\n", targetPokemon)
			return nil
		}
		return err 
	}
	err = json.Unmarshal(content, &data)
	if err != nil { return err }

	var offset int = (config.exp / 10) + len(config.capturedPokemon)
	var strike int = rand.Intn(data.BaseExperience + offset) + scaleTipper

	fmt.Printf("Throwing a Pokeball at %s... ", data.Name)

	// You get experience in all things!
	// Fresh capture: 100% base experience
	// Repeat capture: 50% base experience
	// Failed attempt: 10% base experience

	if strike >= data.BaseExperience {
		fmt.Println(" caught!")
		if _, exists := config.capturedPokemon[targetPokemon]; exists {
			fmt.Printf("  You previously captured %s!\n", data.Name)
			config.exp += data.BaseExperience / 2
			return nil
		}
		config.capturedPokemon[targetPokemon] = data
		config.exp += data.BaseExperience
	} else {
		fmt.Println(" escaped! Better luck next time.")
		config.exp += data.BaseExperience / 10
	}
	return nil
}
