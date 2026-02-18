package main

import (
	"fmt"
)

func commandInfo(config *pokedexConfig, args []string) (err error) {
	var lenCaptured int = len(config.capturedPokemon)
	if lenCaptured != 0 {
		fmt.Printf("You have %d experience from %d captures!\n", config.exp, len(config.capturedPokemon))
		fmt.Println("Captured Pokemons:")
		for _, pokemon := range config.capturedPokemon {
			fmt.Printf("  - [ID: %d] %s (%d exp)\n", pokemon.ID, pokemon.Name, pokemon.BaseExperience)
		}
	} else {
		fmt.Printf("You have your initial %d experience, since you have not yet captured any Pokemon.\n", config.exp)
	}
	return nil
}
