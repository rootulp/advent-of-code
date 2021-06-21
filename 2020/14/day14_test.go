package main

import "testing"

func TestSumOfBitmaskedValues(t *testing.T) {
	got := SumOfBitmaskedValues("example.txt")
	want := 165
	if got != want {
		t.Errorf("TestSumOfBitmaskedValues incorrect got %v want %v", got, want)
	}
}
