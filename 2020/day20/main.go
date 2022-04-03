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

func (t Tile) reversedBorders() (reversed []string) {
	for _, border := range t.borders() {
		reversed = append(reversed, reverse(border))
	}
	return reversed
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

func PartOne(filename string) (productOfCornerIds int) {
	lines := readLines(filename)
	tiles := parseTiles(lines)

	fmt.Printf("tiles: %v\n", tiles)
	fmt.Printf("len(tiles): %v\n", len(tiles))
	fmt.Printf("borders for first tile: %v\n", tiles[0].borders())

	occurences := countBorderOccurences(tiles)
	fmt.Printf("occurences %v\n", occurences)

	cornerTileIds := cornerTiles(tiles, occurences)
	fmt.Printf("cornerTileIds %v\n", cornerTileIds)

	productOfCornerIds = 1
	for _, cornerTileId := range cornerTileIds {
		productOfCornerIds *= cornerTileId
	}

	return productOfCornerIds
}

// cornerTiles returns the tile Ids for all corner tiles
func cornerTiles(tiles []Tile, occurences map[string]int) (cornerTileIds []int) {
	tileToSharedBorders := map[int]int{}
	for _, tile := range tiles {
		numSharedBorders := 0
		borders := tile.borders()
		for _, border := range borders {
			numSharedBorders += (occurences[border])
		}
		reversedBorders := tile.reversedBorders()
		for _, reversedBorder := range reversedBorders {
			numSharedBorders += (occurences[reversedBorder])
		}
		tileToSharedBorders[tile.id] = numSharedBorders
		// cornerTiles have two shared borders and two unique borders
		// (2 * 2) + (2 * 1) == 6
		// However we are double counting each border (because reversed borders) so
		// 2 * 6 == 12
		if numSharedBorders == 12 {
			cornerTileIds = append(cornerTileIds, tile.id)
		}
	}
	fmt.Printf("tileToSharedBorders %v\n", tileToSharedBorders)
	return cornerTileIds
}

func countBorderOccurences(tiles []Tile) (occurences map[string]int) {
	occurences = map[string]int{}
	for _, tile := range tiles {
		borders := tile.borders()
		for _, border := range borders {
			occurences[border] += 1
		}
		reversedBorders := tile.reversedBorders()
		for _, reversedBorder := range reversedBorders {
			occurences[reversedBorder] += 1
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
	tiles = append(tiles, tile)
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

func reverse(input string) (reversed string) {
	for _, r := range input {
		reversed = string(r) + reversed
	}
	return reversed
}
