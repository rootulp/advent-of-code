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

func TestPartTwo(t *testing.T) {
	got := ProductOfDepartureValues("input.txt")
	want := 362974212989

	if got != want {
		t.Errorf("Incorrect ProductOfDepartureValues got %v want %v", got, want)
	}
}
