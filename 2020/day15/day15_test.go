package main

import "testing"

var exampleTests = []struct {
	startingNumbers []int
	turns           int
	want            int
}{
	{[]int{0, 3, 6}, 10, 0},
	{[]int{1, 3, 2}, 2020, 1},
	{[]int{2, 1, 3}, 2020, 10},
	{[]int{1, 2, 3}, 2020, 27},
	{[]int{2, 3, 1}, 2020, 78},
	{[]int{3, 2, 1}, 2020, 438},
	{[]int{3, 1, 2}, 2020, 1836},
}

func TestPartOneExample(t *testing.T) {
	for _, tt := range exampleTests {
		t.Run("Example", func(t *testing.T) {
			got := MemoryGame(tt.startingNumbers, tt.turns)
			want := tt.want

			if got != want {
				t.Errorf("MemoryGame incorrect got %v want %v", got, want)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	input := []int{18, 11, 9, 0, 5, 1}
	got := MemoryGame(input, 2020)
	want := 959

	if got != want {
		t.Errorf("MemoryGame incorrect got %v want %v", got, want)
	}
}
