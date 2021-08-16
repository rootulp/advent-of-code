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
	fmt.Printf("initialState %v\n", state)
	return 0
}

func newState(lines []string) (s state) {
	zeroIndexSlice := getTwoDimensionalSlice(lines)
	s = state{[gridSize][gridSize][gridSize]rune{
		{ // z = -1
			{'.', '.', '.'},
			{'.', '.', '.'},
			{'.', '.', '.'},
		},
		zeroIndexSlice,
		{ // z = 1
			{'.', '.', '.'},
			{'.', '.', '.'},
			{'.', '.', '.'},
		},
	}}
	return s
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
