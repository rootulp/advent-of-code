package main

import (
	"testing"
)

func TestGetCountOfOccupiedSeats(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		got := GetCountOfOccupiedSeatsPartOne("input_test.txt")
		want := 37

		if got != want {
			t.Errorf("GetCountOfOccupiedSeats incorrect got %v want %v", got, want)
		}
	})
	t.Run("input", func(t *testing.T) {
		got := GetCountOfOccupiedSeatsPartOne("input.txt")
		want := 2277

		if got != want {
			t.Errorf("GetCountOfOccupiedSeats incorrect got %v want %v", got, want)
		}
	})
	t.Run("input_test after one round", func(t *testing.T) {
		lines := readFile("input_test.txt")
		grid := getGrid(lines)
		got := toString(tickPartOne(grid))

		lines_after_one_round := readFile("input_test_after_one_round.txt")
		want := toString(getGrid(lines_after_one_round))

		if got != want {
			t.Errorf("Grid after one tick incorrect got %v want %v", got, want)
		}
	})
}
