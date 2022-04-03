package main

import "testing"

type testCase struct {
	filename string
	want     int
}

var tests []testCase = []testCase{
	{"example.txt", 20899048083289},
}

func TestPartOne(t *testing.T) {
	for _, test := range tests {
		got := PartOne(test.filename)
		want := test.want

		if got != want {
			t.Errorf("PartOne(%v) got %v want %v", test.filename, got, want)
		}
	}
}
