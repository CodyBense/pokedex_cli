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
    callBack    func() error
}

// Repl loop: reads in input, finds command and runs it or prints unknown command if not existant
func startRepl() {
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
            err := command.callBack()
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
        // "map": {
        //     name: "map",
        //     description: "Prints 20 locations, subsequent calls prints the next 20",
        //     callBack: commandMap,
        // },
    }
}

// Old
/*
var cliName string = "Pokdex"

// Prints prompt for every loop
func printPrompt() {
    fmt.Print(cliName, "> ")
}

// Prints for unkown functions
func printUnknown(text string) {
    fmt.Println(text, ": command not found")
}

// Prints all hardcoded functions
func displayHelp() {
    fmt.Printf(
        "Welcome to %v! These are the avilable commands: \n", cliName,
    )
    fmt.Println("map    -   Show 20 locations from the map")
    fmt.Println("mapb   -   Show the previous 20 locations")
    fmt.Println("help   -   Show available commands")
    fmt.Println("clear  -   Clear the terminal screen")
    fmt.Println("exit   -   Exit your connection")
}

// Clears the screen
func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

// attempts to recover from a bad command
func handleInvalidCommand(text string) {
    defer printUnknown(text)
}

// parses the given command
func handleCmd(text string) {
    handleInvalidCommand(text)
}

*/
