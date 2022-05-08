package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)


func PartOne(filename string) (numBlackTiles int) {
	lines := readLines(filename)
	floor := Floor{}

	for _, line := range lines {
		point := getPoint(line)
		floor.Flip(point)
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
				y -= 2
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
	tiles map[Point]Tile
}

func (f Floor) Flip(p Point) {
	// TODO
}

func (f Floor) NumBlackTiles() (result int) {
	for _, tile := range f.tiles {
		if tile.IsBlack() {
			result += 1
		}
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

func (t *Tile) Flip() {
	t.isBlack = !t.isBlack
}

func (t *Tile) IsBlack() bool {
	return t.isBlack
}
