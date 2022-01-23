package main

import "testing"

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
		got := Evaluate(test.expression)
		want := test.want
		if got != want {
			t.Errorf("Evaluate got %v want %v", got, want)
		}
	}
}
