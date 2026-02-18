package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"gitlab.com/whyteoni/pokedexcli/internal/pokecache"
)

var replCommands = make(map[string]cliCommand)
func registerCommand(name, desc string, command func(*pokedexConfig, []string) error) {
	replCommands[name] = cliCommand{
		name: name,
		desc: desc,
		callback: command,
	}
}

func init() {
	registerCommand("catch", "Attempt to capture a pokemon.", commandCatch)
	registerCommand("exit", "Closes the pokedex.", commandExit)
	registerCommand("explore", "Explore a given location. Takes a location name or ID.", commandExplore)
	registerCommand("help", "Displays a help message.", commandHelp)
	registerCommand("info", "Show information about your session.", commandInfo)
	registerCommand("inspect", "Inspect a cpatured pokemon.", commandInspect)
	registerCommand("map", "Show the next 20 locations.", commandMap)
	registerCommand("mapb", "Show the previous 20 locations.", commandMapB)
	
}

func cleanInput(text string) (cleanText []string) {
	text = strings.ToLower(text)
	cleanText = strings.Fields(text)
	return
}

func cachedLookup(cache *pokecache.Cache, url string) (content []byte, err error) {
	var exists bool
	if content, exists = cache.Get(url); !exists {	
		var err error
		resp, err := http.Get(url)
		if err != nil { return nil, err }
		defer resp.Body.Close()
		if resp.StatusCode > 299 { 
			err = fmt.Errorf("response failed with status code: %d", resp.StatusCode)
			return nil, err
		}
		// Convert http.response into []byte
		content, err = io.ReadAll(resp.Body)
		if err != nil { return nil, err }

		// Cache body of get response
		cache.Add(url, content)
	}
	return content, nil
}
