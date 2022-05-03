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

func (d Deck) push(card int) Deck {
	return Deck{
		d.name,
		append(d.cards, card),
	}
}

func (d Deck) shift() (newDeck Deck, result int) {
	newDeck = Deck{
		d.name,
		d.cards[1:],
	}
	result = d.cards[0]

	return newDeck, result
}

func main() {
	fmt.Printf("Starting day22...\n")

	partOne := PartOne("input.txt")
	fmt.Printf("PartOne: %v\n", partOne)

	partTwo := PartTwo("example.txt")
	fmt.Printf("PartTwo: %v\n", partTwo)
}

func PartOne(filename string) (score int) {
	lines := readLines(filename)

	playerOneLines := lines[:len(lines) / 2]
	playerTwoLines := lines[len(lines) / 2 + 1:]

	deckOne := NewDeck("Player 1", playerOneLines)
	deckTwo := NewDeck("Player 2", playerTwoLines)

	roundNumber := 1
	for !isGameOver(deckOne, deckTwo) {
		deckOne, deckTwo, roundNumber = playRound(deckOne, deckTwo, roundNumber)
	}

	fmt.Println("== Post-game results ==")
	fmt.Println(deckOne)
	fmt.Println(deckTwo)

	if len(deckOne.cards) > 0 {
		return winningScore(deckOne)
	} else if len(deckTwo.cards) > 0 {
		return winningScore(deckTwo)
	} else {
		log.Fatalf("neither deck has len > 0")
	}
	return 0
}

func PartTwo(filename string) (score int) {
	lines := readLines(filename)

	playerOneLines := lines[:len(lines) / 2]
	playerTwoLines := lines[len(lines) / 2 + 1:]

	deckOne := NewDeck("Player 1", playerOneLines)
	deckTwo := NewDeck("Player 2", playerTwoLines)

	gameNumber := 1
	var winner Deck
	for !isGameOver(deckOne, deckTwo) {
		deckOne, deckTwo, winner = playGame(deckOne, deckTwo, gameNumber)
	}

	fmt.Println("== Post-game results ==")
	fmt.Println(deckOne)
	fmt.Println(deckTwo)
	return winningScore(winner)
}

func winningScore(deck Deck) (score int) {
	for i, card := range reverse(deck.cards) {
		score += card * (i + 1)
	}
	return score
}

func reverse(list []int) (reversed []int) {
	for _, element := range list {
		reversed = append([]int{element}, reversed...)
	}
	return reversed
}

func isGameOver(deckOne Deck, deckTwo Deck) bool {
	return len(deckOne.cards) == 0 || len(deckTwo.cards) == 0
}

func playRound(deckOne Deck, deckTwo Deck, roundNumber int) (newDeckOne Deck, newDeckTwo Deck, newRoundNumber int) {
	fmt.Printf("-- Round %d --\n", roundNumber)
	fmt.Println(deckOne)
	fmt.Println(deckTwo)

	newDeckOne, playerOneCard := deckOne.shift()
	newDeckTwo, playerTwoCard := deckTwo.shift()

	fmt.Printf("Player 1 plays: %d\n", playerOneCard)
	fmt.Printf("Player 2 plays: %d\n", playerTwoCard)

	if playerOneCard > playerTwoCard {
		fmt.Printf("Player 1 wins the round!\n")
		newDeckOne = newDeckOne.push(playerOneCard)
		newDeckOne = newDeckOne.push(playerTwoCard)
	} else {
		fmt.Printf("Player 2 wins the round!\n")
		newDeckTwo = newDeckTwo.push(playerTwoCard)
		newDeckTwo = newDeckTwo.push(playerOneCard)
	}
	fmt.Println()

	return newDeckOne, newDeckTwo, roundNumber + 1
}

func playGame(deckOne Deck, deckTwo Deck, gameNumber int) (deckOne Deck, deckTwo Deck, winner Deck) {
	fmt.Printf("=== Game %d ===\n\n", gameNumber)

	roundNumber := 1
	for !isGameOver(deckOne, deckTwo, subgameWinner) {
		deckOne, deckTwo, roundNumber = playRound(deckOne, deckTwo, roundNumber)
	}

}

func playRecursiveRound(deckOne Deck, deckTwo Deck, roundNumber int, gameNumber int) (newDeckOne Deck, newDeckTwo Deck, newRoundNumber int, newGameNumber int) {
	fmt.Printf("-- Round %d -- (Game %d)\n", roundNumber, gameNumber)
	fmt.Println(deckOne)
	fmt.Println(deckTwo)

	newDeckOne, playerOneCard := deckOne.shift()
	newDeckTwo, playerTwoCard := deckTwo.shift()

	fmt.Printf("Player 1 plays: %d\n", playerOneCard)
	fmt.Printf("Player 2 plays: %d\n", playerTwoCard)

	if shouldRecurse(playerOneCard, len(newDeckOne.cards)) && shouldRecurse(playerTwoCard, len(newDeckTwo.cards)) {
		fmt.Println("Playing a sub-game to determine the winner...")


	} else if playerOneCard > playerTwoCard {
		fmt.Printf("Player 1 wins round %d of game %d!\n", roundNumber, gameNumber)
		newDeckOne = newDeckOne.push(playerOneCard)
		newDeckOne = newDeckOne.push(playerTwoCard)
	} else {
		fmt.Printf("Player 2 wins the round!\n")
		newDeckTwo = newDeckTwo.push(playerTwoCard)
		newDeckTwo = newDeckTwo.push(playerOneCard)
	}
	fmt.Println()

	return newDeckOne, newDeckTwo, roundNumber + 1, gameNumber
}

func shouldRecurse(currentCard int, numRemainingCards int) bool {
	return currentCard < numRemainingCards
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
