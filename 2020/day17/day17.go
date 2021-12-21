package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cell struct {
	z int
	x int
	y int
}

const ACTIVE_CELL = '#'

func main() {
	fmt.Println("Starting day17")

	// Part One
	partOne := PartOne("example.txt")
	fmt.Printf("Part one: %v\n", partOne)
}

// PartOne returns the number of active cubes after simulating six cycles of
// life for the Conway Cubes given initial state in filename
func PartOne(filename string) int {
	lines := readFile(filename)
	activeCells := setActiveCells(lines)
	for i := 0; i < 6; i++ {
		fmt.Printf("i=%d, len(activeCells)=%d\n", i, len(activeCells))
		activeCells = nextCycle(activeCells)
	}
	return len(activeCells)
}

func nextCycle(current map[cell]bool) (next map[cell]bool) {
	next = make(map[cell]bool)
	min, max := getBounds(current)
	for z := min; z <= max; z++ {
		for x := min; x <= max; x++ {
			for y := min; y <= max; y++ {
				c := cell{z, x, y}
				activeNeighbors := getCountActiveNeighbors(current, c)
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

func getBounds(current map[cell]bool) (min int, max int) {
	for c := range current {
		min = minimum(min, minimum(c.z, minimum(c.x, c.y)))
		max = maximum(max, maximum(c.z, maximum(c.x, c.y)))
	}
	return min - 1, max + 1
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

func getCountActiveNeighbors(current map[cell]bool, c cell) (count int) {
	for dz := -1; dz <= 1; dz++ {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dz == 0 && dx == 0 && dy == 0 {
					continue
				}
				neighbor := cell{
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

func setActiveCells(lines []string) (activeCells map[cell]bool) {
	activeCells = make(map[cell]bool)
	for x, line := range lines {
		for y, r := range line {
			if r == ACTIVE_CELL {
				c := cell{
					z: 0, // every element in input is at z=0
					x: x,
					y: y,
				}
				activeCells[c] = true
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
