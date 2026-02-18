package main

import (
	"time"

	"gitlab.com/whyteoni/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name    string
	desc	string
	callback    func(*pokedexConfig, []string) error
}

type pokedexConfig struct {
	next			string
	prev			string
	cache			*pokecache.Cache
	capturedPokemon	map[string]Pokemon
	exp				int
}

func NewConfig(duration int, initialExp int) (config *pokedexConfig){
	config = &pokedexConfig{
		next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20", 
		prev: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		cache: pokecache.NewCache(time.Duration(duration) * time.Second),
		capturedPokemon: make(map[string]Pokemon),
		exp: initialExp,
	}
	return
}

// API Endpoint based structs

type locationAreaLookup struct {
	Count	int		`json:"count"`
	Next	*string	`json:"next"`
	Prev	*string	`json:"previous"`
	Results	[]struct{
		Name	string	`json:"name"`
		Url		string	`json:"url"`
	}	`jon:"results"`
}

type locationDetails struct {
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Name  string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	BaseExperience 	int `json:"base_experience"`
	Height    		int `json:"height"`
	ID              int    `json:"id"`
	Name          	string `json:"name"`
	Stats []struct {
		BaseStat 	int `json:"base_stat"`
		Effort   	int `json:"effort"`
		Stat     	struct {
			Name 	string `json:"name"`
			URL  	string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot 		int `json:"slot"`
		Type struct {
			Name 	string `json:"name"`
			URL  	string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight 			int `json:"weight"`
}
