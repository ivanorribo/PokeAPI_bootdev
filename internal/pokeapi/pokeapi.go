package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ivanorribo/PokeAPI_bootdev/internal/pokecache"
)

type Client struct {
	client http.Client
	cache  *pokecache.Cache
}

type response struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name              string             `json:"name"`
	Url               string             `json:"url"`
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon"`
}

func NewClient() *Client {
	return &Client{
		client: http.Client{
			Timeout: time.Second * 10,
		},
		cache: pokecache.NewCache(time.Second * 60), // Set TTL to 60 seconds
	}
}

func (c *Client) GetLocation(url string) (*response, error) {
	var r response
	var body []byte
	if cachedResponse, found := c.cache.Get(url); found {
		body = cachedResponse
	} else {
		resp, err := c.client.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode > 299 {
			return nil, fmt.Errorf("error code: %d\n on body: %s", resp.StatusCode, body)
		}
		c.cache.Add(url, body)
	}

	err := json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *Client) PokemonEncounters(url string) ([]pokemonEncounter, error) {
	var body []byte
	if cachedResponse, found := c.cache.Get(url); found {
		body = cachedResponse
	} else {
		resp, err := c.client.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode > 299 {
			return nil, fmt.Errorf("error code: %d\n on body: %s", resp.StatusCode, body)
		}
		c.cache.Add(url, body)
	}
	location := location{}
	err := json.Unmarshal(body, &location)
	if err != nil {
		return nil, err
	}
	return location.PokemonEncounters, nil
}
