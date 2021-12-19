package main

import "fmt"

const gridSize = 3

type cell struct {
	val rune
	z   int
	y   int
	x   int
}

type state struct {
	grid  [gridSize][gridSize][gridSize]*cell
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
func (s *state) NextCycle() {
	for _, z := range s.grid {
		for _, x := range z {
			for _, y := range x {
				y.NextCycle()
			}
		}
	}

	s.cycle += 1
}

func (s state) String() (result string) {
	for zIndex, z := range s.grid {
		result += fmt.Sprintf("\nz=%d\n", zIndex)
		for _, x := range z {
			for _, y := range x {
				result += string(y.val)
			}
			result += "\n"
		}
	}
	return result
}

func emptyGrid() (grid [gridSize][gridSize][gridSize]*cell) {
	for z := 0; z < gridSize; z++ {
		for x := 0; x < gridSize; x++ {
			for y := 0; y < gridSize; y++ {
				grid[z][x][y] = &cell{
					val: '.',
					z:   z,
					x:   x,
					y:   y,
				}
			}
		}
	}
	return grid
}

func (c cell) NextCycle() {
	return
}

// getTwoDimensionalSlice returns a two dimensional slice for the input
func getTwoDimensionalSlice(lines []string) (result [gridSize][gridSize]*cell) {
	for x, line := range lines {
		for y, r := range line {
			result[x][y] = &cell{
				val: r,
				z:   0,
				x:   x,
				y:   y,
			}
		}
	}
	return result
}
