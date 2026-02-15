package main

import (
	"fmt"
	"os"
)

func commandExit(config *pokedexConfig) (err error) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return
}
