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
