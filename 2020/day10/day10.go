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

	// Part two
	result2 := GetNumberOfArrangements("input.txt")
	fmt.Printf("The number of distinct arrangements you can arrange the adapters to connect the charging outlet to your device is: %d\n", result2)
}

// GetNumberOfArrangements returns the number of distinct arrangements that can
// be used to arrange the adapters to connect the charging outlet to the device.
func GetNumberOfArrangements(filename string) int {
	adapters := readFile(filename)

	slice := addThreeHigherThanMaxJolts(addZeroJolts(adapters))
	sort.Ints(slice)

	var paths = make(map[int]int)
	paths[getMax(slice)] = 1
	return nPaths(paths, slice, 0)
}

func nPaths(paths map[int]int, slice []int, start int) int {
	fmt.Printf("Start %v Paths %#v\n", start, paths)
	if paths[start] == 0 {
		if contains(slice, start) {
			paths[start] = nPaths(paths, slice, start+1) + nPaths(paths, slice, start+2) + nPaths(paths, slice, start+3)
		} else {
			paths[start] = 0
		}
	}
	return paths[start]
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

func contains(slice []int, target int) bool {
	for _, current := range slice {
		if current == target {
			return true
		}
	}
	return false
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
