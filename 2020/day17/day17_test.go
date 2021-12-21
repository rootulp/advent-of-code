package main

import (
	"testing"
)

func TestPartOneExample(t *testing.T) {
	got := PartOne("example.txt")
	want := 112

	if got != want {
		t.Errorf("Incorrect result for part one example got %v want %v", got, want)
	}
}

func TestPartOne(t *testing.T) {
	got := PartOne("input.txt")
	want := 382

	if got != want {
		t.Errorf("Incorrect result for part one input got %v want %v", got, want)
	}
}

func TestPartTwoExample(t *testing.T) {
	got := PartTwo("example.txt")
	want := 848

	if got != want {
		t.Errorf("Incorrect result for part two example got %v want %v", got, want)
	}
}

func TestPartExample(t *testing.T) {
	got := PartTwo("input.txt")
	want := 2552

	if got != want {
		t.Errorf("Incorrect result for part two example got %v want %v", got, want)
	}
}
