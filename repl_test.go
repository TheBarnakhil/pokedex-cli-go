package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello    world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " foo    Bar ",
			expected: []string{"foo", "bar"},
		},
		{
			input:    " Star   Platinum   Za   Warudo ",
			expected: []string{"star", "platinum", "za", "warudo"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("The length do not match!\n")
			t.Errorf("Expected: %d vs Actual %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("The words do not match!\n")
				t.Errorf("Expected: %s vs Actual %s", expectedWord, word)
			}
		}
	}

}
