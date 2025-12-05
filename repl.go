package main

import "strings"

//function to split the user's input into "words" based on whitespace
func cleanInput(text string) []string {
	text = strings.ToLower(text)
	names := strings.Fields(text)
	return names
}