package main

import (
    "errors"
    "fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

    name := args[0]
    pokemon, err := cfg.pokeapiClient.GetPokemon(name)
    if err != nil {
        return err
    }
    for _, p := range cfg.caughtPokemon {
        if p.Name == name {
            fmt.Printf("Name: %s\n", pokemon.Name)
            fmt.Printf("Height: %d\n", pokemon.Height)
            fmt.Printf("Weight: %d\n", pokemon.Weight)
        }
        return nil
    }
    return errors.New("pokemon not caught")
}
