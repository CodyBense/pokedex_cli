package main

import (
    "errors"
    "fmt"
)

func commandPokedex(cfg *config, args ...string) error {
    pokemon:= cfg.caughtPokemon
    if len(pokemon) == 0 {
        return errors.New("No pokemon have been caught")
    }
    fmt.Println("Your Pokedex:")
    for _, p := range pokemon {
        fmt.Printf(" - %s\n", p.Name)
    }
    return nil
}
