package main

import (
    "os"
    "os/exec"
    "fmt"
)

func commandExit(cfg *Config) error {
    os.Exit(0)
    return nil
}

func commandHelp(cfg *Config) error {
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

func commandClear(cfg *Config) error {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
    return nil
}
