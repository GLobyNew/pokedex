package main 

import (
	"testing"
)


func TestCleanInput(t *testing.T) {

cases := []struct {
	input string
	expected []string
}{
	{
	input: "   hello world    ",
	expected: []string{"hello", "world"},
	},
	{
	input: "shit BOG DOGshiet ",
	expected: []string{"shit", "bog", "dogshiet"},
	},
	{
	input: "icarus flame is in hand   ",
	expected: []string{"icarus", "flame", "is", "in", "hand"},
	},
}

for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("Error: actual and expected lengths are not the same")
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("Error: actual word is not an expected word")
		}
	}
}

}
