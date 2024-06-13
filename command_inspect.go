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
            fmt.Println("Name:", pokemon.Name)
            fmt.Println("Height:", pokemon.Height)
            fmt.Println("Weight:", pokemon.Weight)
            fmt.Println("Stats:")
            for _, stat := range pokemon.Stats {
                fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
            }
            fmt.Println("Types:")
            for _, typeInfo := range pokemon.Types {
                fmt.Printf(" -%s\n", typeInfo.Type.Name)
            }
        }
        return nil
    }
    return errors.New("pokemon not caught")
}
