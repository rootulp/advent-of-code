package main

import "fmt"

func main() {
	fmt.Println("Starting day 15")

	// Part one
	input := []int{18, 11, 9, 0, 5, 1}
	partOne := MemoryGame(input, 2020)
	fmt.Printf("Part one %v\n", partOne)

	// Part two
	// example := []int{0, 3, 6}
	// partTwo := MemoryGameOptimized(example, 10)
	partTwo := MemoryGameOptimized(input, 30000000)
	fmt.Printf("Part two %v\n", partTwo)
}

func MemoryGame(startingNumbers []int, turns int) int {
	memory := map[int]int{}
	turn := 0
	for _, num := range startingNumbers {
		fmt.Printf("Turn %v, spoke %v because starting number\n", turn, num)
		memory[turn] = num
		turn += 1
	}
	for turn < turns {
		lastNumberSpoken := memory[turn-1]
		mostRecentTurn := getMostRecentTurn(lastNumberSpoken, turn, memory)
		if mostRecentTurn == -1 {
			fmt.Printf("Turn %v, spoke %v because last number %v hasn't been seen\n", turn, 0, lastNumberSpoken)
			memory[turn] = 0
		} else {
			spoke := (turn - 1) - mostRecentTurn
			fmt.Printf("Turn %v, spoke %v because last number %v was spoken on %v and %v\n", turn, spoke, lastNumberSpoken, turn-1, mostRecentTurn)
			memory[turn] = spoke
		}
		turn += 1
	}

	return memory[turns-1]
}

func getMostRecentTurn(lastNumberSpoken int, turn int, memory map[int]int) int {
	for i := turn - 2; i >= 0; i-- {
		if memory[i] == lastNumberSpoken {
			return i
		}
	}
	return -1
}

func MemoryGameOptimized(startingNumbers []int, turns int) int {
	turnToNumberSpoken := map[int]int{}
	numberToMostRecentTurn := map[int]int{}
	turn := 0
	for _, num := range startingNumbers {
		fmt.Printf("Turn %v, spoke %v because starting number\n", turn, num)
		turnToNumberSpoken[turn] = num

		lastTurn := turn - 1
		lastNumberSpoken := turnToNumberSpoken[lastTurn]
		numberToMostRecentTurn[lastNumberSpoken] = lastTurn
		turn += 1
	}
	for turn < turns {
		lastTurn := turn - 1
		lastNumberSpoken := turnToNumberSpoken[lastTurn]
		if mostRecentTurn, ok := numberToMostRecentTurn[lastNumberSpoken]; ok {
			spoke := lastTurn - mostRecentTurn
			fmt.Printf("Turn %v, spoke %v because last number %v was spoken on %v and %v\n", turn, spoke, lastNumberSpoken, lastTurn, mostRecentTurn)
			turnToNumberSpoken[turn] = spoke
		} else {
			spoke := 0
			fmt.Printf("Turn %v, spoke %v because last number %v hasn't been seen\n", turn, spoke, lastNumberSpoken)
			turnToNumberSpoken[turn] = spoke
		}
		numberToMostRecentTurn[lastNumberSpoken] = lastTurn
		turn += 1
	}
	return turnToNumberSpoken[turns-1]
}
