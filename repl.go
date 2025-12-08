package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct{
	name string
	description string
	callback func() error
}

//function to split the user's input into "words" based on whitespace
func cleanInput(text string) []string {
	text = strings.ToLower(text)
	names := strings.Fields(text)
	return names
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
			"exit": {
					name:        "exit",
					description: "Exit the Pokedex",
					callback:    commandExit,
			},
			"help": {
				name: "help",
				description: "Displays a help message",
				callback: commandHelp,
			},
		}
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for k, c := range commands {
		fmt.Printf("%s: %s\n", k, c.description)
	}
	return nil
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}