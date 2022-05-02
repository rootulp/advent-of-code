package main

import "fmt"

func main() {
	fmt.Printf("Starting day 21...\n")

	partOne := PartOne("example.txt")
	fmt.Printf("Part One: %v", partOne)
}

// PartOne returns the number of times a "safe ingredient" appears in the list of recipes in filename
// A "safe ingredient" is one that can not possibly contain any of the allergens
func PartOne(filename string) int {
	return 0
}
