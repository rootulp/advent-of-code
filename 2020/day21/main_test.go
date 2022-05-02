package main

import "testing"

type testCase struct {
	input string
	want int
}

func TestPartOne(t *testing.T) {
	tests := []testCase{
		{"example.txt", 5},
		{"input.txt", 2659},
	}

	for _, test := range tests {
		got := PartOne(test.input)

		if got != test.want {
			t.Errorf("PartOne(%v) got %v want %v", test.input, got, test.want)
		}
	}
}
