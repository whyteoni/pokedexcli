package main

import (
	"fmt"
	"testing"
	"time"

	"gitlab.com/whyteoni/pokedexcli/internal/pokecache"
)


func TestAddGet(t *testing.T) {
	const interval = 10 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
		{
			key: "nonURLtestKey",
			val: []byte("nonURLtestVal"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			if val, exists := cache.Get(c.key); exists {
				if string(val) != string(c.val) {
					t.Errorf("expected to find value")
					return
				}
			} else {
				t.Errorf("expected to find key")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 1*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
