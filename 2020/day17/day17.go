package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const gridSize = 3

type state struct {
	grid [gridSize][gridSize][gridSize]rune
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

func New() (s state) {
	grid := emptyGrid()
	return state{
		grid: grid,
	}
}

// NextCycle advances state by one cycle
func (s state) NextCycle() (nextState state) {
	nextState = New()
	return nextState
}

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
	state := newState(lines)
	fmt.Printf("initialState: \n%v\n", state)
	return 0
}

func newState(lines []string) (s state) {
	zeroIndexSlice := getTwoDimensionalSlice(lines)
	grid := emptyGrid()
	grid[1] = zeroIndexSlice
	return state{
		grid: grid,
	}
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

// getTwoDimensionalSlice returns a two dimensional slice for the input
func getTwoDimensionalSlice(lines []string) (result [3][3]rune) {
	result = [3][3]rune{}
	for x, line := range lines {
		for y, r := range line {
			result[x][y] = r
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
