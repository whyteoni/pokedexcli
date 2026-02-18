package main

import (
	"fmt"
	"strings"
)

func commandHelp(config *pokedexConfig, args []string) (err error) {
	const padding string = " "
	var offset int = 0

	fmt.Printf("\nWelcome to the Pokedex!\n\nCommands:\n\n")

	for _, command := range replCommands {
		if len(command.name) > offset {
			offset = len(command.name)
		}
	}

	for _, command := range replCommands {
		filler := strings.Repeat(padding, offset - len(command.name))
		fmt.Printf("%s%s: %s\n", filler, command.name, command.desc)
	}
	fmt.Println()
	return
}
