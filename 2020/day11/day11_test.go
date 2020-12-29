package main

import "testing"

func TestGetCountOfOccupiedSeats(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		got := GetCountOfOccupiedSeats("input_test.txt")
		want := 37

		if got != want {
			t.Errorf("GetCountOfOccupiedSeats incorrect got %v want %v", got, want)
		}
	})
	t.Run("input", func(t *testing.T) {
		got := GetCountOfOccupiedSeats("input.txt")
		want := 2277

		if got != want {
			t.Errorf("GetCountOfOccupiedSeats incorrect got %v want %v", got, want)
		}
	})
}
