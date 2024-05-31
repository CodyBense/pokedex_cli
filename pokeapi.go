package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TODO: *UPDATE to new format, found in repl.go* map: get Config arg working, work on subsquent call feature; mapb: all of it

// Struct for response
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
func commandMap(cfg *Config) error{
    // GET api request and check for errors
    res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
    if err != nil {
        log.Fatal(err)
        return err
    }
    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if res.StatusCode > 299 {
        log.Fatalf("Response failed with status code: %d and\nbody:%s\n", res.StatusCode, body)
    }
    if err != nil {
        log.Fatal(err)
        return err
    }

    // Slice GET request into Config struct
    dat := []byte(body)
    
    // cfg := Config{}

    errUm := json.Unmarshal(dat, &cfg)
    if errUm != nil {
        fmt.Println(errUm)
    }

    // Prints location names
    fmt.Println()
    for location := range len(cfg.Results){
        fmt.Println(cfg.Results[location].Name)
    }
    return nil
}

// Prints the previous 20 locaitons
func commandMapb(cfg *Config) {
}
