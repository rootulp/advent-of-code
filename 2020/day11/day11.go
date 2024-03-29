package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting day11")

	// Part one
	result1 := GetCountOfOccupiedSeatsPartOne("input.txt")
	fmt.Printf("Part one: the number of occupied seats after the grid stabalizes is %v\n", result1)

	// Part two
	result2 := GetCountOfOccupiedSeatsPartTwo("input.txt")
	fmt.Printf("Part two: the number of occupied seats after the grid stabalizes is %v\n", result2)
}

// GetCountOfOccupiedSeatsPartOne returns the number of occupied seats after the grid stabalizes.
func GetCountOfOccupiedSeatsPartOne(filename string) int {
	lines := readFile(filename)
	grid := getGrid(lines)

	for i := 0; i < 100; i++ {
		grid = tickPartOne(grid)
	}

	return getCountOfOccupiedSeats(grid)
}

// GetCountOfOccupiedSeatsPartTwo returns the number of occupied seats after the grid stabalizes.
func GetCountOfOccupiedSeatsPartTwo(filename string) int {
	lines := readFile(filename)
	grid := getGrid(lines)

	for i := 0; i < 100; i++ {
		grid = tickPartTwo(grid)
	}

	return getCountOfOccupiedSeats(grid)
}

func tickPartOne(grid [][]gridValue) [][]gridValue {
	next := duplicateGrid(grid)

	for x, row := range grid {
		for y, val := range row {
			countOfOccupiedNeighbors := getCountOfOccupiedNeighbors(grid, x, y)
			if val == *registry.emptySeat && countOfOccupiedNeighbors == 0 {
				// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
				next[x][y] = *registry.occupiedSeat
			} else if val == *registry.occupiedSeat && countOfOccupiedNeighbors >= 4 {
				// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
				next[x][y] = *registry.emptySeat
			}
			// Otherwise, the seat's state does not change.
		}
	}

	// printGrid(next)
	return next
}

func tickPartTwo(grid [][]gridValue) [][]gridValue {
	next := duplicateGrid(grid)

	for x, row := range grid {
		for y, val := range row {
			countOfOccupiedVisible := getCountOfOccupiedVisible(grid, x, y)
			if val == *registry.emptySeat && countOfOccupiedVisible == 0 {
				next[x][y] = *registry.occupiedSeat
			} else if val == *registry.occupiedSeat && countOfOccupiedVisible >= 5 {
				next[x][y] = *registry.emptySeat
			}
		}
	}

	// printGrid(next)
	return next
}

func duplicateGrid(grid [][]gridValue) [][]gridValue {
	duplicate := make([][]gridValue, len(grid))
	for i := range grid {
		duplicate[i] = make([]gridValue, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}

func getCountOfOccupiedNeighbors(grid [][]gridValue, row int, col int) int {
	count := 0
	for _, diffRow := range []int{-1, 0, 1} {
		for _, diffCol := range []int{-1, 0, 1} {
			if diffRow == 0 && diffCol == 0 {
				continue
			}
			if row+diffRow >= len(grid) || row+diffRow < 0 {
				continue
			}
			if col+diffCol >= len(grid[0]) || col+diffCol < 0 {
				continue
			}
			neighbor := grid[row+diffRow][col+diffCol]
			if neighbor == *registry.occupiedSeat {
				count++
			}
		}
	}

	return count
}

func getCountOfOccupiedVisible(grid [][]gridValue, row int, col int) int {
	count := 0
	// fmt.Printf("Getting count of visible for (%d, %d)\n", row, col)
	for _, directionRow := range []int{-1, 0, 1} {
		for _, directionCol := range []int{-1, 0, 1} {
			// fmt.Printf("Visible direction (%d, %d)\n", directionRow, directionCol)
			if directionRow == 0 && directionCol == 0 {
				continue
			}
			diffRow := directionRow
			diffCol := directionCol
			for row+diffRow < len(grid) && row+diffRow >= 0 &&
				col+diffCol < len(grid[0]) && col+diffCol >= 0 {
				visible := grid[row+diffRow][col+diffCol]
				if visible == *registry.occupiedSeat {
					// fmt.Printf("(%d, %d) is occupied seat\n", row+diffX, col+diffY)
					count++
					break
				} else if visible == *registry.emptySeat {
					// fmt.Printf("(%d, %d) is empty seat\n", row+diffX, col+diffY)
					break
				} else if visible == *registry.floor {
					// fmt.Printf("(%d, %d) is floor\n", row+diffX, col+diffY)
					diffRow += directionRow
					diffCol += directionCol
				}
			}
		}
	}

	return count
}

func getCountOfOccupiedSeats(grid [][]gridValue) int {
	count := 0
	for _, row := range grid {
		for _, val := range row {
			if val == *registry.occupiedSeat {
				count++
			}
		}
	}
	return count
}

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

func printGrid(grid [][]gridValue) {
	fmt.Print(toString(grid))
}

func toString(grid [][]gridValue) (result string) {
	for _, line := range grid {
		for _, r := range line {
			result += string(r.characterRepresentation)
		}
		result += "\n"
	}
	result += "\n"

	return result
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
