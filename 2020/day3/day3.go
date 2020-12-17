package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting Toboggan Trajectory")
	grid := readFile("./input.txt")

	currentRow := 0
	currentCol := 0
	numTreesEncountered := 0

	for currentRow < len(grid) {
		currentPos := grid[currentRow][currentCol]
		if isTree(currentPos) {
			numTreesEncountered++
		}
		fmt.Printf("currentRow %v, currentCol %v, currentPos %v, numTreesEncountered %v", currentRow, currentCol, currentPos, numTreesEncountered)
		currentRow++
		currentCol = (currentCol + 3) % len(grid[0])
	}

	fmt.Printf("Total # of trees encountered: %d", numTreesEncountered)
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
