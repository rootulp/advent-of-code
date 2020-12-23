package main

import "testing"

func TestHelloWord(t *testing.T) {
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
}