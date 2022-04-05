package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point3d struct {
	z int
	x int
	y int
}

type point4d struct {
	z int
	w int
	x int
	y int
}

const ACTIVE = '#'

func main() {
	fmt.Println("Starting day17")

	// Part One
	partOne := PartOne("input.txt")
	fmt.Printf("Part one: %v\n", partOne)

	// Part Two
	partTwo := PartTwo("input.txt")
	fmt.Printf("Part two: %v\n", partTwo)
}

// PartOne returns the number of active cubes after simulating six cycles of
// life for the Conway Cubes given initial state in filename. Assumes three
// dimensions.
func PartOne(filename string) int {
	lines := readFile(filename)
	activeCells := getActiveCells3d(lines)
	for i := 0; i < 6; i++ {
		// fmt.Printf("i=%d, len(activeCells)=%d\n", i, len(activeCells))
		activeCells = nextCycle3d(activeCells)
	}
	return len(activeCells)
}

// PartTwo returns the number of active cubes after simulating six cycles of
// life for the Conway Cubes given initial state in filename. Assumes four
// dimensions.
func PartTwo(filename string) int {
	lines := readFile(filename)
	activeCells := getActiveCells4d(lines)
	for i := 0; i < 6; i++ {
		// fmt.Printf("i=%d, len(activeCells)=%d\n", i, len(activeCells))
		activeCells = nextCycle4d(activeCells)
	}
	return len(activeCells)
}

func nextCycle3d(current map[point3d]bool) (next map[point3d]bool) {
	next = make(map[point3d]bool)
	min, max := getBounds3d(current)
	for z := min; z <= max; z++ {
		for x := min; x <= max; x++ {
			for y := min; y <= max; y++ {
				c := point3d{z, x, y}
				activeNeighbors := getCountActiveNeighbors3d(current, c)
				if current[c] && (activeNeighbors == 2 || activeNeighbors == 3) {
					next[c] = true
				}
				if !current[c] && activeNeighbors == 3 {
					next[c] = true
				}
			}
		}
	}
	return next
}

func nextCycle4d(current map[point4d]bool) (next map[point4d]bool) {
	next = make(map[point4d]bool)
	min, max := getBounds4d(current)
	for z := min; z <= max; z++ {
		for w := min; w <= max; w++ {
			for x := min; x <= max; x++ {
				for y := min; y <= max; y++ {
					c := point4d{z, w, x, y}
					activeNeighbors := getCountActiveNeighbors4d(current, c)
					if current[c] && (activeNeighbors == 2 || activeNeighbors == 3) {
						next[c] = true
					}
					if !current[c] && activeNeighbors == 3 {
						next[c] = true
					}
				}
			}
		}
	}
	return next
}

func getBounds3d(current map[point3d]bool) (min int, max int) {
	for c := range current {
		min = getMin(min, c.z, c.x, c.y)
		max = getMax(max, c.z, c.x, c.y)
	}
	return min - 1, max + 1
}

func getBounds4d(current map[point4d]bool) (min int, max int) {
	for c := range current {
		min = getMin(min, c.z, c.w, c.x, c.y)
		max = getMax(max, c.z, c.w, c.x, c.y)
	}
	return min - 1, max + 1
}

func getCountActiveNeighbors3d(current map[point3d]bool, c point3d) (count int) {
	neighbors := getNeighbors3d(current, c)
	for _, neighbor := range neighbors {
		if current[neighbor] {
			count += 1
		}
	}
	return count
}

func getCountActiveNeighbors4d(current map[point4d]bool, c point4d) (count int) {
	neighbors := getNeighbors4d(current, c)
	for _, neighbor := range neighbors {
		if current[neighbor] {
			count += 1
		}
	}
	return count
}

func getNeighbors3d(current map[point3d]bool, c point3d) (neighbors []point3d) {
	for dz := -1; dz <= 1; dz++ {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dz == 0 && dx == 0 && dy == 0 {
					continue
				}
				neighbor := point3d{
					z: c.z + dz,
					x: c.x + dx,
					y: c.y + dy,
				}
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}

func getNeighbors4d(current map[point4d]bool, c point4d) (neighbors []point4d) {
	for dz := -1; dz <= 1; dz++ {
		for dw := -1; dw <= 1; dw++ {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dz == 0 && dw == 0 && dx == 0 && dy == 0 {
						continue
					}
					neighbor := point4d{
						z: c.z + dz,
						w: c.w + dw,
						x: c.x + dx,
						y: c.y + dy,
					}
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}
	return neighbors
}

func getActiveCells3d(lines []string) (activeCells map[point3d]bool) {
	activeCells = make(map[point3d]bool)
	for x, line := range lines {
		for y, r := range line {
			if r == ACTIVE {
				activeCells[point3d{
					z: 0, // every element in input is at z=0
					x: x,
					y: y,
				}] = true
			}
		}
	}
	return activeCells
}

func getActiveCells4d(lines []string) (activeCells map[point4d]bool) {
	activeCells = make(map[point4d]bool)
	for x, line := range lines {
		for y, r := range line {
			if r == ACTIVE {
				activeCells[point4d{
					z: 0, // every element in input is at z=0
					w: 0, // every element in input is at w=0
					x: x,
					y: y,
				}] = true
			}
		}
	}
	return activeCells
}

func readFile(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func getMin(elements ...int) (min int) {
	for _, element := range elements {
		if element < min {
			min = element
		}
	}
	return min
}

func getMax(elements ...int) (result int) {
	for _, element := range elements {
		if element > result {
			result = element
		}
	}
	return result
}
