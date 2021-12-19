package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	state := NewState(lines)
	fmt.Printf("initialState: \n%v\n", state)
	return 0
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
