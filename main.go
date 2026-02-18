package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := NewConfig(10, 200)

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil { panic("error reading from stdin")	}
		userInput := cleanInput(scanner.Text())
		if len(userInput) != 0 {
			command := userInput[0]
			if replCommand, ok := replCommands[command]; ok {
				if err := replCommand.callback(config, userInput[1:]); err != nil {
					fmt.Printf("error when running command: %v\n", err)
					os.Exit(1)
				}
			} else { fmt.Println("Unknown command") }
		}
	}

}
