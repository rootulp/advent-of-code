package main

import (
	"testing"
)

func TestGetManhattanDistance(t *testing.T) {
	t.Run("GetManhattanDistancePartOne example", func(t *testing.T) {
		got := GetManhattanDistancePartOne("example.txt")
		want := 25

		if got != want {
			t.Errorf("GetManhattanDistancePartOne example incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetManhattanDistancePartOne input", func(t *testing.T) {
		got := GetManhattanDistancePartOne("input.txt")
		want := 1294

		if got != want {
			t.Errorf("GetManhattanDistancePartOne example incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetManhattanDistancePartTwo example", func(t *testing.T) {
		got := GetManhattanDistancePartTwo("example.txt")
		want := 286

		if got != want {
			t.Errorf("GetManhattanDistancePartTwo example incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetManhattanDistancePartTwo example", func(t *testing.T) {
		got := GetManhattanDistancePartTwo("input.txt")
		want := 20592

		if got != want {
			t.Errorf("GetManhattanDistancePartTwo example incorrect got %v want %v", got, want)
		}
	})
}
