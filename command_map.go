package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type locationAreaLookup struct {
	Count	int		`json:"count"`
	Next	*string	`json:"next"`
	Prev	*string	`json:"previous"`
	Results	[]struct{
		Name	string	`json:"name"`
		Url		string	`json:"url"`
	}	`jon:"results"`
}

func getLocationArea(config *pokedexConfig, url string) (data locationAreaLookup, err error) {
	resp, err := http.Get(url)
	if err != nil { return }
	defer resp.Body.Close()
	if resp.StatusCode > 299 { 
		err = fmt.Errorf("response failed with status code: %d", resp.StatusCode)
		return  
	}

	// Convert http.response into JSON
	body, err := io.ReadAll(resp.Body)
	if err != nil { return }
	err = json.Unmarshal(body, &data)
	if err != nil { return }

	// Update global config
	if data.Next != nil { config.next = *data.Next }
	if data.Prev != nil { config.prev = *data.Prev }

	return
}

func commandMap(config *pokedexConfig) (err error) {
	data, err := getLocationArea(config, config.next)
	if err != nil { return }
	for _, location := range data.Results {	fmt.Println(location.Name) }
	return
}

func commandMapB(config *pokedexConfig) (err error) {
	data, err := getLocationArea(config, config.prev)
	if err != nil { return }
	for _, location := range data.Results {	fmt.Println(location.Name) }
	return
}
