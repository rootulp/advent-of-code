package main

import "testing"

type testCase struct {
	input string
	want  int
}

var partOneTests []testCase = []testCase{
	{"example.txt", 2},
	{"input.txt", 142},
}

func TestPartOne(t *testing.T) {
	for _, test := range partOneTests {
		got := PartOne(test.input)

		if got != test.want {
			t.Errorf("PartOne(%v) got %v want %v", test.input, got, test.want)
		}
	}
}

var partTwoTests []testCase = []testCase{
	{"example2.txt", 12},
	{"input.txt", 294},
}

func TestPartTwo(t *testing.T) {
	for _, test := range partTwoTests {
		got := PartTwo(test.input)

		if got != test.want {
			t.Errorf("PartTwo(%v) got %v want %v", test.input, got, test.want)
		}
	}
}
