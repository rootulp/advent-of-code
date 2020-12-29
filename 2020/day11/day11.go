package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type gridValue struct {
	characterRepresentation rune
}

type gridRegistry struct {
	floor        *gridValue
	emptySeat    *gridValue
	occupiedSeat *gridValue
	tokens       []*gridValue
}

func newGridRegistry() *gridRegistry {
	floor := &gridValue{'.'}
	emptySeat := &gridValue{'L'}
	occupiedSeat := &gridValue{'#'}

	return &gridRegistry{
		floor:        floor,
		emptySeat:    emptySeat,
		occupiedSeat: occupiedSeat,
		tokens:       []*gridValue{floor, emptySeat, occupiedSeat},
	}
}

func (g *gridRegistry) List() []*gridValue {
	return g.tokens
}

func (g *gridRegistry) Parse(r rune) (*gridValue, error) {
	for _, token := range g.List() {
		if token.characterRepresentation == r {
			return token, nil
		}
	}
	return nil, fmt.Errorf("couldn't find rune %v in list %v", r, g.List())
}

var registry = newGridRegistry()

func main() {
	fmt.Println("Starting day11")

	// Part one
	lines := readFile("input_test.txt")
	grid := getGrid(lines)
	printGrid(grid)
}

// GetCountOfOccupiedSeats returns the number of occupied seats after the grid stabalizes.
func GetCountOfOccupiedSeats(filename string) int {
	return 0
}

func printGrid(grid [][]gridValue) {
	for _, line := range grid {
		for _, r := range line {
			fmt.Printf("%c", r.characterRepresentation)
		}
		fmt.Println()
	}
}

// getGrid converts a slice of lines into a matrix of gridValues
func getGrid(lines []string) [][]gridValue {
	grid := [][]gridValue{}
	for _, line := range lines {
		gridLine := []gridValue{}
		for _, ch := range line {
			gridValue, err := registry.Parse(ch)
			if err != nil {
				log.Fatal(err)
			}
			gridLine = append(gridLine, *gridValue)
		}
		grid = append(grid, gridLine)
	}
	return grid
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
