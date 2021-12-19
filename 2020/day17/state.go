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
	grid  [gridSize][gridSize][gridSize]cell
	cycle int
}

func NewState(lines []string) (s state) {
	grid := emptyGrid()
	zeroIndexSlice := getTwoDimensionalSlice(lines)
	// s.grid is a slice of z = -1, 0, 1.
	grid[1] = zeroIndexSlice

	return state{
		grid:  grid,
		cycle: 0,
	}
}

func (source state) Clone() (cloned state) {
	grid := emptyGrid()

	for z, plane := range source.grid {
		for x, line := range plane {
			for y, c := range line {
				grid[z][x][y] = cell{
					val: c.val,
					z:   z,
					x:   x,
					y:   y,
				}
			}
		}
	}

	return state{
		grid:  grid,
		cycle: source.cycle,
	}
}

// NextCycle advances state by one cycle
func (current state) NextCycle() (next state) {
	next = current.Clone()

	for z, plane := range current.grid {
		for x, line := range plane {
			for y, c := range line {
				next.grid[z][x][y] = cell{
					val: c.NextCycle(current),
					z:   z,
					x:   x,
					y:   y,
				}
			}
		}
	}

	next.cycle = current.cycle + 1
	return next
}

func (s state) String() (result string) {
	for zIndex, z := range s.grid {
		// Z is 0 based and does not reflect negative values
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

func emptyGrid() (grid [gridSize][gridSize][gridSize]cell) {
	for z := 0; z < gridSize; z++ {
		for x := 0; x < gridSize; x++ {
			for y := 0; y < gridSize; y++ {
				grid[z][x][y] = cell{
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

func (c cell) NextCycle(s state) rune {
	if c.isActiveNextCycle(s) {
		return '#'
	} else {
		return '.'
	}
}

func (c cell) isActiveNextCycle(s state) (isActive bool) {
	activeNeighbors := s.getActiveNeighbors(c)
	if c.isActive() {
		return activeNeighbors == 2 || activeNeighbors == 3
	} else {
		return activeNeighbors == 3
	}
}

func (c cell) isActive() bool {
	return c.val == '#'
}

func (s state) getActiveNeighbors(c cell) (activeNeighbors int) {
	neighbors := s.getNeighbors(c)
	for _, neighbor := range neighbors {
		if neighbor.isActive() {
			activeNeighbors += 1
		}
	}
	return activeNeighbors
}

func (s state) getNeighbors(c cell) (neighbors []cell) {
	for dz := -1; dz < 1; dz++ {
		for dx := -1; dx < 1; dx++ {
			for dy := -1; dy < 1; dy++ {
				targetZ := c.z + dz
				targetX := c.x + dx
				targetY := c.z + dy
				if targetZ >= 0 && targetZ < len(s.grid) &&
					targetX >= 0 && targetX < len(s.grid[0]) &&
					targetY >= 0 && targetY < len(s.grid[0][0]) {
					neighbors = append(neighbors, s.grid[targetZ][targetX][targetY])
				}
			}
		}
	}
	return neighbors
}

// getTwoDimensionalSlice returns a two dimensional slice for the input
func getTwoDimensionalSlice(lines []string) (result [gridSize][gridSize]cell) {
	for x, line := range lines {
		for y, r := range line {
			result[x][y] = cell{
				val: r,
				z:   0,
				x:   x,
				y:   y,
			}
		}
	}
	return result
}
