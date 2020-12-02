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
}

// FindProductOfTwoExpensesThatSumTo first finds two expenses in the provided
// expenses array that sum to sum. Then returns their product.
func FindProductOfTwoExpensesThatSumTo(sum int, expenses []int) (product int) {
	expense1, expense2 := FindTwoExpensesThatSumTo(sum, expenses)
	product = expense1 * expense2
	log.Printf("Product of (%d, %d) is %d.", expense1, expense2, product)
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
			log.Printf("Found two expenses (%d, %d) that sum to %d.", result[0], result[1], sum)
			return result[0], result[1]
		}
		seen[expense] = true
	}

	log.Fatalf("Failed to find two expenses that sum to %d in %v", sum, expenses)
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
