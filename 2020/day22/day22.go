package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Deck struct {
	name  string
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

func (d Deck) Push(card int) Deck {
	return Deck{
		d.name,
		append(d.cards, card),
	}
}

func (d Deck) Shift() (newDeck Deck, result int) {
	newDeck = Deck{
		d.name,
		d.cards[1:],
	}
	result = d.cards[0]

	return newDeck, result
}

func (d Deck) Copy(numCards int) (newDeck Deck) {
	newCards := make([]int, numCards)
	copy(newCards, d.cards)

	newDeck = Deck{
		d.name,
		newCards,
	}

	return newDeck
}

func (d Deck) Len() int {
	return len(d.cards)
}

func (d Deck) Score() (score int) {
	for i, card := range reverse(d.cards) {
		score += card * (i + 1)
	}
	return score
}

func main() {
	fmt.Printf("Starting day22...\n")

	partOne := PartOne("input.txt")
	fmt.Printf("PartOne: %v\n", partOne)

	partTwo := PartTwo("input.txt")
	fmt.Printf("PartTwo: %v\n", partTwo)
}

func PartOne(filename string) (score int) {
	lines := readLines(filename)
	playerOneLines, playerTwoLines := splitLines(lines)

	deckOne := NewDeck("Player 1", playerOneLines)
	deckTwo := NewDeck("Player 2", playerTwoLines)

	roundNumber := 1
	for !isGameOver(deckOne, deckTwo) {
		deckOne, deckTwo, roundNumber = playRound(deckOne, deckTwo, roundNumber)
	}

	fmt.Println("== Post-game results ==")
	fmt.Println(deckOne)
	fmt.Println(deckTwo)

	winningDeck, err := getWinningDeck(deckOne, deckTwo)
	if err != nil {
		log.Fatal(err)
	}
	return winningDeck.Score()
}

func PartTwo(filename string) (score int) {
	lines := readLines(filename)
	playerOneLines, playerTwoLines := splitLines(lines)

	deckOne := NewDeck("Player 1", playerOneLines)
	deckTwo := NewDeck("Player 2", playerTwoLines)

	newDeckOne, newDeckTwo, _ := playSubGame(deckOne, deckTwo, 1)

	fmt.Println("== Post-game results ==")
	fmt.Println(newDeckOne)
	fmt.Println(newDeckTwo)

	winningDeck, err := getWinningDeck(newDeckOne, newDeckTwo)
	if err != nil {
		log.Fatal(err)
	}
	return winningDeck.Score()
}

// playRound plays a round according to the rules of PartOne
func playRound(deckOne Deck, deckTwo Deck, roundNumber int) (newDeckOne Deck, newDeckTwo Deck, newRoundNumber int) {
	fmt.Printf("-- Round %d --\n", roundNumber)
	fmt.Println(deckOne)
	fmt.Println(deckTwo)

	newDeckOne, playerOneCard := deckOne.Shift()
	newDeckTwo, playerTwoCard := deckTwo.Shift()

	fmt.Printf("Player 1 plays: %d\n", playerOneCard)
	fmt.Printf("Player 2 plays: %d\n", playerTwoCard)

	if playerOneCard > playerTwoCard {
		fmt.Printf("Player 1 wins the round!\n")
		newDeckOne = newDeckOne.Push(playerOneCard)
		newDeckOne = newDeckOne.Push(playerTwoCard)
	} else {
		fmt.Printf("Player 2 wins the round!\n")
		newDeckTwo = newDeckTwo.Push(playerTwoCard)
		newDeckTwo = newDeckTwo.Push(playerOneCard)
	}
	fmt.Println()

	return newDeckOne, newDeckTwo, roundNumber + 1
}

// playSubGame plays a game according to the rules of PartTwo
func playSubGame(deckOne Deck, deckTwo Deck, gameNumber int) (newDeckOne Deck, newDeckTwo Deck, winner string) {
	fmt.Printf("=== Game %d ===\n\n", gameNumber)

	// seen is a map from serialized game state to true (if seen)
	seen := map[string]bool{}

	roundNumber := 1
	for !isGameOver(deckOne, deckTwo) {
		if hasSeenBefore(seen, deckOne, deckTwo) {
			return deckOne, deckTwo, deckOne.name
		}
		serialized := serialize(deckOne, deckTwo)
		seen[serialized] = true

		fmt.Printf("=== Round %d (Game %d) --\n", roundNumber, gameNumber)
		fmt.Println(deckOne)
		fmt.Println(deckTwo)

		newDeckOne, playerOneCard := deckOne.Shift()
		newDeckTwo, playerTwoCard := deckTwo.Shift()

		fmt.Printf("Player 1 plays: %d\n", playerOneCard)
		fmt.Printf("Player 2 plays: %d\n", playerTwoCard)
		if playerOneCard <= len(newDeckOne.cards) && playerTwoCard <= len(newDeckTwo.cards) {
			fmt.Printf("Playing a sub-game to determine the winner...\n")
			copiedDeckOne := newDeckOne.Copy(playerOneCard)
			copiedDeckTwo := newDeckTwo.Copy(playerTwoCard)
			_, _, winner = playSubGame(copiedDeckOne, copiedDeckTwo, gameNumber+1)
		} else if playerOneCard > playerTwoCard {
			winner = deckOne.name
		} else {
			winner = deckTwo.name
		}

		if winner == deckOne.name {
			fmt.Printf("Player 1 wins the round!\n")
			newDeckOne = newDeckOne.Push(playerOneCard)
			newDeckOne = newDeckOne.Push(playerTwoCard)
		} else if winner == deckTwo.name {
			fmt.Printf("Player 2 wins the round!\n")
			newDeckTwo = newDeckTwo.Push(playerTwoCard)
			newDeckTwo = newDeckTwo.Push(playerOneCard)
		} else {
			log.Fatal("unexpected winner")
		}

		deckOne = newDeckOne
		deckTwo = newDeckTwo
		roundNumber += 1
	}

	fmt.Printf("...anyway, back to game %d.\n", gameNumber-1)
	return deckOne, deckTwo, winner
}

func hasSeenBefore(seen map[string]bool, deckOne Deck, deckTwo Deck) bool {
	serialized := serialize(deckOne, deckTwo)
	return seen[serialized]
}

func serialize(deckOne Deck, deckTwo Deck) (serialized string) {
	serializedOne := deckOne.String()
	serializedTwo := deckTwo.String()

	return strings.Join([]string{serializedOne, serializedTwo}, "|")
}

func reverse(list []int) (reversed []int) {
	for _, element := range list {
		reversed = append([]int{element}, reversed...)
	}
	return reversed
}

func isGameOver(deckOne Deck, deckTwo Deck) bool {
	return deckOne.Len() == 0 || deckTwo.Len() == 0
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

func splitLines(lines []string) (playerOneLines []string, playerTwoLines []string) {
	playerOneLines = lines[:len(lines)/2]
	playerTwoLines = lines[len(lines)/2+1:]

	return playerOneLines, playerTwoLines
}

func getWinningDeck(deckOne Deck, deckTwo Deck) (winner Deck, err error) {
	if deckOne.Len() > 0 {
		return deckOne, nil
	} else if deckTwo.Len() > 0 {
		return deckTwo, nil
	} else {
		return Deck{}, errors.New("neither deck has len > 0")
	}
}
