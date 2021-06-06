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
	t.Run("input_test part one after five rounds", func(t *testing.T) {
		lines := readFile("input_test.txt")
		grid := getGrid(lines)

		// Advance five rounds
		grid = tickPartOne(grid)
		grid = tickPartOne(grid)
		grid = tickPartOne(grid)
		grid = tickPartOne(grid)
		grid = tickPartOne(grid)
		got := toString(grid)

		lines_after_five_rounds := readFile("input_test_part_one_after_five_rounds.txt")
		want := toString(getGrid(lines_after_five_rounds))

		if got != want {
			t.Errorf("Grid after five rounds incorrect got %v want %v", got, want)
		}
	})
	t.Run("input part one", func(t *testing.T) {
		got := GetCountOfOccupiedSeatsPartOne("input.txt")
		want := 2277

		if got != want {
			t.Errorf("GetCountOfOccupiedSeats incorrect got %v want %v", got, want)
		}
	})
}
