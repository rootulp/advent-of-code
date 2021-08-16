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
	s = state{[gridSize][gridSize][gridSize]rune{
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
