package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cliCommands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		pokedexPrompt := "Pokedex > "
		for fmt.Print(pokedexPrompt); scanner.Scan(); fmt.Print(pokedexPrompt) {
			input := scanner.Text()
			inputSlice := cleanInput(input)
			command := strings.Join(inputSlice[:1], "")
			input = ""
			switch command {
				case "exit": {
					cliCommands[command].callback()
				}
				case "help": {
					cliCommands[command].callback()
				}
				default: {
					fmt.Println(pokedexPrompt + "Unknown command")
				}
			}
		}
	}
	
}