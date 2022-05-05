package main

import "testing"


func TestPartOne(t *testing.T) {
	type testCase struct {
		input string
		numMoves int
		want string
	}

	tests := []testCase{
		{EXAMPLE_INPUT, 10, "92658374"},
		{EXAMPLE_INPUT, 100, "67384529"},
		{INPUT, 100, "82635947"},
	}

	for _, test := range tests {
		got := PartOne(test.input, test.numMoves)

		if got != test.want {
			t.Fatalf("PartOne(%v, %v) got %v want %v", test.input, test.numMoves, got, test.want)
		}
	}
}

// func TestPartTwo(t *testing.T) {
// 	type testCase struct {
// 		input string
// 		numMoves int
// 		want int
// 	}

// 	tests := []testCase{
// 		{EXAMPLE_INPUT, 10_000_000, 92658374},
// 		// {INPUT, 10_000_000, 0}, // TODO 0 -> real result
// 	}

// 	for _, test := range tests {
// 		got := PartTwo(test.input, test.numMoves)

// 		if got != test.want {
// 			t.Fatalf("PartTwo(%v, %v) got %v want %v", test.input, test.numMoves, got, test.want)
// 		}
// 	}
// }
