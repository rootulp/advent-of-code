package main

import "testing"

func TestGetFirstNumberThatIsNotSumOfPair(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		got := GetFirstNumberThatIsNotSumOfPair("input_test.txt", 5)
		want := 127

		if got != want {
			t.Errorf("GetFirstNumberThatIsNotSumOfPair got %d want %d", got, want)
		}
	})
	t.Run("input", func(t *testing.T) {
		got := GetFirstNumberThatIsNotSumOfPair("input_test.txt", 25)
		want := 217430975

		if got != want {
			t.Errorf("GetFirstNumberThatIsNotSumOfPair got %d want %d", got, want)
		}
	})
}
