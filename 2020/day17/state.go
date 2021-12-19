package main

import "fmt"

const gridSize = 3

type state struct {
	grid  [gridSize][gridSize][gridSize]rune
	cycle int
}

func NewState() (s *state) {
	return &state{
		grid:  emptyGrid(),
		cycle: 0,
	}
}

func (s *state) Initialize(lines []string) {
	zeroIndexSlice := getTwoDimensionalSlice(lines)
	// s.grid is a slice of z = -1, 0, 1.
	s.grid[1] = zeroIndexSlice
}

// NextCycle advances state by one cycle
func (s state) NextCycle() (nextState state) {
	return state{
		grid:  emptyGrid(),
		cycle: s.cycle + 1,
	}
}

func (s state) String() (result string) {
	for zIndex, z := range s.grid {
		result += fmt.Sprintf("\nz=%d\n", zIndex)
		for _, x := range z {
			for _, y := range x {
				result += string(y)
			}
			result += "\n"
		}
	}
	return result
}

func emptyGrid() [gridSize][gridSize][gridSize]rune {
	return [gridSize][gridSize][gridSize]rune{
		{ // z = -1
			{'.', '.', '.'},
			{'.', '.', '.'},
			{'.', '.', '.'},
		},
		{ // z = 0
			{'.', '.', '.'},
			{'.', '.', '.'},
			{'.', '.', '.'},
		},
		{ // z = 1
			{'.', '.', '.'},
			{'.', '.', '.'},
			{'.', '.', '.'},
		},
	}
}
