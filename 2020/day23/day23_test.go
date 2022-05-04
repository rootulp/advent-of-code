package main

import "testing"

func TestPartOne(t *testing.T) {
	type testCase struct {
		input string
		numMoves int
		want string
	}

	tests := []testCase{
		{"389125467", 10, "92658374"},
	}

	for _, test := range tests {
		got := PartOne(test.input, test.numMoves)

		if got != test.want {
			t.Fatalf("PartOne(%v, %v) got %v want %v", test.input, test.numMoves, got, test.want)
		}
	}

}
