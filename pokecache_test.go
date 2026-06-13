package main

import (
	"testing"
	"time"

	"github.com/ivanorribo/PokeAPI_bootdev/internal/pokecache"
)

func TestAddGet(t *testing.T) {
	cache := pokecache.NewCache(0)
	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)
	retrievedValue, found := cache.Get(key)

	if !found {
		t.Errorf("Expected to find key %q, but it was not found.", key)
	}

	if string(retrievedValue) != string(value) {
		t.Errorf("Expected value %q, got %q", value, retrievedValue)
	}
}

func TestReapLoop(t *testing.T) {
	ttl := 1 * time.Second
	cache := pokecache.NewCache(ttl)
	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	time.Sleep(2 * time.Second) // Wait for the reap loop to run

	_, found := cache.Get(key)
	if found {
		t.Errorf("Expected key %q to be reaped, but it was found.", key)
	}
}
