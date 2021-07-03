package main

import "testing"

func TestPartOneExample(t *testing.T) {
	got := PartOne("example.txt")
	want := 165
	if got != want {
		t.Errorf("PartOne incorrect got %v want %v", got, want)
	}
}

func TestPartOne(t *testing.T) {
	got := PartOne("input.txt")
	want := 9967721333886
	if got != want {
		t.Errorf("PartOne incorrect got %v want %v", got, want)
	}
}

func TestPartTwoExample(t *testing.T) {
	got := PartTwo("example2.txt")
	want := 208
	if got != want {
		t.Errorf("PartTwo incorrect got %v want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got := PartTwo("input.txt")
	want := 4355897790573
	if got != want {
		t.Errorf("PartTwo incorrect got %v want %v", got, want)
	}
}
