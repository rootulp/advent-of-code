package main

import "testing"

func TestPartOne(t *testing.T) {
	filename := "example.txt"
	got := PartOne(filename)
	want := 5

	if got != want {
		t.Errorf("PartOne(%v) got %v want %v", filename, got, want)
	}
}
