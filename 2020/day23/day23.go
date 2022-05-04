package main

import (
	"fmt"
	"log"
	"strconv"
)

type Game struct {
	cups []int
	currentCupIndex int
}

func NewGame(input string) (Game) {
	return Game{parseCups(input), 0}
}

func (g Game) Move() (newGame Game) {
	// TODO
	return g
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
