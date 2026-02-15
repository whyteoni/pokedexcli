package main

import (
	"fmt"
	"bufio"
	"os"
)

type pokedexConfig struct {
	next	string
	prev	string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := pokedexConfig{
		next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20", 
		prev: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	}

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil { panic("error reading from stdin")	}
		userInput := scanner.Text()
		command := cleanInput(userInput)[0]
		if replCommand, ok := replCommands[command]; ok {
			if err := replCommand.callback(&config); err != nil {
				fmt.Printf("error when running command: %v\n", err)
				os.Exit(1)
			}
		} else { fmt.Println("Unknown command") }
	}

}
