package main

import "testing"

type testCase struct {
	filename string
	want     int
}

func TestPartOne(t *testing.T) {
	var tests []testCase = []testCase{
		{"example.txt", 20899048083289},
		{"input.txt", 18262194216271},
	}

	for _, test := range tests {
		got := PartOne(test.filename)
		want := test.want

		if got != want {
			t.Errorf("PartOne(%v) got %v want %v", test.filename, got, want)
		}
	}
}

func TestPartTwo(t *testing.T) {
	var tests []testCase = []testCase{
		{"example.txt", 273},
		{"input.txt", 2023},
	}

	for _, test := range tests {
		got := PartTwo(test.filename)
		want := test.want

		if got != want {
			t.Errorf("PartTwo(%v) got %v want %v", test.filename, got, want)
		}
	}
}
