package main

import "testing"

func TestPartOneExample(t *testing.T) {
	example := []int{0, 3, 6}
	got := MemoryGame(example, 10)
	want := 0

	if got != want {
		t.Errorf("MemoryGame incorrect got %v want %v", got, want)
	}
}

func TestPartOne(t *testing.T) {
	input := []int{18, 11, 9, 0, 5, 1}
	got := MemoryGame(input, 2020)
	want := 959

	if got != want {
		t.Errorf("MemoryGame incorrect got %v want %v", got, want)
	}
}
