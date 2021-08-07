package main

import "testing"

func TestGetProductOfOneJoltDifferencesAndThreeJoltDifferences(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		got := GetProductOfOneJoltDifferencesAndThreeJoltDifferences("example.txt")
		want := 35

		if got != want {
			t.Errorf("GetProductOfOneJoltDifferencesAndThreeJoltDifferences failed got %d want %d", got, want)
		}
	})
	t.Run("input.txt", func(t *testing.T) {
		got := GetProductOfOneJoltDifferencesAndThreeJoltDifferences("input.txt")
		want := 2400

		if got != want {
			t.Errorf("GetProductOfOneJoltDifferencesAndThreeJoltDifferences failed got %d want %d", got, want)
		}
	})
	t.Run("example2.txt", func(t *testing.T) {
		got := GetProductOfOneJoltDifferencesAndThreeJoltDifferences("example2.txt")
		want := 220

		if got != want {
			t.Errorf("GetProductOfOneJoltDifferencesAndThreeJoltDifferences failed got %d want %d", got, want)
		}
	})
}

func TestGetNumberOfArrangements(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		got := GetNumberOfArrangements("example.txt")
		want := 8

		if got != want {
			t.Errorf("GetNumberOfArrangements failed got %d want %d", got, want)
		}
	})
	t.Run("input.txt", func(t *testing.T) {
		got := GetNumberOfArrangements("input.txt")
		want := 338510590509056

		if got != want {
			t.Errorf("GetNumberOfArrangements failed got %d want %d", got, want)
		}
	})
	t.Run("example2.txt", func(t *testing.T) {
		got := GetNumberOfArrangements("example2.txt")
		want := 19208

		if got != want {
			t.Errorf("GetNumberOfArrangements failed got %d want %d", got, want)
		}
	})
}
