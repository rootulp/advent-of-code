package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting day18")

	// Part One
	partOne, err := PartOne("input.txt")
	if err != nil {
		fmt.Printf("Part one encountered err: %v", err)
	}
	fmt.Printf("Part one: %v\n", partOne)
}

func PartOne(filename string) (sum int, err error) {
	expressions, err := readLines(filename)
	if err != nil {
		return sum, err
	}
	for _, expression := range expressions {
		sum += Evaluate(expression)
	}
	return sum, nil
}

func Evaluate(expression string) (result int) {
	// TODO
	return 0
}

func readLines(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}
