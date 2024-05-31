package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

// helper function to cleanup input, removes white space and converts to lowercase
func cleanInput(text string) string {
    output := strings.TrimSpace(text)
    output = strings.ToLower(output)
    return output
}

