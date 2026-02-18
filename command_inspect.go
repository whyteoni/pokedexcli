package main

import (
	"fmt"
)

func commandInspect(config *pokedexConfig, args []string) error {
	if len(args) == 0 { return fmt.Errorf("missing target Pokemon declaration")	}
	var targetPokemon string = args[0]

	pokemon, exists := config.capturedPokemon[targetPokemon]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		name := stat.Stat.Name
		value := stat.BaseStat
		fmt.Printf(" - %s: %d\n", name, value)
	}

	fmt.Println("Types:")
	for _, ptype := range pokemon.Types {
		fmt.Printf(" - %s\n", ptype.Type.Name)
	}

	return nil
}
