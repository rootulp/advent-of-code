package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

type Direction string

const (
	North   = "North"
	East    = "East"
	South   = "South"
	West    = "West"
	Left    = "Left"
	Right   = "Right"
	Forward = "Forward"
)

var runeToDirection = map[rune]Direction{
	'N': North,
	'E': East,
	'S': South,
	'W': West,
	'L': Left,
	'R': Right,
	'F': Forward,
}

func main() {
	fmt.Printf("Starting day 12")

	GetManhattanDistance("input_test.txt")
}

func GetManhattanDistance(filename string) (distance int) {
	lines := readFile(filename)
	instructions := parseInstructions(lines)
	log.Print(instructions)

	return 0
}

type Instruction struct {
	command  Direction
	distance int
}

func parseInstructions(lines []string) (instructions []Instruction) {
	for _, line := range lines {
		instruction := parseInstruction(line)
		instructions = append(instructions, instruction)
	}
	return instructions
}

func parseInstruction(line string) (instruction Instruction) {
	command, i := utf8.DecodeRuneInString(line)
	distance, err := strconv.Atoi(line[i:])
	if err != nil {
		log.Fatal(err)
	}
	return Instruction{
		command:  runeToDirection[command],
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
