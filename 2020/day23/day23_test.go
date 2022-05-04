package main

import "testing"

func TestPartOne(t *testing.T) {
	type testCase struct {
		input string
		numMoves int
		want string
	}

	tests := []testCase{
		// Example
		{"389125467", 10, "92658374"},
		{"389125467", 100, "67384529"},
		// Input
		{"685974213", 100, "82635947"},
	}

	for _, test := range tests {
		got := PartOne(test.input, test.numMoves)

		if got != test.want {
			t.Fatalf("PartOne(%v, %v) got %v want %v", test.input, test.numMoves, got, test.want)
		}
	}

}
