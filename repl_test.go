package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Lorem ipsum dolor sit amet  ",
			expected: []string{"lorem", "ipsum", "dolor", "sit", "amet"},
		},
		{
			input:    "  This CRAZY idea  ",
			expected: []string{"this", "crazy", "idea"},
		},
		{
			input:    "  more CrAzY text  ",
			expected: []string{"more", "crazy", "text"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual length is %d; expected %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v; expected %v", c.input, actual, c.expected)
			}
		}
	}
}
