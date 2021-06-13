package main

import (
	"testing"
)

func TestGetManhattanDistance(t *testing.T) {
	t.Run("GetManhattanDistancePartOne input_test", func(t *testing.T) {
		got := GetManhattanDistancePartOne("input_test.txt")
		want := 25

		if got != want {
			t.Errorf("GetManhattanDistancePartOne input_test incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetManhattanDistancePartOne input", func(t *testing.T) {
		got := GetManhattanDistancePartOne("input.txt")
		want := 1294

		if got != want {
			t.Errorf("GetManhattanDistancePartOne input_test incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetManhattanDistancePartTwo input_test", func(t *testing.T) {
		got := GetManhattanDistancePartTwo("input_test.txt")
		want := 286

		if got != want {
			t.Errorf("GetManhattanDistancePartTwo input_test incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetManhattanDistancePartTwo input_test", func(t *testing.T) {
		got := GetManhattanDistancePartTwo("input_test.txt")
		want := 20592

		if got != want {
			t.Errorf("GetManhattanDistancePartTwo input_test incorrect got %v want %v", got, want)
		}
	})
}
