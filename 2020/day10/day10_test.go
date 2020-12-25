package main

import "testing"

func TestGetProductOfOneJoltDifferencesAndThreeJoltDifferences(t *testing.T) {
	t.Run("input_test.txt", func(t *testing.T) {
		got := GetProductOfOneJoltDifferencesAndThreeJoltDifferences("input_test.txt")
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
	t.Run("input_test2.txt", func(t *testing.T) {
		got := GetProductOfOneJoltDifferencesAndThreeJoltDifferences("input_test2.txt")
		want := 220

		if got != want {
			t.Errorf("GetProductOfOneJoltDifferencesAndThreeJoltDifferences failed got %d want %d", got, want)
		}
	})
}

func TestGetNumberOfArrangements(t *testing.T) {
	t.Run("input_test.txt", func(t *testing.T) {
		got := GetNumberOfArrangements("input_test.txt")
		want := 8

		if got != want {
			t.Errorf("GetNumberOfArrangements failed got %d want %d", got, want)
		}
	})
	t.Run("input_test2.txt", func(t *testing.T) {
		got := GetNumberOfArrangements("input_test2.txt")
		want := 19208

		if got != want {
			t.Errorf("GetNumberOfArrangements failed got %d want %d", got, want)
		}
	})
}
