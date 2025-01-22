package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "blaziken lucario", expected: []string{"blaziken", "lucario"}},
		{input: "mewtwo mew", expected: []string{"mewtwo", "mew"}},
		{input: "ditto ,eve", expected: []string{"ditto", ",eve"}},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len of input does not match with expected")
		}

		for i, _ := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word does not match with expected word")
			}
		}
	}
}
