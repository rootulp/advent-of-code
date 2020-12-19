package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"unicode/utf8"
)

func main() {
	boardingPasses := ReadFile("input.txt")

	// Part one
	maxSeatID := GetMaxSeatID(boardingPasses)
	fmt.Printf("Max boarding pass seatID %v\n", maxSeatID)

	// Part two
	PrintOccupiedSeats(boardingPasses)
	// Unoccupied seat is row 64 col 7 = seatId 519
}

// PrintOccupiedSeats prints a map from row => col that describes the occupied
// seats.
func PrintOccupiedSeats(boardingPasses []string) {
	occupied := make(map[int][]int)
	for _, boardingPass := range boardingPasses {
		row, col := GetRowAndCol(boardingPass)
		newRow := append(occupied[row], col)
		sort.Ints(newRow)
		occupied[row] = newRow
	}

	// To store the keys in occupied in sorted order
	keys := make([]int, len(occupied))
	i := 0
	for k := range occupied {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	for _, row := range keys {
		fmt.Printf("Row %v: %v\n", row, occupied[row])
	}
}

// GetMaxSeatID returns the max seatId from the list of provided boarding
// passes.
func GetMaxSeatID(boardingPasses []string) int {
	maxSeatID := 0
	for _, boardingPass := range boardingPasses {
		seatID := GetSeatID(boardingPass)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	return maxSeatID
}

// ReadFile reads a file line by line
func ReadFile(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	lines = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

// GetSeatID gets the seatId for the provided boarding pass.
func GetSeatID(boardingPass string) (seatID int) {
	row, col := GetRowAndCol(boardingPass)
	return (row * 8) + col
}

// GetRowAndCol gets the row and col for a boarding pass.
func GetRowAndCol(boardingPass string) (row int, col int) {
	boardingRow := boardingPass[:7]
	boardingCol := boardingPass[7:]
	// log.Printf("boardingPass %v boardingRow %v boardingCol %v\n", boardingPass, boardingRow, boardingCol)

	row = GetRow(boardingRow)
	col = GetCol(boardingCol)
	return
}

// GetRow returns a number between 0 and 127 (inclusive) that represents the row
// number associated with the provided boardingRow.
func GetRow(boardingRow string) (row int) {
	rows := makeRange(0, 127)
	return binarySearch(rows, 0, 127, boardingRow)
}

func binarySearch(rows []int, min int, max int, searchString string) int {
	// log.Printf("binarySearch invoked with min %v max %v searchString %v\n", min, max, searchString)
	if min > max || max < min {
		log.Fatalf("binarySearch invoked with invalid rows %v min %v max %v searchString %v\n", rows, min, max, searchString)
	}
	if min == max {
		result := rows[min]
		// log.Printf("binarySearch completed successfully with result %v\n", result)
		return result
	}

	remainingString, currentSearch := trimFirstRune(searchString)
	midpoint := ((max - min) / 2) + min
	if currentSearch == 'F' || currentSearch == 'L' {
		return binarySearch(rows, min, midpoint, remainingString)
	} else if currentSearch == 'B' || currentSearch == 'R' {
		return binarySearch(rows, midpoint+1, max, remainingString)
	} else {
		log.Fatalf("binarySearch invoked with invalid searchString %v currentSearch %v remainingString %v\n", searchString, currentSearch, remainingString)
	}
	return 0
}

// GetCol returns a number between 0 and 7 (inclusive) that represents the col
// number associated with the provided boardingCol.
func GetCol(boardingCol string) (col int) {
	rows := makeRange(0, 7)
	return binarySearch(rows, 0, 7, boardingCol)
}

func makeRange(min int, max int) []int {
	result := make([]int, max-min+1)
	for i := range result {
		result[i] = min + i
	}
	return result
}

func trimFirstRune(s string) (string, rune) {
	r, i := utf8.DecodeRuneInString(s)
	return s[i:], r
}
