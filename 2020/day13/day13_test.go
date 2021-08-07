package main

import "testing"

func TestDay13(t *testing.T) {
	t.Run("GetProductOfEarliestBusAndTimeToWait example", func(t *testing.T) {
		got := GetProductOfEarliestBusAndTimeToWait("example.txt")
		want := 295

		if got != want {
			t.Errorf("GetProductOfEarliestBusAndTimeToWait incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetProductOfEarliestBusAndTimeToWait input", func(t *testing.T) {
		got := GetProductOfEarliestBusAndTimeToWait("input.txt")
		want := 6568

		if got != want {
			t.Errorf("GetProductOfEarliestBusAndTimeToWait incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetEarliestBusWithSubsequentDepartures example", func(t *testing.T) {
		got := GetEarliestBusWithSubsequentDepartures("example.txt")
		want := 1068781

		if got != want {
			t.Errorf("GetProductOfEarliestBusAndTimeToWait incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetEarliestBusWithSubsequentDepartures example", func(t *testing.T) {
		got := GetEarliestBusWithSubsequentDepartures("input.txt")
		want := 554865447501099

		if got != want {
			t.Errorf("GetProductOfEarliestBusAndTimeToWait incorrect got %v want %v", got, want)
		}
	})
}
