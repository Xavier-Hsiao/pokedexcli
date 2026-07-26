package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	// create the test suite
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
	}

	// loop over the cases and run the tests
	for _, c := range cases {
		got := cleanInput(c.input)
		if len(got) != len(c.expected) {
			t.Errorf("Lengths don't match: '%v' vs. '%v", got, c.expected)
			continue
		}
		for i := range got {
			word := got[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("After cleanup: %v, but expected %v", got, c.expected)
			}
		}
	}
}
