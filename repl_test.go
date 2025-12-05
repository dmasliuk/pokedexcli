package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " pikachu Bulbasaur Squirtle charmander ",
			expected: []string {"pikachu", "bulbasaur", "squirtle", "charmander"},
		},
		{
			input: " some pokemon In Pokedex ",
			expected: []string {"some", "pokemon", "in", "pokedex"},
		},
		{
			input: " Venusaur butterfree Pidgey",
			expected: []string {"venusaur", "butterfree", "pidgey"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Incorrect input: %v, expected: %v", actual, c.expected)
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Incorrect input: %v, expected: %v", actual, c.expected)
				return
			}
		}
	}
}