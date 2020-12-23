package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Part one
	preambleLength := 25
	result1 := GetFirstNumberThatIsNotSumOfPair("input.txt", preambleLength)
	fmt.Printf("The first number that is not the sum of a pair in the previous %d numbers is: %d\n", preambleLength, result1)

	// Part two
	result2 := GetSumOfSmallestAndLargestInContiguousRange("input.txt", preambleLength)
	fmt.Printf("The sum of the smallest and largest numbers in the contiguous range that sums to invalid number %d is: %d\n", result1, result2)
}

// GetSumOfSmallestAndLargestInContiguousRange first finds the invalid number
// from part one. Then it finds a contiguous range that sums to this invalid
// number. This function returns the sum of the smallest and largest numbers in
// this contiguous range.
func GetSumOfSmallestAndLargestInContiguousRange(filename string, preambleLength int) int {
	invalidNumber := GetFirstNumberThatIsNotSumOfPair(filename, preambleLength)
	numbers := readFile(filename)

	for rangeLength := 2; rangeLength < len(numbers); rangeLength++ {
		for i := 0; i < len(numbers)-rangeLength; i++ {
			targetRange := numbers[i : i+rangeLength]
			if isNumberTheSumOfRange(targetRange, invalidNumber) {
				return min(targetRange) + max(targetRange)
			}
		}
	}

	log.Fatalf("failed to find a contiguous range that sums to %d", invalidNumber)
	return 0 // unreached
}

// GetFirstNumberThatIsNotSumOfPair returns the first number that isn't the sum
// of a pair of numbers in the prior preambleLength set of numbers.
func GetFirstNumberThatIsNotSumOfPair(filename string, preambleLength int) int {
	numbers := readFile(filename)
	// fmt.Printf("numbers %v\n", numbers)
	for i := preambleLength; i < len(numbers); i++ {
		if !isNumberTheSumOfPair(numbers[i-preambleLength:i], preambleLength, numbers[i]) {
			return numbers[i]
		}
	}

	log.Fatalf("failed to find a number that is not the sum of a pair")
	return 0 // unreached
}

func isNumberTheSumOfPair(rangeToSearch []int, preambleLength int, target int) bool {
	// log.Printf("isNumberInTheSumOfPair %v, %v, %v", rangeToSearch, preambleLength, target)
	for j := 0; j < len(rangeToSearch); j++ {
		for k := j + 1; k < len(rangeToSearch); k++ {
			if rangeToSearch[j]+rangeToSearch[k] == target {
				return true
			}
		}
	}
	return false
}

func isNumberTheSumOfRange(rangeToSearch []int, target int) bool {
	sum := 0
	for _, num := range rangeToSearch {
		sum += num
	}
	return sum == target
}

func min(slice []int) int {
	min := slice[0]
	for _, num := range slice {
		if num < min {
			min = num
		}
	}
	return min
}

func max(slice []int) int {
	max := slice[0]
	for _, num := range slice {
		if num > max {
			max = num
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
