package main

import "testing"

func TestFindTwoExpensesThatSumTo(t *testing.T) {
	expenses := []int{1721, 979, 366, 299, 675, 1456}
	sum := 2020
	expected1, expected2 := 299, 1721

	result1, result2 := FindTwoExpensesThatSumTo(sum, expenses)
	if result1 != expected1 || result2 != expected2 {
		t.Errorf("FindTwoExpensesThatSumTo was incorrect. Receieved: %d, %d. Wanted: %d, %d.", result1, result2, expected1, expected2)
	}
}

func TestFindProductOfTwoExpensesThatSumTo(t *testing.T) {
	expenses := []int{1721, 979, 366, 299, 675, 1456}
	sum := 2020
	expected := 514579

	result := FindProductOfTwoExpensesThatSumTo(sum, expenses)
	if result != expected {
		t.Errorf("FindProductOfTwoExpensesThatSumTo was incorrect. Received: %d. Wanted: %d.", result, expected)
	}
}

func TestFindThreeExpensesThatSumTo(t *testing.T) {
	expenses := []int{1721, 979, 366, 299, 675, 1456}
	sum := 2020
	expected1, expected2, expected3 := 366, 675, 979

	result1, result2, result3 := FindThreeExpensesThatSumTo(sum, expenses)
	if result1 != expected1 || result2 != expected2 || result3 != expected3 {
		t.Errorf("FindTwoExpensesThatSumTo was incorrect. Receieved: %d, %d %d. Wanted: %d, %d, %d.", result1, result2, result3, expected1, expected2, expected3)
	}
}

func TestFindProductOfThreeExpensesThatSumTo(t *testing.T) {
	expenses := []int{1721, 979, 366, 299, 675, 1456}
	sum := 2020
	expected := 241861950

	result := FindProductOfTwoExpensesThatSumTo(sum, expenses)
	if result != expected {
		t.Errorf("FindProductOfTwoExpensesThatSumTo was incorrect. Received: %d. Wanted: %d.", result, expected)
	}
}
