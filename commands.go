package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func commandExit(cfg *config, args ...string) error {
    os.Exit(0)
    return nil
}

func commandHelp(cfg *config, args ...string) error {
    fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandClear(cfg *config, args ...string) error {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
    return nil
}

func commandMapf(cfg *config, args ...string) error {
    locationsResp, err := cfg.pokeapiClient.ListLocaitons(cfg.nextLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL = locationsResp.Next
    cfg.prevLocationsURL = locationsResp.Previous

    for _, loc := range locationsResp.Results {
        fmt.Println(loc.Name)
    }
    return nil
}

func commandMapb(cfg *config, args ...string) error {
    if cfg.prevLocationsURL == nil {
        return errors.New("you're on the first page")
    }

    locationResp, err := cfg.pokeapiClient.ListLocaitons(cfg.prevLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL = locationResp.Next
    cfg.prevLocationsURL = locationResp.Previous

    for _, loc := range locationResp.Results {
        fmt.Println(loc.Name)
    }
    return nil
}

func commandExplore(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("you must provide a location name")
    }

    name := args[0]
    location, err := cfg.pokeapiClient.GetLocation(name)
    if err != nil {
        return err
    }
    fmt.Printf("Exploring %s...\n", location.Name)
    fmt.Println("Found Pokemon: ")
    for _, enc := range location.PokemonEncounters {
        fmt.Printf(" - %s\n", enc.Pokemon.Name)
    }
    return nil
}
