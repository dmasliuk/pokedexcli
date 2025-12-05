package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		pokedexPrompt := "Pokedex > "
		for fmt.Print(pokedexPrompt); scanner.Scan(); fmt.Print(pokedexPrompt) {
			input := scanner.Text()
			inputSlice := cleanInput(input)
			command := strings.Join(inputSlice[:1], "")
			input = ""
			fmt.Printf("Your command was: %s\n", command)
		}
	}
	
}