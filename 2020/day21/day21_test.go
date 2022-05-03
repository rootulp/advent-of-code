package main

import "testing"


func TestPartOne(t *testing.T) {
	type testCase struct {
		input string
		want int
	}

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


func TestPartTwo(t *testing.T) {
	type testCase struct {
		input string
		want string
	}

	tests := []testCase{
		{"example.txt", "mxmxvkd,sqjhc,fvjkl"},
		{"input.txt", "rcqb,cltx,nrl,qjvvcvz,tsqpn,xhnk,tfqsb,zqzmzl"},
	}

	for _, test := range tests {
		got := PartTwo(test.input)

		if got != test.want {
			t.Errorf("PartTwo(%v) got %v want %v", test.input, got, test.want)
		}
	}
}
