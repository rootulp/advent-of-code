package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Part one
	responses := ReadFileIntoResponses("input.txt")
	PrintSumOfUnique(responses)

	// Part two
	groups := ReadFileIntoGroups("input.txt")
	PrintSumOfCommon(groups)
}

// PrintSumOfUnique prints the sum of unique characters in each response
func PrintSumOfUnique(responses []string) {
	sum := 0
	for _, response := range responses {
		unique := GetNumUnique(response)
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
func GetNumUnique(response string) int {
	seen := make(map[rune]bool)
	for _, char := range response {
		seen[char] = true
	}
	return len(seen)
}

// ReadFileIntoResponses returns a list responses per group. Each response represents one group.
func ReadFileIntoResponses(filename string) (responses []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	responses = []string{}
	response := ""
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			responses = append(responses, response)
			response = ""
		} else {
			response += text
		}
	}
	responses = append(responses, response)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
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
