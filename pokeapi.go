package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TODO: *UPDATE to new format, found in repl.go* map: get Config arg working, work on subsquent call feature; mapb: all of it

// Struct for response
type Config struct {
    Count       int 
    Next        string 
    Previous    string 
    Results     []struct {
        Name string 
        Url  string 
    } 
}

// Prints 20 locations on the pokemon map, subsequent calls should print the next 20
func commandMap(cfg *Config) error{
    var helperResults []byte

    if cfg.Next == "" {
        helperResults, _ = helperMap("https://pokeapi.co/api/v2/location-area/")
    } else {
        helperResults, _ = helperMap(cfg.Next)
    }

    // Slice GET request into Config struct
    dat := []byte(helperResults)

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
func commandMapb(cfg *Config) error {
    if cfg.Previous == ""{
        return errors.New("No previous pages")
    }
    helperResults, _ := helperMap(cfg.Previous)

    // Slice GET request inot Config struct
    dat := []byte(helperResults)

    errUm := json.Unmarshal(dat, &cfg)
    if errUm != nil {
        fmt.Println(errUm)
    }

    // Prints locaiton names
    fmt.Println()
    for location := range len(cfg.Results){
        fmt.Println(cfg.Results[location].Name)
    }
    return nil
}

// Helper func for map func
func helperMap(url string) ([]byte, error) {
    // GET api request and check for errors
    // res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
    res, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if res.StatusCode > 299 {
        log.Fatalf("Response failed with status code: %d and\nbody:%s\n", res.StatusCode, body)
    }
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    return body, nil
}
