package main

import "testing"

func TestPartOne(t *testing.T) {
	type testCase struct {
		input string
		want int
	}

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
