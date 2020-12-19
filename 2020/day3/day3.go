package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting Toboggan Trajectory")
	grid := readFile("input.txt")
	trees1 := GetNumTreesEncountered(grid, 1, 1)
	trees2 := GetNumTreesEncountered(grid, 3, 1)
	trees3 := GetNumTreesEncountered(grid, 5, 1)
	trees4 := GetNumTreesEncountered(grid, 7, 1)
	trees5 := GetNumTreesEncountered(grid, 1, 2)
	result := trees1 * trees2 * trees3 * trees4 * trees5
	fmt.Printf("Product of trees encountered for different trajectories %d\n", result)
}

// GetNumTreesEncountered gets the number of trees encountered when navigating
// the provided grid from the (0, 0) to the last row.
func GetNumTreesEncountered(grid [][]rune, trajectoryRight int, trajectoryDown int) int {
	currentRow := 0
	currentCol := 0
	numTreesEncountered := 0

	for currentRow < len(grid) {
		currentPos := grid[currentRow][currentCol]
		if isTree(currentPos) {
			numTreesEncountered++
		}
		// fmt.Printf("currentRow %v, currentCol %v, currentPos %v, numTreesEncountered %v\n", currentRow, currentCol, currentPos, numTreesEncountered)
		currentRow += trajectoryDown
		currentCol = (currentCol + trajectoryRight) % len(grid[0])
	}

	fmt.Printf("%d trees encountered for trajectoryRight %d and trajectoryDown %d\n", numTreesEncountered, trajectoryRight, trajectoryDown)
	return numTreesEncountered
}

func isTree(char rune) bool {
	const tree = '#'
	return char == tree
}

func readFile(filename string) (grid [][]rune) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" { // Reached end of file
			file.Close()
			return grid
		}

		grid = append(grid, []rune(text))
	}

	return grid
}
