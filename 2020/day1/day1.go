package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	sum := 2020
	expenses := readFile("./input.txt")

	FindProductOfTwoExpensesThatSumTo(sum, expenses)
	FindProductOfThreeExpensesThatSumTo(sum, expenses)
}

// FindProductOfTwoExpensesThatSumTo first finds two expenses in the provided
// expenses array that sum to sum. Then returns their product.
func FindProductOfTwoExpensesThatSumTo(sum int, expenses []int) (product int) {
	expense1, expense2 := FindTwoExpensesThatSumTo(sum, expenses)
	product = expense1 * expense2
	log.Printf("%d * %d = %d", expense1, expense2, product)
	return product
}

// FindProductOfThreeExpensesThatSumTo first finds two expenses in the provided
// expenses array that sum to sum. Then returns their product.
func FindProductOfThreeExpensesThatSumTo(sum int, expenses []int) (product int) {
	expense1, expense2, expense3 := FindThreeExpensesThatSumTo(sum, expenses)
	product = expense1 * expense2 * expense3
	log.Printf("%d * %d * %d = %d", expense1, expense2, expense3, product)
	return product
}

// FindTwoExpensesThatSumTo finds two expenses in the provided array that sum to
// sum.
func FindTwoExpensesThatSumTo(sum int, expenses []int) (expense1, expense2 int) {
	seen := make(map[int]bool)
	for _, expense := range expenses {
		complement := sum - expense
		// log.Printf("Expense %v, complement %v, seen %v", expense, complement, seen)
		if seen[complement] {
			result := []int{expense, complement}
			sort.Ints(result)

			// TODO how to unpack result into expense1 and expense2?
			// expense1, expense2 := result...
			log.Printf("Found two expenses that sum to %d = %d + %d", sum, result[0], result[1])
			return result[0], result[1]
		}
		seen[expense] = true
	}

	log.Fatalf("Failed to find two expenses that sum to %d in %v", sum, expenses)
	return // unreached
}

// FindThreeExpensesThatSumTo finds three expenses in the provided array that sum to
// sum.
func FindThreeExpensesThatSumTo(sum int, expenses []int) (expense1, expense2, expense3 int) {

	pairsSeen := make(map[int][]int)
	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			sum := expenses[i] + expenses[j]
			pair := []int{expenses[i], expenses[j]}
			pairsSeen[sum] = pair
		}
	}

	for _, expense := range expenses {
		complement := sum - expense
		if val, ok := pairsSeen[complement]; ok {
			// result := make([]int, 3)
			// copy(result, val)
			// result[2] = expense
			var result []int
			result = append(result, expense)
			result = append(result, val...)

			// Sort result so we get a reliable output and can write a unit test against it
			sort.Ints(result)

			// TODO how to unpack result into expense1 and expense2?
			// expense1, expense2 := result...
			log.Printf("Found three expenses that sum to %d = %d + %d + %d", sum, result[0], result[1], result[2])
			return result[0], result[1], result[2]
		}
	}

	log.Fatalf("Failed to find three expenses that sum to %d in %v", sum, expenses)
	return // unreached
}

func readFile(filename string) (expenses []int) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" { // Reached end of file
			file.Close()
			return expenses
		}

		expense, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
		expenses = append(expenses, expense)
	}

	file.Close()
	return expenses
}
