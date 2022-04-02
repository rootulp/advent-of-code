package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Starting day 19\n")
	partOne := PartOne("input.txt")
	fmt.Printf("Part one: %v\n", partOne)
}

func PartOne(filename string) (result int) {
	lines := readLines(filename)
	fmt.Printf("lines: %v\n", lines)

	return 0
}

func readLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
