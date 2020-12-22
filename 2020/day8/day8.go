package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Part one
	log.Println("Starting day8")
	instructions := readFile("input_test.txt")
	value := GetAccumulatorValuePriorToFirstRepeatedInstruction(instructions)
	fmt.Printf("The value in the accumulator is %d prior to the first repeated instruction.", value)
}

// GetAccumulatorValuePriorToFirstRepeatedInstruction returns the value in the
// accumulator prior to the first repeated instruction.
func GetAccumulatorValuePriorToFirstRepeatedInstruction(instructions []string) int {
	return 0
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
