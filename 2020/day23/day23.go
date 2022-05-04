package main

import (
	"fmt"
	"log"
	"strconv"
)

const NUM_CUPS_TO_PICK_UP int = 3

type Game struct {
	cups []int
	currentCupIndex int
}

func NewGame(input string) (Game) {
	return Game{parseCups(input), 0}
}

func (g Game) Move() (newGame Game) {
	fmt.Print(g)

	pickedUp, remainingCups := g.PickUp()

	destination := g.DestinationCup(pickedUp)
	fmt.Printf("destination: %v\n", destination)

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
	fmt.Printf("pick up: %v\n", pickedUp)
	fmt.Printf("remaining: %v\n", remainingCups)

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

func (g Game) decrementCup(cup int) (int) {
	if cup <= 1 {
		return g.maxCup()
	}
	return cup - 1
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

func (g Game) maxCup() (max int) {
	for _, cup := range g.cups {
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

// PartOne retruns the order of the cups after applying numMoves
func PartOne(input string, numMoves int) (result string) {
	game := NewGame(input)

	for move := 1; move <= numMoves; move += 1 {
		fmt.Printf("-- move %v --\n", move)
		game = game.Move()
	}

	fmt.Printf("-- final --\n")
	fmt.Println(game)
	return game.CupOrder()
}

func indexOf(list []int, item int) (index int, err error) {
	for i, v := range list {
		if item == v {
			return i, nil
		}
	}
	return 0, fmt.Errorf("item %v not found in list %v", item, list)
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
