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
	partTwo := PartTwo("example.txt")
	fmt.Printf("Part two: %v\n", partTwo)
}

// PartOne returns the number of active cubes after simulating six cycles of
// life for the Conway Cubes given initial state in filename. Assumes three
// dimensions.
func PartOne(filename string) int {
	lines := readFile(filename)
	activeCells := getActiveCells3d(lines)
	for i := 0; i < 6; i++ {
		fmt.Printf("i=%d, len(activeCells)=%d\n", i, len(activeCells))
		activeCells = nextCycle3d(activeCells)
	}
	return len(activeCells)
}

// PartTwo returns the number of active cubes after simulating six cycles of
// life for the Conway Cubes given initial state in filename. Assumes four
// dimensions.
func PartTwo(filename string) int {
	lines := readFile(filename)
	activeCells := getActiveCells3d(lines)
	for i := 0; i < 6; i++ {
		fmt.Printf("i=%d, len(activeCells)=%d\n", i, len(activeCells))
		activeCells = nextCycle3d(activeCells)
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

// func nextCycle4d(current map[point4d]bool) (next map[point4d]bool) {
// 	next = make(map[point4d]bool)
// 	min, max := getBounds(current)
// 	for z := min; z <= max; z++ {
// 		for x := min; x <= max; x++ {
// 			for y := min; y <= max; y++ {
// 				c := point3d{z, x, y}
// 				activeNeighbors := getCountActiveNeighbors(current, c)
// 				if current[c] && (activeNeighbors == 2 || activeNeighbors == 3) {
// 					next[c] = true
// 				}
// 				if !current[c] && activeNeighbors == 3 {
// 					next[c] = true
// 				}
// 			}
// 		}
// 	}
// 	return next
// }

func getBounds3d(current map[point3d]bool) (min int, max int) {
	for c := range current {
		min = minimum(min, minimum(c.z, minimum(c.x, c.y)))
		max = maximum(max, maximum(c.z, maximum(c.x, c.y)))
	}
	return min - 1, max + 1
}

func getCountActiveNeighbors3d(current map[point3d]bool, c point3d) (count int) {
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
				if current[neighbor] {
					count += 1
				}
			}
		}
	}
	return count
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
	// fmt.Printf("activeCells %v\n", activeCells)
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

func minimum(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
