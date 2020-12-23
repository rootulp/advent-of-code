package main

import "fmt"

func main() {
	fmt.Println("Starting day10")

	// Part one
	result1 := GetProductOfOneJoltDifferencesAndThreeJoltDifferences("input.txt")
	fmt.Printf("The product of 1-jolt differences and 3-jolt differences is: %d\n", result1)
}

// GetProductOfOneJoltDifferencesAndThreeJoltDifferences returns the number of
// 1-jolt differences * the number of 3-jolt differences.
func GetProductOfOneJoltDifferencesAndThreeJoltDifferences(filename string) int {
	return 0
}
