package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	row int
	col int
}

type Tile struct {
	id        int
	contents  []string
	neighbors map[int]string // map from tile id to side
}

func (t Tile) String() string {
	return fmt.Sprintf("Tile: %v\nContents:\n%v\nNeighbors: %v\n\n", t.id, strings.Join(t.contents, "\n"), t.neighbors)
}

func (t Tile) getAllSides() (allSides []string) {
	allSides = append(allSides, t.getSides()...)
	allSides = append(allSides, t.getSidesReversed()...)
	return allSides
}

func (t Tile) getSides() []string {
	return []string{
		t.getSide("top"),
		t.getSide("bottom"),
		t.getSide("left"),
		t.getSide("right"),
	}
}

func (t Tile) getSide(side string) string {
	switch side {
	case "top":
		return t.contents[0]
	case "bottom":
		return t.contents[len(t.contents)-1]
	case "right":
		return t.getRightSide()
	case "left":
		return t.getLeftSide()
	default:
		log.Fatalf("unrecognized side %v", side)
	}
	return ""
}

func (t Tile) getSidesReversed() (reversed []string) {
	for _, border := range t.getSides() {
		reversed = append(reversed, reverse(border))
	}
	return reversed
}

func (t Tile) getLeftSide() (result string) {
	for _, row := range t.contents {
		vals := strings.Split(row, "")
		result += vals[0]
	}
	return result
}

func (t Tile) getRightSide() (result string) {
	for _, row := range t.contents {
		vals := strings.Split(row, "")
		result += vals[len(vals)-1]
	}
	return result
}

func (t *Tile) rotate() {
	var rotated [][]string = [][]string{}
	// var R = len(t.contents)
	var C = len(t.contents[0])

	for r, row := range t.contents {
		for c := range row {
			rotated[r][c] = string(t.contents[C-c-1][r])
		}
	}

	var newContents []string
	for _, r := range rotated {
		newContents = append(newContents, strings.Join(r, ""))
	}
	t.contents = newContents
}

func (t *Tile) flip() {
	t.contents = reverseSlice(t.contents)
}

func main() {
	fmt.Printf("Starting day20...\n")

	partOne := PartOne("input.txt")
	fmt.Printf("Part one: %v\n", partOne)

	partTwo := PartTwo("example.txt")
	fmt.Printf("Part one: %v\n", partTwo)
}

func PartOne(filename string) (productOfCornerIds int) {
	lines := readLines(filename)
	tiles := parseTiles(lines)

	// fmt.Printf("tiles: %v\n", tiles)
	// fmt.Printf("len(tiles): %v\n", len(tiles))
	// fmt.Printf("borders for first tile: %v\n", tiles[0].getSides())

	occurences := countBorderOccurences(tiles)
	// fmt.Printf("occurences %v\n", occurences)

	cornerTileIds := cornerTiles(tiles, occurences)
	// fmt.Printf("cornerTileIds %v\n", cornerTileIds)

	productOfCornerIds = 1
	for _, cornerTileId := range cornerTileIds {
		productOfCornerIds *= cornerTileId
	}

	return productOfCornerIds
}

func PartTwo(filename string) (numberOfPoundSignsNotPartOfSeaMonsters int) {
	lines := readLines(filename)
	tiles := parseTiles(lines)
	fmt.Printf("tiles %v\n", tiles)
	// tilesMap := getTilesMap(tiles)
	// occurences := countBorderOccurences(tiles)
	// cornerTileIds := cornerTiles(tiles, occurences)

	// pic := assemble(tilesMap, cornerTileIds)
	// fmt.Printf("pic\n%v\n", pic)

	return numberOfPoundSignsNotPartOfSeaMonsters
}

// cornerTiles returns the tile Ids for all corner tiles
func cornerTiles(tiles []Tile, occurences map[string]int) (cornerTileIds []int) {
	tileToSharedBorders := map[int]int{}
	for _, tile := range tiles {
		numSharedBorders := 0
		borders := tile.getSides()
		for _, border := range borders {
			numSharedBorders += (occurences[border])
		}
		reversedBorders := tile.getSidesReversed()
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
	// fmt.Printf("tileToSharedBorders %v\n", tileToSharedBorders)
	return cornerTileIds
}

func countBorderOccurences(tiles []Tile) (occurences map[string]int) {
	occurences = map[string]int{}
	for _, tile := range tiles {
		borders := tile.getSides()
		for _, border := range borders {
			occurences[border] += 1
		}
		reversedBorders := tile.getSidesReversed()
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
			tile = Tile{id, []string{}, map[int]string{}}
		} else if strings.HasPrefix(line, ".") || strings.HasPrefix(line, "#") {
			tile.contents = append(tile.contents, line)
		} else {
			tiles = append(tiles, tile)
		}
	}
	tiles = append(tiles, tile)

	for a, tileA := range tiles {
		for b, tileB := range tiles {
			if a == b {
				continue
			}
			sharedSides := getSharedSides(tileA.getAllSides(), tileB.getAllSides())
			for _, side := range sharedSides {
				tileA.neighbors[tileB.id] = side
				tileB.neighbors[tileA.id] = side
			}
		}
	}
	return tiles
}

func getSharedSides(sidesA []string, sidesB []string) (shared []string) {
	for _, a := range sidesA {
		for _, b := range sidesB {
			if a == b {
				shared = append(shared, a)
			}
		}
	}
	return shared
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

func assemble(tilesMap map[int]Tile, cornerTileIds []int) (fullpic []string) {
	var r = 0
	var c = 0
	var gridmap map[Point]int
	// var fullpic = []string{}

	for len(gridmap) < len(tilesMap) {
		if r == 0 && c == 0 {
			cornerTileId := cornerTileIds[0]
			cornerTile := tilesMap[cornerTileId]
			fmt.Printf("cornerTile %v\n", cornerTile)

			// neighbors = list(tiles[idn]["neighbors"].values())
			// neighbors += [n[::-1] for n in neighbors]
			// while True:
			// 	bottom, right = get_side(tile, "bottom"), get_side(tile, "right")
			// 	if bottom in neighbors and right in neighbors:
			// 		break
			// 	tile = rotate(tile)
			// gridmap[(r, c)] = idn
			// tiles[idn]["grid"] = tile
			// add_fullpic(tile)

			c += 1
		} else if c == 0 {
			// 	# add new tile based on first tile in previous row's bottom
			// 	...[snip]...
			c += 1
		} else {
			// 	# look for next in row
			// 	...[snip]...
			// 	if find next:
			// 		# add to row
			// 		...[snip]...
			// 		c += 1
			// 	else:
			// 		r, c = r + 1, 0
		}
	}
	return fullpic
}

func getTilesMap(tiles []Tile) (result map[int]Tile) {
	result = map[int]Tile{}
	for _, tile := range tiles {
		result[tile.id] = tile
	}
	return result
}

func addTileToPic(pic []string, tile Tile, newline bool) (newPic []string) {
	copy(newPic, pic)

	if len(newPic) == 0 {
		copy(newPic, tile.contents)
		return newPic
	}
	if newline {
		newPic = append(newPic, tile.contents...)
		return newPic
	}
	var R = len(newPic) - len(tile.contents)
	for index, row := range tile.contents {
		newPic[R+index] += row
	}
	return newPic
}

func reverse(input string) (reversed string) {
	for _, r := range input {
		reversed = string(r) + reversed
	}
	return reversed
}

func reverseSlice(input []string) (reversed []string) {
	for _, r := range input {
		reversed = append([]string{r}, reversed...)
	}
	return reversed
}
