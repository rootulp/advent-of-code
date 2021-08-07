package main

import "testing"

func TestGetFirstNumberThatIsNotSumOfPair(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := GetFirstNumberThatIsNotSumOfPair("example.txt", 5)
		want := 127

		if got != want {
			t.Errorf("GetFirstNumberThatIsNotSumOfPair got %d want %d", got, want)
		}
	})
	t.Run("input", func(t *testing.T) {
		got := GetFirstNumberThatIsNotSumOfPair("input.txt", 25)
		want := 217430975

		if got != want {
			t.Errorf("GetFirstNumberThatIsNotSumOfPair got %d want %d", got, want)
		}
	})
}

func TestGetSumOfSmallestAndLargestInContiguousRange(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := GetSumOfSmallestAndLargestInContiguousRange("example.txt", 5)
		want := 62

		if got != want {
			t.Errorf("GetSumOfSmallestAndLargestInContiguousRange got %d want %d", got, want)
		}
	})
	t.Run("input", func(t *testing.T) {
		got := GetSumOfSmallestAndLargestInContiguousRange("input.txt", 25)
		want := 28509180

		if got != want {
			t.Errorf("GetSumOfSmallestAndLargestInContiguousRange got %d want %d", got, want)
		}
	})
}
