package main

import "testing"

func TestDay13(t *testing.T) {
	t.Run("GetProductOfEarliestBusAndTimeToWait input_test", func(t *testing.T) {
		got := GetProductOfEarliestBusAndTimeToWait("input_test.txt")
		want := 295

		if got != want {
			t.Errorf("GetProductOfEarliestBusAndTimeToWait incorrect got %v want %v", got, want)
		}
	})

	t.Run("GetProductOfEarliestBusAndTimeToWait input", func(t *testing.T) {
		got := GetProductOfEarliestBusAndTimeToWait("input.txt")
		want := 6568

		if got != want {
			t.Errorf("GetProductOfEarliestBusAndTimeToWait incorrect got %v want %v", got, want)
		}
	})

}
