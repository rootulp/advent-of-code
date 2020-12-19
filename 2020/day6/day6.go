package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	groups := ReadFileIntoGroups("input.txt")

	// Part one
	PrintSumOfUnique(groups)

	// Part two
	PrintSumOfCommon(groups)
}

// PrintSumOfUnique prints the sum of unique characters in each response
func PrintSumOfUnique(groups [][]string) {
	sum := 0
	for _, group := range groups {
		unique := GetNumUnique(group)
		sum += unique
	}
	fmt.Printf("Sum of unique counts %v\n", sum)
}

// PrintSumOfCommon prints the sum of common responses in each group
func PrintSumOfCommon(groups [][]string) {
	sum := 0
	for _, response := range groups {
		unique := GetNumCommon(response)
		sum += unique
	}
	fmt.Printf("Sum of common counts %v\n", sum)
}

// GetNumCommon returns the number of responses that the group has in common
func GetNumCommon(group []string) (numCommon int) {
	countOfResponses := make(map[rune]int)
	for _, response := range group {
		for _, answer := range response {
			countOfResponses[answer]++
		}
	}
	numCommon = 0
	for response := range countOfResponses {
		if countOfResponses[response] == len(group) {
			numCommon++
		}
	}
	return numCommon
}

// GetNumUnique returns the number of unique characters in the provided response
func GetNumUnique(group []string) int {
	seen := make(map[rune]bool)
	for _, response := range group {
		for _, char := range response {
			seen[char] = true
		}
	}
	return len(seen)
}

// ReadFileIntoGroups returns a list of groups. Each group is a list of
// responses (one per person).
func ReadFileIntoGroups(filename string) (groupResponses [][]string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	groupResponses = [][]string{}
	response := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			groupResponses = append(groupResponses, response)
			response = []string{}
		} else {
			response = append(response, text)
		}
	}
	groupResponses = append(groupResponses, response)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
