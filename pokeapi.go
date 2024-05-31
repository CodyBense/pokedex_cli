package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Config struct {
    Count int 
    Next string 
    Previous any 
    Results []struct {
        Name string 
        Url string 
    } 
}

// Prints 20 locations on the pokemon map, subsequent calls should print the next 20
func pokemon_map(/* cfg *Config */) {
    res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
    if err != nil {
        log.Fatal(err)
    }
    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if res.StatusCode > 299 {
        log.Fatalf("Response failed with status code: %d and\nbody:%s\n", res.StatusCode, body)
    }
    if err != nil {
        log.Fatal(err)
    }
    /* dat := []byte(`{
        "count" : 1054,
        "next" : "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
        "previous" : null,
        "results":
        [{
            "name" : "canalave-city-area",
            "url" : "https://pokeapi.co/api/v2/location-area/1/"
        }]
    }`) */
    dat := []byte(body)
    
    cfg := Config{}

    errUm := json.Unmarshal(dat, &cfg)
    if errUm != nil {
        fmt.Println(errUm)
    }
    for location := range len(cfg.Results){
        fmt.Println(cfg.Results[location].Name)
    }
}

// Prints the previous 20 locaitons
func pokemon_mapb(cfg *Config) {
}
