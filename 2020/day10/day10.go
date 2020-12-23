package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Starting day10")

	// Part one
	result1 := GetProductOfOneJoltDifferencesAndThreeJoltDifferences("input.txt")
	fmt.Printf("The product of 1-jolt differences and 3-jolt differences is: %d\n", result1)
}

// GetProductOfOneJoltDifferencesAndThreeJoltDifferences returns the number of
// 1-jolt differences * the number of 3-jolt differences.
func GetProductOfOneJoltDifferencesAndThreeJoltDifferences(filename string) int {
	adapters := readFile(filename)
	// log.Printf("adapters %v\n", adapters)

	slice := addThreeHigherThanMaxJolts(addZeroJolts(adapters))
	sort.Ints(slice)

	// log.Printf("slice %v\n", slice)
	return getNumberOfOneJoltDifferences(slice) * getNumberOfThreeJoltDifferences(slice)
}

func getNumberOfOneJoltDifferences(slice []int) int {
	return getNumberOfNJoltDifferences(slice, 1)
}

func getNumberOfThreeJoltDifferences(slice []int) int {
	return getNumberOfNJoltDifferences(slice, 3)
}

func getNumberOfNJoltDifferences(slice []int, n int) int {
	differences := 0
	for i := 1; i < len(slice); i++ {
		if slice[i]-slice[i-1] == n {
			differences++
		}
	}
	return differences
}

func addZeroJolts(adapters []int) []int {
	return append(adapters, 0)
}

func addThreeHigherThanMaxJolts(adapters []int) []int {
	max := getMax(adapters)
	return append(adapters, max+3)
}

func getMax(slice []int) int {
	max := slice[0]
	for _, current := range slice {
		if current > max {
			max = current
		}
	}
	return max
}

func readFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	adapters := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		adapter, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		adapters = append(adapters, adapter)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return adapters
}
