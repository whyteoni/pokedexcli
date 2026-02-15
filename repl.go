package main

import (
	"strings"
)

type cliCommand struct {
	name    string
	desc	string
	callback    func(*pokedexConfig) error
}

var replCommands = make(map[string]cliCommand)
func registerCommand(name, desc string, command func(*pokedexConfig) error) {
	replCommands[name] = cliCommand{
		name: name,
		desc: desc,
		callback: command,
	}
}

func init() {
	registerCommand("exit", "Closing the Pokedex... Goodbye!", commandExit)
	registerCommand("help", "Displays a help message", commandHelp)
	registerCommand("map", "Show the next 20 locations", commandMap)
	registerCommand("mapb", "Show the previous 20 locations", commandMapB)
}

func cleanInput(text string) (cleanText []string) {
	text = strings.ToLower(text)
	cleanText = strings.Fields(text)
	return
}
