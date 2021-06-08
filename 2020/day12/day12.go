package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Printf("Starting day 12")

	GetManhattanDistance("input_test.txt")
}

func GetManhattanDistance(filename string) (distance int) {
	lines := readFile(filename)
	log.Print(lines)

	return 0
}

type Instruction struct {
	command  rune
	distance int
}

func parseInstruction(line string) (instruction Instruction) {
	command, i := utf8.DecodeRuneInString(line)
	distance, err := strconv.Atoi(line[i:])
	if err != nil {
		log.Fatal(err)
	}
	return Instruction{
		command:  command,
		distance: distance,
	}
}

func readFile(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
