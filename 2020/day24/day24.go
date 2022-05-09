package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Starting day24...\n")

	partOne := PartOne("input.txt")
	fmt.Printf("PartOne: %v\n", partOne)

	partTwo := PartTwo("input.txt", 100)
	fmt.Printf("PartTwo: %v\n", partTwo)
}

func PartOne(filename string) (numBlackTiles int) {
	lines := readLines(filename)
	floor := NewFloor()

	for _, line := range lines {
		point := getPoint(line)
		floor.Flip(point)
	}
	fmt.Println(floor)

	return floor.NumBlackTiles()
}

func PartTwo(filename string, numDays int) (numBlackTiles int) {
	lines := readLines(filename)
	floor := NewFloor()

	for _, line := range lines {
		point := getPoint(line)
		floor.Flip(point)
	}

	// fmt.Printf("Day 0\n")
	// fmt.Println(floor)

	for day := 1; day <= numDays; day += 1 {
		floor.AdvanceOneDay()

		// fmt.Printf("Day %d\n", day)
		// fmt.Printf("Num black tiles %v\n", floor.NumBlackTiles())
		// fmt.Println(floor)
	}

	return floor.NumBlackTiles()
}

func readLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func getPoint(line string) (Point) {
	x := 0
	y := 0
	for _, token := range tokenize(line) {
		switch token {
		case "nw":
			x -= 1
			y += 1
		case "ne":
			x += 1
			y += 1
		case "sw":
				x -= 1
				y -= 1
		case "se":
				x += 1
				y -= 1
		case "e":
				x += 2
		case "w":
				x -= 2
		default:
				log.Fatalf("unrecognized token %v", token)
		}
	}
	return Point{x, y}
}

func tokenize(line string) (tokens []string) {
	for line != "" {
		if strings.HasPrefix(line, "nw") {
			tokens = append(tokens, "nw")
			line = strings.TrimPrefix(line, "nw")
		}
		if strings.HasPrefix(line, "ne") {
			tokens = append(tokens, "ne")
			line = strings.TrimPrefix(line, "ne")
		}
		if strings.HasPrefix(line, "sw") {
			tokens = append(tokens, "sw")
			line = strings.TrimPrefix(line, "sw")
		}
		if strings.HasPrefix(line, "se") {
			tokens = append(tokens, "se")
			line = strings.TrimPrefix(line, "se")
		}
		if strings.HasPrefix(line, "e") {
			tokens = append(tokens, "e")
			line = strings.TrimPrefix(line, "e")
		}
		if strings.HasPrefix(line, "w") {
			tokens = append(tokens, "w")
			line = strings.TrimPrefix(line, "w")
		}
	}
	return tokens
}

type Floor struct {
	tiles map[Point]*Tile
}

func NewFloor() (*Floor) {
	tiles := map[Point]*Tile{}
	return &Floor{tiles}
}

func (f *Floor) Flip(point Point) {
	if _, ok := f.tiles[point]; !ok {
		f.tiles[point] = NewTile()
	}
	tile := f.tiles[point]
	tile.Flip()
}

func (f *Floor) NumBlackTiles() (result int) {
	for _, tile := range f.tiles {
		if tile.IsBlack() {
			result += 1
		}
	}
	return result
}

func (f *Floor) AdvanceOneDay() {
	f.populateMissingTiles()

	tilesToFlip := []*Tile{}
	for point, tile := range f.tiles{
		if f.shouldFlip(point, tile) {
			tilesToFlip = append(tilesToFlip, tile)
		}
	}

	for _, tile := range tilesToFlip {
		tile.Flip()
	}
}

func (f *Floor) numBlackNeighbors(point Point, tile *Tile) (result int) {
	neighbors := f.neighbors(point, tile)
	for _, neighbor := range neighbors {
		if neighbor.IsBlack() {
			result += 1
		}
	}
	return result
}

func (f *Floor) neighbors(point Point, tile *Tile) (result []*Tile) {
	possibleNeighbors := possibleNeighbors(point)

	for point, tile := range f.tiles {
		if contains(possibleNeighbors, point) {
			result = append(result, tile)
		}
	}
	return result
}

func (f *Floor) shouldFlip(point Point, tile *Tile) bool {
	if tile.IsBlack() && (f.numBlackNeighbors(point, tile) == 0 || f.numBlackNeighbors(point, tile) > 2) {
		return true
	}
	if tile.IsWhite() && (f.numBlackNeighbors(point, tile) == 2) {
		return true
	}
	return false
}

func (f *Floor) blackPoints() (result []Point) {
	for point, tile := range f.tiles {
		if tile.IsBlack() {
			result = append(result, point)
		}
	}
	return result
}

func (f *Floor) populateMissingTiles() {
	blackPoints := f.blackPoints()
	for _, point := range blackPoints {
		neighbors := possibleNeighbors(point)
		for _, neighbor := range neighbors {
			if _, ok := f.tiles[neighbor]; !ok {
				f.tiles[neighbor] = NewTile()
			}
		}
	}
}

func (f *Floor) String() (result string) {
	for point, tile := range f.tiles {
		result += fmt.Sprintf("%v: %v\n", point, tile)
	}
	return result
}

type Point struct {
	x int
	y int
}

type Tile struct {
	isBlack bool
}

func NewTile() (*Tile) {
	// tiles start off white
	return &Tile{isBlack: false}
}

func (t *Tile) Flip() {
	t.isBlack = !t.isBlack
}

func (t *Tile) IsBlack() bool {
	return t.isBlack
}

func (t *Tile) IsWhite() bool {
	return !t.isBlack
}

func (t *Tile) String() string {
	if t.IsBlack() {
		return "black"
	} else {
		return "white"
	}
}

// Utils

func contains(points []Point, val Point) bool {
	for _, point := range points {
		if point == val {
			return true
		}
	}
	return false
}

func possibleNeighbors(p Point) (result []Point) {
	return []Point{
		{p.x - 1, p.y - 1},
		{p.x + 1, p.y + 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y + 1},
		{p.x + 2, p.y},
		{p.x - 2, p.y},
	}
}
