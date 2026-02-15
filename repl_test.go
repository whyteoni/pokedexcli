package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  ",
			expected: []string{},
		},
		{
			input: "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input: "   multiple   spaces   ",
			expected: []string{"multiple", "spaces"},
		},
		{
			input: "single",
			expected: []string{"single"},
		},
		{
			input: "\t\n\r",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("wrong number of elements: actual '%v', expected '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("unexpected value: cleanInput(%v) yielded '%v', expected '%v'", c.input, actual, c.expected)
			}
		}
	}
}
