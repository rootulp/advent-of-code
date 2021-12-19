package main

import (
	"testing"
)

func TestPartOneExampleInitialState(t *testing.T) {
	lines := readFile("example.txt")
	s := NewState()
	s.Initialize(lines)
	got := s.String()
	want := `
z=0
...
...
...

z=1
.#.
..#
###

z=2
...
...
...
`
	if got != want {
		t.Errorf("Incorrect example initial state got %v want %v", s, want)
	}
}

// func TestPartOneExampleAfterOneCycle(t *testing.T) {
// 	lines := readFile("example.txt")
// 	got := newState(lines)
// 	want := state{[gridSize][gridSize][gridSize]rune{
// 		{ // z = -1
// 			{'#', '.', '.'},
// 			{'.', '.', '#'},
// 			{'.', '#', '.'},
// 		},
// 		{ // z = 0
// 			{'#', '.', '#'},
// 			{'.', '#', '#'},
// 			{'.', '#', '.'},
// 		},
// 		{ // z = 1
// 			{'#', '.', '.'},
// 			{'.', '.', '#'},
// 			{'.', '#', '.'},
// 		},
// 	}}
// 	if got != want {
// 		t.Errorf("Incorrect example initial state got %v want %v", got, want)
// 	}
// }
