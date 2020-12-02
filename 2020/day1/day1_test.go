package main

import "testing"

func TestFindTwoExpensesThatSumTo(t *testing.T) {
	expenses := []int{1721, 979, 366, 299, 675, 1456}
	sum := 2020
	expected1, expected2 := 299, 1721

	result1, result2 := FindTwoExpensesThatSumTo(sum, expenses)
	if result1 != expected1 || result2 != expected2 {
		t.Errorf("FindTwoExpensesThatSumTo was incorrect, got: %d and %d, want: %d and %d.", result1, result2, expected1, expected2)
	}
}

func TestFindProductOfTwoExpensesThatSumTo(t *testing.T) {
	expenses := []int{1721, 979, 366, 299, 675, 1456}
	sum := 2020
	expected := 514579

	result := FindProductOfTwoExpensesThatSumTo(sum, expenses)
	if result != expected {
		t.Errorf("FindProductOfTwoExpensesThatSumTo was incorrect, got: %d, want: %d.", result, expected)
	}
}
