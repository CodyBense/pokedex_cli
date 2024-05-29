package main

import (
    "errors"
)

type cliCommand struct {
    name string
    description string
    callback func() error
}

return map[string]cliCommand{
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
}
