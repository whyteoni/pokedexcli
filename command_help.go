package main

import "fmt"

func commandHelp(config *pokedexConfig) (err error) {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range replCommands {
		fmt.Printf("%s: %s\n", command.name, command.desc)
	}
	return
}
