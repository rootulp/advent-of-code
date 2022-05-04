package main

import (
	"bufio"
	"encoding/json"
	"errors"
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

	winner, err := getWinner(deckOne, deckTwo)
	if err != nil {
		log.Fatal(err)
	}
	return winningScore(winner)
}

func PartTwo(filename string) (score int) {
	lines := readLines(filename)
	playerOneLines, playerTwoLines := splitLines(lines)

	deckOne := NewDeck("Player 1", playerOneLines)
	deckTwo := NewDeck("Player 2", playerTwoLines)

	gameNumber := 1
	fmt.Printf("=== Game %d ===\n\n", gameNumber)

	// seen is a map from serialized game state to true (if seen)
	seen := map[string]bool{}

	roundNumber := 1
	for !isGameOver(deckOne, deckTwo) {
		if hasSeenBefore(seen, deckOne, deckTwo) {
			return winningScore(deckOne)
		}
		serialized := serialize(deckOne, deckTwo)
		seen[serialized] = true

		fmt.Printf("-- Round %d (Game %d) --\n", roundNumber, gameNumber)
		fmt.Println(deckOne)
		fmt.Println(deckTwo)

		newDeckOne, playerOneCard := deckOne.Shift()
		newDeckTwo, playerTwoCard := deckTwo.Shift()

		fmt.Printf("Player 1 plays: %d\n", playerOneCard)
		fmt.Printf("Player 2 plays: %d\n", playerTwoCard)
		shouldRecurse := playerOneCard <= len(newDeckOne.cards) && playerTwoCard <= len(newDeckTwo.cards)
		if shouldRecurse{
			fmt.Printf("Playing a sub-game to determine the winner...\n")
			copiedDeckOne := newDeckOne.Copy(playerOneCard)
			copiedDeckTwo := newDeckTwo.Copy(playerTwoCard)
			winner := playSubGame(copiedDeckOne, copiedDeckTwo, gameNumber + 1)
			if winner == deckOne.name {
				fmt.Printf("Player 1 wins the round!\n")
				newDeckOne = newDeckOne.Push(playerOneCard)
				newDeckOne = newDeckOne.Push(playerTwoCard)
			} else if winner == deckTwo.name {
				fmt.Printf("Player 2 wins the round!\n")
				newDeckTwo = newDeckTwo.Push(playerTwoCard)
				newDeckTwo = newDeckTwo.Push(playerOneCard)
			}
		} else if playerOneCard > playerTwoCard {
			fmt.Printf("Player 1 wins the round!\n")
			newDeckOne = newDeckOne.Push(playerOneCard)
			newDeckOne = newDeckOne.Push(playerTwoCard)
		} else {
			fmt.Printf("Player 2 wins the round!\n")
			newDeckTwo = newDeckTwo.Push(playerTwoCard)
			newDeckTwo = newDeckTwo.Push(playerOneCard)
		}
		deckOne = newDeckOne
		deckTwo = newDeckTwo
		roundNumber += 1
	}

	fmt.Println("== Post-game results ==")
	fmt.Println(deckOne)
	fmt.Println(deckTwo)

	winner, err := getWinner(deckOne, deckTwo)
	if err != nil {
		log.Fatal(err)
	}
	return winningScore(winner)
}

func playSubGame(deckOne Deck, deckTwo Deck, gameNumber int) (winner string) {
	fmt.Printf("=== Game %d ===\n\n", gameNumber)

	// seen is a map from serialized game state to true (if seen)
	seen := map[string]bool{}

	roundNumber := 1
	for !isGameOver(deckOne, deckTwo) {
		if hasSeenBefore(seen, deckOne, deckTwo) {
			fmt.Printf("hasSeenBefore triggered\n")
			return deckOne.name
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
			winner := playSubGame(copiedDeckOne, copiedDeckTwo, gameNumber + 1)
			if winner == deckOne.name {
				fmt.Printf("Player 1 wins the round!\n")
				newDeckOne = newDeckOne.Push(playerOneCard)
				newDeckOne = newDeckOne.Push(playerTwoCard)
			} else if winner == deckTwo.name {
				fmt.Printf("Player 2 wins the round!\n")
				newDeckTwo = newDeckTwo.Push(playerTwoCard)
				newDeckTwo = newDeckTwo.Push(playerOneCard)
			}
		} else if playerOneCard > playerTwoCard {
			fmt.Printf("Player 1 wins the round!\n")
			newDeckOne = newDeckOne.Push(playerOneCard)
			newDeckOne = newDeckOne.Push(playerTwoCard)
			winner = deckOne.name
		} else {
			fmt.Printf("Player 2 wins the round!\n")
			newDeckTwo = newDeckTwo.Push(playerTwoCard)
			newDeckTwo = newDeckTwo.Push(playerOneCard)
			winner = deckTwo.name
		}
		deckOne = newDeckOne
		deckTwo = newDeckTwo
		roundNumber += 1
	}

	fmt.Printf("...anyway, back to game %d.\n", gameNumber - 1)
	return winner
}

func hasSeenBefore(seen map[string]bool, deckOne Deck, deckTwo Deck) bool {
	serialized := serialize(deckOne, deckTwo)
	result := seen[serialized]
	if result {
		fmt.Printf("serizlied %v\n", serialized)
		// PrettyPrint(seen)
	}
	return result
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
			fmt.Println(string(b))
	}
	return
}

func serialize(deckOne Deck, deckTwo Deck) (serialized string) {
	serializedOne := deckOne.String()
	serializedTwo := deckTwo.String()

	return strings.Join([]string{serializedOne, serializedTwo}, "|")
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
	return deckOne.Len() == 0 || deckTwo.Len() == 0
}

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
	playerOneLines = lines[:len(lines) / 2]
	playerTwoLines = lines[len(lines) / 2 + 1:]

	return playerOneLines, playerTwoLines
}

func getWinner(deckOne Deck, deckTwo Deck) (winner Deck, err error) {
	if deckOne.Len() > 0 {
		return deckOne, nil
	} else if deckTwo.Len() > 0 {
		return deckTwo, nil
	} else {
		return Deck{}, errors.New("neither deck has len > 0")
	}
}
