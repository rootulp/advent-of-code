package main

import (
	"testing"
)

func TestGetCountOfOccupiedSeats(t *testing.T) {
	t.Run("input_test part one count of occupied seats", func(t *testing.T) {
		got := GetCountOfOccupiedSeatsPartOne("input_test.txt")
		want := 37

		if got != want {
			t.Errorf("GetCountOfOccupiedSeats part one incorrect got %v want %v", got, want)
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
			t.Errorf("Input test part one after five rounds incorrect got %v want %v", got, want)
		}
	})
	t.Run("input_test part two count of occupied seats", func(t *testing.T) {
		got := GetCountOfOccupiedSeatsPartTwo("input_test.txt")
		want := 26

		if got != want {
			t.Errorf("GetCountOfOccupiedSeats part two incorrect got %v want %v", got, want)
		}
	})
	t.Run("input_test part two after five rounds", func(t *testing.T) {
		lines := readFile("input_test.txt")
		grid := getGrid(lines)

		// Advance five rounds
		grid = tickPartTwo(grid)
		grid = tickPartTwo(grid)
		grid = tickPartTwo(grid)
		grid = tickPartTwo(grid)
		grid = tickPartTwo(grid)
		got := toString(grid)

		lines_after_five_rounds := readFile("input_test_part_two_after_five_rounds.txt")
		want := toString(getGrid(lines_after_five_rounds))

		if got != want {
			t.Errorf("Input test part two after five rounds incorrect got \n%v want \n%v", got, want)
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
