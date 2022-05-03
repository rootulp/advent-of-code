package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Deck struct {
	name string
	cards []int
}

func NewDeck(name string, lines []string) Deck {
	cards := []int{}
	for _, line := range lines {
		if strings.HasPrefix(line, "Player") {
			continue
		}
		card, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
	}
	return Deck{name, cards}
}

func (d Deck) String() string {
	return fmt.Sprintf("%s: %v", d.name, d.cards)
}

func main() {
	fmt.Printf("Starting day22...\n")

	partOne := PartOne("example.txt")
	fmt.Printf("PartOne: %v\n", partOne)

	// partTwo := PartTwo("example.txt")
	// fmt.Printf("PartTwo: %v\n", partTwo)
}

func PartOne(filename string) (score int) {
	lines := readLines(filename)

	playerOneLines := lines[:len(lines) / 2]
	playerTwoLines := lines[len(lines) / 2 + 1:]

	deckOne := NewDeck("Player 1", playerOneLines)
	deckTwo := NewDeck("Player 2", playerTwoLines)

	fmt.Println(deckOne)
	fmt.Println(deckTwo)
	return 0
}

func readLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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
