package main

import "testing"

func TestPartOneExample(t *testing.T) {
	got := TicketScanningErrorRate("example.txt")
	want := 71

	if got != want {
		t.Errorf("Incorrect TicketScanningErrorRate got %v want %v", got, want)
	}
}

func TestPartOne(t *testing.T) {
	got := TicketScanningErrorRate("input.txt")
	want := 23122

	if got != want {
		t.Errorf("Incorrect TicketScanningErrorRate got %v want %v", got, want)
	}
}
