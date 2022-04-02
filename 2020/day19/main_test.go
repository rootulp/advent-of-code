package main

import "testing"

type testCase struct {
	input string
	want  int
}

var tests []testCase = []testCase{
	{"example.txt", 2},
	{"input.txt", 0},
}

func TestPartOne(t *testing.T) {
	for _, testCase := range tests {
		got := PartOne(testCase.input)

		if got != testCase.want {
			t.Errorf("PartOne(%v) got %v want %v", testCase.input, got, testCase.want)
		}
	}
}
