package main

import "testing"

func TestDay13(t *testing.T) {
	t.Run("GetProductOfEarliestBusAndTimeToWait", func(t *testing.T) {
		got := GetProductOfEarliestBusAndTimeToWait("input_test.txt")
		want := 295

		if got != want {
			t.Errorf("GetProductOfEarliestBusAndTimeToWait incorrect got %v want %v", got, want)
		}
	})

}
