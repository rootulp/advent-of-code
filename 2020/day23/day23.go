package main

import (
	"fmt"
	"log"
	"strconv"
)

const NUM_CUPS_TO_PICK_UP = 3
const EXAMPLE_INPUT = "389125467"
const INPUT = "685974213"

func main() {
	fmt.Printf("Starting day23...\n")

	partOne := PartOne(INPUT, 100)
	fmt.Printf("PartOne %v\n", partOne)

	partTwo := PartTwo(INPUT, 10_000_000)
	fmt.Printf("PartTwo %v\n", partTwo)
}

// PartOne retruns the order of the cups after applying numMoves
func PartOne(input string, numMoves int) (result string) {
	totalCups := len(input)
	cups := parseCups(input)
	game := NewGame(cups, totalCups, numMoves)

	for moveNumber := 1; moveNumber <= numMoves; moveNumber += 1 {
		fmt.Printf("-- move %d --\n", moveNumber)
		fmt.Println(game)
		game.Move()
	}

	fmt.Printf("-- final --\n")
	fmt.Println(game)

	return game.CupOrderWithoutOne()
}

// PartTwo returns the product of the two cups immediately clockwise of cup 1 after applying numMoves
func PartTwo(input string, numMoves int) (product int) {
	totalCups := 1_000_000
	cups := parseCups(input)
	game := NewGame(cups, totalCups, numMoves)

	for moveNumber := 1; moveNumber <= numMoves; moveNumber += 1 {
		game.Move()
	}

	return game.ProductOfTwoCupsAfterCupOne()
}

type Game struct {
	cupToNextCup map[int]int
	pointer      int
	totalCups    int
}

func NewGame(cups []int, totalCups int, numMoves int) (result *Game) {
	cupToNextCup := map[int]int{}

	for i := 0; i < totalCups; i += 1 {
		if i < len(cups)-1 {
			cupToNextCup[cups[i]] = cups[i+1]
		} else if i == len(cups)-1 && len(cups) == totalCups {
			// reached last cup, point to first cup
			cupToNextCup[cups[i]] = cups[0]
		} else if i == len(cups)-1 && len(cups) < totalCups {
			cupToNextCup[cups[i]] = getMax(cups) + 1
		} else if i < totalCups-1 {
			cupToNextCup[i+1] = i + 2
		} else if i == totalCups-1 {
			cupToNextCup[i+1] = cups[0]
		}
	}
	return &Game{cupToNextCup, cups[0], totalCups}
}

func (g *Game) String() (result string) {
	result += fmt.Sprintf("cupToNextCup: (%v)\n", g.cupToNextCup)
	result += fmt.Sprintf("cups: (%v) ", g.pointer)

	current := g.cupToNextCup[g.pointer]
	for current != g.pointer {
		result += fmt.Sprintf("%d ", current)
		current = g.cupToNextCup[current]
	}
	return result
}

func (g *Game) Move() {
	// remove three cups
	cup1 := g.cupToNextCup[g.pointer]
	cup2 := g.cupToNextCup[cup1]
	cup3 := g.cupToNextCup[cup2]
	g.cupToNextCup[g.pointer] = g.cupToNextCup[cup3]
	// fmt.Printf("pick up: %v, %v, %v\n", cup1, cup2, cup3)

	// find destination cup
	destination := decrementInRange(1, g.totalCups, g.pointer)
	for includes([]int{cup1, cup2, cup3}, destination) {
		destination = decrementInRange(1, g.totalCups, destination)
	}
	// fmt.Printf("destination: %v\n\n", destination)

	// reinsert cups after dest
	g.cupToNextCup[cup3] = g.cupToNextCup[destination]
	g.cupToNextCup[destination] = cup1

	// move pointer forward
	g.pointer = g.cupToNextCup[g.pointer]
}

func (g *Game) CupOrderWithoutOne() (result string) {
	current := g.cupToNextCup[1]
	for current != 1 {
		result += fmt.Sprintf("%d", current)
		current = g.cupToNextCup[current]
	}
	return result
}

func (g *Game) ProductOfTwoCupsAfterCupOne() (product int) {
	firstCup := g.cupToNextCup[1]
	secondCup := g.cupToNextCup[firstCup]

	return firstCup * secondCup
}

func decrementInRange(min int, max int, val int) (decremented int) {
	// fmt.Printf("decrementInRange(%v, %v, %v)\n", min, max, val)
	if val-1 < min {
		return max
	}
	return val - 1
}

func parseCups(input string) (cups []int) {
	for _, r := range input {
		cup, err := strconv.Atoi(string(r))
		if err != nil {
			log.Fatal(err)
		}
		cups = append(cups, cup)
	}
	return cups
}

func getMax(list []int) (max int) {
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}

func includes(list []int, val int) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}
