package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"gitlab.com/whyteoni/pokedexcli/internal/pokecache"
)

type pokedexConfig struct {
	next	string
	prev	string
	cache	*pokecache.Cache
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := pokedexConfig{
		next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20", 
		prev: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		cache: pokecache.NewCache(5 * time.Second),
	}

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil { panic("error reading from stdin")	}
		userInput := cleanInput(scanner.Text())
		if len(userInput) != 0 {
			command := userInput[0]
			if replCommand, ok := replCommands[command]; ok {
				if err := replCommand.callback(&config, userInput[1:]); err != nil {
					fmt.Printf("error when running command: %v\n", err)
					os.Exit(1)
				}
			} else { fmt.Println("Unknown command") }
		}
	}

}
