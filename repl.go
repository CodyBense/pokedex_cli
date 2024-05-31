package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Command struct
type cliCommand struct {
    name        string
    description string
    callBack    func(cfg *Config) error
}

// Repl loop: reads in input, finds command and runs it or prints unknown command if not existant
func startRepl() {
    cfg := Config{}
    reader := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        reader.Scan()
        words := cleanInput(reader.Text())
        if len(words) == 0 {
            continue
        }

        commandName := words[0]

        command, exits := getCommands()[commandName]
        if exits {
            err := command.callBack(&cfg)
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            fmt.Println("Unkown Command")
            continue
        }
    }
}

// Helper function to cleanup input
func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

// Gets commands
func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Displays a help message",
            callBack: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the pokedex",
            callBack: commandExit,
        },
        "clear": {
            name: "clear",
            description: "Clears the screen",
            callBack: commandClear,
        },
        "map": {
            name: "map",
            description: "Prints 20 locations, subsequent calls prints the next 20",
            callBack: commandMap,
        },
        "mapb": {
            name: "mapb",
            description: "Prints the previous 20 locations",
            callBack: commandMapb,
        },
    }
}
