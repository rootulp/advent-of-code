package main

import "testing"

type testCase struct {
	input string
	want int
}

func TestPartOne(t *testing.T) {
	tests := []testCase{
		{"example.txt", 306},
		{"input.txt", 33403},
	}

	for _, test := range tests {
		got := PartOne(test.input)

		if got != test.want {
			t.Errorf("PartOne(%v) got %v want %v", test.input, got, test.want)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []testCase {
		{"example.txt", 291},
		// takes a while to run
		// {"input.txt", 29177},
	}

	for _, test := range tests {
		got := PartTwo(test.input)

		if got != test.want {
			t.Errorf("PartTwo(%v) got %v want %v", test.input, got, test.want)
		}
	}
}
