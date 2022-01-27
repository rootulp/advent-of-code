package main

import "testing"

var operatorPrecedence map[string]int = map[string]int{
	"+": 1,
	"*": 1,
}

func TestPartOne(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{"example.txt", 26335},
		{"input.txt", 5019432542701},
	}

	for _, test := range tests {
		got := PartOne(test.filename)
		if got != test.want {
			t.Errorf("PartOne(%v) got %v want %v", test.filename, got, test.want)
		}
	}
}

func TestPartTwo(t *testing.T) {
	type test struct {
		filename string
		want     int
	}

	tests := []test{
		{"example.txt", 693891},
		{"input.txt", 70518821989947},
	}

	for _, test := range tests {
		got := PartTwo(test.filename)
		if got != test.want {
			t.Errorf("PartTwo(%v) got %v want %v", test.filename, got, test.want)
		}
	}
}

func TestEvaluateExpression(t *testing.T) {
	type test struct {
		expression string
		want       int
	}

	tests := []test{
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, test := range tests {
		got := Evaluate(test.expression, operatorPrecedence)
		want := test.want
		if got != want {
			t.Errorf("Evaluate(%v) got %v want %v", test.expression, got, want)
		}
	}
}

func TestReversePolishNotation(t *testing.T) {
	type test struct {
		expression string
		want       string
	}

	tests := []test{
		{"3 + 4", "3 4 +"},
		{"7 + 8 * 6", "7 8 + 6 *"},
		{"2 * 3 + (4 * 5)", "2 3 * 4 5 * +"},
	}

	for _, test := range tests {
		got := ReversePolishNotation(test.expression, operatorPrecedence)
		want := test.want
		if got != want {
			t.Errorf("ReversePolishNotation(%v) got %v want %v", test.expression, got, want)
		}
	}
}

func TestEvaluateReversePolishNotation(t *testing.T) {
	type test struct {
		expression string
		want       int
	}

	tests := []test{
		{"3 4 +", 7},
		{"7 8 + 6 *", 90},
		{"2 3 * 4 5 * +", 26},
	}

	for _, test := range tests {
		got := EvaluateReversePolishNotation(test.expression)
		want := test.want
		if got != want {
			t.Errorf("EvaluateReversePolishNotation(%v) got %v want %v", test.expression, got, want)
		}
	}
}
