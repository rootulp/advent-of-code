package main

import "testing"

func TestTicketScanningErrorRate(t *testing.T) {
	got := TicketScanningErrorRate("example.txt")
	want := 71

	if got != want {
		t.Errorf("Incorrect TicketScanningErrorRate got %v want %v", got, want)
	}
}
