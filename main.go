package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
)

func main() {
    // Hardcoded commands
    commands := map[string]interface{} {
        "help": displayHelp,
        "clear": clearScreen,
    }

    // Begin repl loop
    reader := bufio.NewScanner(os.Stdin)
    printPrompt()
    for reader.Scan() {
        text := cleanInput(reader.Text())
        if command, exists := commands[text]; exists {
            // call hardcoded function
            command.(func())()
        } else if strings.EqualFold("exit", text) {
            // close program on exit
            return
        } else {
            // pass command to parser
            handleCmd(text)
        }
        printPrompt()
    }
    // Print and aditional line if we encounter an EOF character
    fmt.Println()
}
