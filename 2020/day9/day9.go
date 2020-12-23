package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting day9")

	// Part one
	preambleLength := 5
	result1 := GetFirstNumberThatIsNotSumOfPair("input_test.txt", preambleLength)
	fmt.Printf("The first number that is not the sum of a pair in the previous %d numbers is: %d\n", preambleLength, result1)
}

func GetFirstNumberThatIsNotSumOfPair(filename string, preambleLength int) int {
	numbers := readFile(filename)
	fmt.Printf("numbers %v\n", numbers)
	return 0
}

func readFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	numbers := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return numbers
}
