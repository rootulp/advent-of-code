package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type GridValue struct {
	CharacterRepresentation rune
}

type GridRegistry struct {
	floor        *GridValue
	emptySeat    *GridValue
	occupiedSeat *GridValue
	tokens       []*GridValue
}

func newGridRegistry() *GridRegistry {
	floor := &GridValue{'.'}
	emptySeat := &GridValue{'L'}
	occupiedSeat := &GridValue{'#'}

	return &GridRegistry{
		floor:        floor,
		emptySeat:    emptySeat,
		occupiedSeat: occupiedSeat,
		tokens:       []*GridValue{floor, emptySeat, occupiedSeat},
	}
}

func (g *GridRegistry) List() []*GridValue {
	return g.tokens
}

func (g *GridRegistry) Parse(r rune) (*GridValue, error) {
	for _, token := range g.List() {
		if token.CharacterRepresentation == r {
			return token, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("couldn't find rune %v in list %v", r, g.List()))
}

var gridRegistry = newGridRegistry()

func main() {
	fmt.Println("Starting day11")

	// Part one
	lines := readFile("input_test.txt")
	grid := getGrid(lines)
	printGrid(grid)
}

func printGrid(grid [][]GridValue) {
	for _, line := range grid {
		for _, r := range line {
			fmt.Printf("%c", r.CharacterRepresentation)
		}
		fmt.Println()
	}
}

// getGrid converts a slice of lines into a matrix of gridValues
func getGrid(lines []string) [][]GridValue {
	grid := [][]GridValue{}
	for _, line := range lines {
		gridLine := []GridValue{}
		for _, ch := range line {
			gridValue, err := gridRegistry.Parse(ch)
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
