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

	partTwo := PartTwo(EXAMPLE_INPUT, 10_000_000)
	fmt.Printf("PartTwo %v\n", partTwo)
}

// PartOne retruns the order of the cups after applying numMoves
func PartOne(input string, numMoves int) (result string) {
	shouldAppendCupsUpToOneMillion := false
	game := NewGame(input, shouldAppendCupsUpToOneMillion)

	for move := 1; move <= numMoves; move += 1 {
		fmt.Printf("-- move %v --\n", move)
		game = game.Move()
	}

	fmt.Printf("-- final --\n")
	fmt.Println(game)
	return game.CupOrder()
}


// PartTwo returns the product of the two cups immediately clockwise of cup 1 after applying numMoves
func PartTwo(input string, numMoves int) (product int) {
	shouldAppendCupsUpToOneMillion := true
	game := NewGame(input, shouldAppendCupsUpToOneMillion)

	for move := 1; move <= numMoves; move += 1 {
		// fmt.Printf("-- move %v --\n", move)
		game = game.Move()
	}

	// fmt.Printf("-- final --\n")
	// fmt.Println(game)

	indexOfOne, err := game.Index(1)
	if err != nil {
		log.Fatal(err)
	}
	result1, result2 := game.cups[indexOfOne + 1], game.cups[indexOfOne + 2]
	return result1 * result2
}

type Game struct {
	cups []int
	currentCupIndex int
}

func NewGame(input string, shouldAppendCupsUpToOneMillion bool) (Game) {
	cups := parseCups(input)
	if shouldAppendCupsUpToOneMillion {
		cups = appendCupsUpToOneMillion(cups)
	}
	return Game{cups, 0}
}

func (g Game) Move() (newGame Game) {
	// fmt.Print(g)

	pickedUp, remainingCups := g.PickUp()

	destination := g.DestinationCup(pickedUp)
	// fmt.Printf("destination: %v\n", destination)

	destinationIndex, err := indexOf(remainingCups, destination)
	if err != nil {
		log.Fatal(err)
	}

	newCups := insertInto(remainingCups, pickedUp, destinationIndex)
	newIndex := incrementCurrentCupIndex(newCups, g.CurrentCup())

	return Game{newCups, newIndex}
}

func (g Game) PickUp() (pickedUp []int, remainingCups []int) {
	remainingCups = append([]int{}, g.cups...)

	for i := 0; i < NUM_CUPS_TO_PICK_UP; i += 1 {
		remaining, popped := pop(remainingCups, g.currentCupIndex + 1)
		remainingCups = remaining
		pickedUp = append(pickedUp, popped)
	}
	// fmt.Printf("pick up: %v\n", pickedUp)
	// fmt.Printf("remaining: %v\n", remainingCups)

	return pickedUp, remainingCups
}

func (g Game) CupOrder() (result string) {
	indexOfOne, err := indexOf(g.cups, 1)
	if err != nil {
		log.Fatal(err)
	}
	startIndex := indexOfOne + 1
	stopIndex := startIndex + len(g.cups) - 1

	for i := startIndex; i < stopIndex; i += 1 {
		cupIndex := i % len(g.cups)
		result += strconv.Itoa(g.cups[cupIndex])
	}
	return result
}

func (g Game) DestinationCup(pickedUp []int) (destinationCup int) {
	destinationCup = g.CurrentCup() - 1
	if destinationCup == 0 {
		destinationCup = g.maxCup()
	}
	for contains(pickedUp, destinationCup) {
		destinationCup = g.decrementCup(destinationCup)
	}
	return destinationCup
}

func (g Game) CurrentCup() (cup int) {
	return g.cups[g.currentCupIndex]
}

func (g Game) String() (result string) {
	result += fmt.Sprintf("cups: %v\n", g.cups)
	result += fmt.Sprintf("currentCupIndex: %v\n", g.currentCupIndex)
	return result
}

func (g Game) Index(cup int) (index int, err error) {
	return indexOf(g.cups, cup)
}

func (g Game) maxCup() (max int) {
	return getMax(g.cups)
}

func (g Game) decrementCup(cup int) (int) {
	if cup <= 1 {
		return g.maxCup()
	}
	return cup - 1
}

// Utils

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

func appendCupsUpToOneMillion(cups []int) ([]int) {
	max := getMax(cups)
	for i := max + 1; i <= 1_000_000; i += 1 {
		cups = append(cups, i)
	}
	return cups
}

func incrementCurrentCupIndex(newCups []int, currentCup int) (newIndex int) {
	currentIndex, err := indexOf(newCups, currentCup)
	if err != nil {
		log.Fatal(err)
	}

	if currentIndex == len(newCups) - 1 {
		return 0
	}
	return currentIndex + 1
}

func getMax(cups []int) (max int) {
	for _, cup := range cups {
		if cup > max {
			max = cup
		}
	}
	return max
}

func pop(list []int, index int) (remaining []int, popped int) {
	if index >= len(list) {
		return pop(list, 0)
	}
	popped = list[index]
	remaining = append([]int{}, list[:index]...)
	remaining = append(remaining, list[index + 1:]...)

	return remaining, popped
}

func contains(list []int, item int) bool {
	for _, val := range list {
		if val == item {
			return true
		}
	}
	return false
}

func insertInto(list []int, itemsToInsert []int, index int) (result []int) {
	result = append(result, list[:index+1]...)
	result = append(result, itemsToInsert...)
	result = append(result, list[index+1:]...)

	return result
}

func indexOf(list []int, item int) (index int, err error) {
	for i, v := range list {
		if item == v {
			return i, nil
		}
	}
	return 0, fmt.Errorf("item %v not found in list %v", item, list)
}
