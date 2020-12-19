package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	responses := ReadFile("input.txt")

	sum := 0
	for _, response := range responses {
		unique := GetNumUnique(response)
		sum += unique
	}
	fmt.Printf("Sum of unique counts %v\n", sum)
}

// GetNumUnique returns the number of unique characters in the provided response
func GetNumUnique(response string) int {
	seen := make(map[rune]bool)
	for _, char := range response {
		seen[char] = true
	}
	return len(seen)
}

// ReadFile returns a list responses per group
func ReadFile(filename string) (responses []string) {
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
