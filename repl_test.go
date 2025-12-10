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
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "H   e     L       l  O",
			expected: []string{"h", "e", "l", "l", "o"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// 1) First check: length of slices
		if len(actual) != len(c.expected) {
        	t.Errorf("for input %q got %v, want %v", c.input, actual, c.expected)
        	continue
		}

		// 2) Second check: each word
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("for input %q word %d = %q, want %q",
                    c.input, i, actual[i], c.expected[i])
			}
		}
	}
}