package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	id       int
	contents []string
}

func (t Tile) borders() []string {
	return []string{
		t.contents[0],
		t.contents[len(t.contents)-1],
		t.leftBorder(),
		t.rightBorder(),
	}
}

func (t Tile) leftBorder() (result string) {
	for _, row := range t.contents {
		vals := strings.Split(row, "")
		result += vals[0]
	}
	return result
}

func (t Tile) rightBorder() (result string) {
	for _, row := range t.contents {
		vals := strings.Split(row, "")
		result += vals[len(vals)-1]
	}
	return result
}

func main() {
	fmt.Printf("Starting day20...\n")

	partOne := PartOne("example.txt")
	fmt.Printf("Part one: %v\n", partOne)
}

func PartOne(filename string) int {
	lines := readLines(filename)
	tiles := parseTiles(lines)

	fmt.Printf("tiles: %v\n", tiles)
	fmt.Printf("borders for first tile: %v\n", tiles[0].borders())

	borderOccurrences := countBorderOccurences(tiles)
	fmt.Printf("borderOccurences %v\n", borderOccurrences)

	return 0
}

func countBorderOccurences(tiles []Tile) (occurences map[string]int) {
	occurences = map[string]int{}
	for _, tile := range tiles {
		borders := tile.borders()
		for _, border := range borders {
			occurences[border] += 1
		}
	}
	return occurences
}

func parseTiles(lines []string) (tiles []Tile) {
	var tile Tile
	for _, line := range lines {
		if strings.HasPrefix(line, "Tile ") {
			tileId := strings.TrimPrefix(line, "Tile ")
			tileId = strings.TrimSuffix(tileId, ":")
			id, err := strconv.Atoi(tileId)
			if err != nil {
				log.Fatal(err)
			}
			tile = Tile{id: id}
		} else if strings.HasPrefix(line, ".") || strings.HasPrefix(line, "#") {
			tile.contents = append(tile.contents, line)
		} else {
			tiles = append(tiles, tile)
		}
	}
	return tiles
}

func readLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
