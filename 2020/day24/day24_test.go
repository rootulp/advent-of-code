package main

import "testing"

func TestPartOne(t *testing.T) {
	type testCase struct {
		input string
		want int
	}
	tests := []testCase {
		{"example.txt", 10},
		{"input.txt", 375},
	}

	for _, test := range tests {
		got := PartOne(test.input)

		if got != test.want {
			t.Errorf("PartOne(%s) got %d want %d", test.input, got, test.want)
		}
	}
}
