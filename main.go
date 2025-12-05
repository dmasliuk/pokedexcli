package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
		cliCommands := map[string]cliCommand{
			"exit": {
					name:        "exit",
					description: "Exit the Pokedex",
					callback:    commandExit,
			},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		pokedexPrompt := "Pokedex > "
		for fmt.Print(pokedexPrompt); scanner.Scan(); fmt.Print(pokedexPrompt) {
			input := scanner.Text()
			inputSlice := cleanInput(input)
			command := strings.Join(inputSlice[:1], "")
			input = ""
			if command == "exit" {
				cliCommands[command].callback()
			}
		}
	}
	
}