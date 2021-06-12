package main

import (
	"testing"
)

func TestGetManhattanDistance(t *testing.T) {
	t.Run("GetManhattanDistance input_test", func(t *testing.T) {
		got := GetManhattanDistance("input_test.txt")
		want := 25

		if got != want {
			t.Errorf("GetManhattanDistance input_test incorrect got %v want %v", got, want)
		}

	})
}
