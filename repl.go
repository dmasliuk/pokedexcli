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

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}