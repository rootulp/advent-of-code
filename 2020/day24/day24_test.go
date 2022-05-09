package main

import "testing"

func TestPartOne(t *testing.T) {
	type testCase struct {
		input string
		want  int
	}
	tests := []testCase{
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

func TestPartTwo(t *testing.T) {
	type testCase struct {
		input   string
		numDays int
		want    int
	}
	tests := []testCase{
		{"example.txt", 1, 15},
		{"example.txt", 2, 12},
		{"example.txt", 3, 25},
		{"example.txt", 10, 37},
		{"example.txt", 20, 132},
		{"example.txt", 30, 259},
		{"example.txt", 40, 406},
		{"example.txt", 50, 566},
		{"example.txt", 60, 788},
		{"example.txt", 70, 1106},
		{"example.txt", 100, 2208},
		{"input.txt", 100, 3937},
	}

	for _, test := range tests {
		got := PartTwo(test.input, test.numDays)

		if got != test.want {
			t.Errorf("PartTwo(%s, %d) got %d want %d", test.input, test.numDays, got, test.want)
		}
	}

}
