package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "picachu CHARMANDER MeW  ",
			expected: []string{"picachu", "charmander", "mew"},
		},
		{
			input:    "Pidgeot jigggggglY",
			expected: []string{"pidgeot", "jiggggggly"},
		},
	}


	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("fail: not expected word")
			}
		}
	}
}

