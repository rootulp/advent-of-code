package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation string

const (
	accumulate  operation = "acc"
	jump        operation = "jmp"
	noOperation operation = "nop"
)

type instruction struct {
	operation operation
	argument  int
}

func main() {
	// Part one
	log.Println("Starting day8")
	lines := readFile("input_test.txt")
	value := GetAccumulatorValuePriorToFirstRepeatedInstruction(lines)
	fmt.Printf("The value in the accumulator is %d prior to the first repeated instruction.", value)
}

// GetAccumulatorValuePriorToFirstRepeatedInstruction returns the value in the
// accumulator prior to the first repeated instruction.
func GetAccumulatorValuePriorToFirstRepeatedInstruction(lines []string) int {
	instructions := parseInstructions(lines)
	executed := make(map[int]bool)
	result, err := ExecuteInstruction(instructions, 0, executed, 0)
	if err != nil {
		log.Fatal("Failed to GetAccumulatorValuePriorToFirstRepeatedInstruction")
	}
	return result
}

// ExecuteInstruction recursively executed the instruction at the provided
// index. It maintains a map of executed instructions and the value present in
// the accumulator.
func ExecuteInstruction(instructions []instruction, index int, executed map[int]bool, accumulator int) (int, error) {
	log.Printf("Executing instructions index %v, executed %v, accumulator %v", index, executed, accumulator)
	if executed[index] == true {
		// We have already executed the current instruction. Therefore return
		// the value present in the accumulator.
		return accumulator, nil
	}

	// Otherwise, mark the current instruction as executed and then execute it.
	executed[index] = true

	switch instructions[index].operation {
	case noOperation:
		return ExecuteInstruction(instructions, index+1, executed, accumulator)
	case accumulate:
		accumulator += instructions[index].argument
		return ExecuteInstruction(instructions, index+1, executed, accumulator)
	case jump:
		return ExecuteInstruction(instructions, index+instructions[index].argument, executed, accumulator)
	default:
		log.Fatalf("instruction %v did not match an expected operation", instructions[index].operation)
	}
	return accumulator, errors.New("Failed to execute instruction")

}

func parseInstructions(lines []string) (instructions []instruction) {
	instructions = []instruction{}
	for _, line := range lines {
		instruction := parseInstruction(line)
		instructions = append(instructions, instruction)
	}
	return instructions
}

func parseInstruction(line string) instruction {
	fields := strings.Fields(line)
	// log.Printf("fields %#v", fields)
	if len(fields) != 2 {
		log.Fatalf("fields %#v is not of length 2", fields)
	}

	var operation operation

	switch fields[0] {
	case string(accumulate):
		operation = accumulate
	case string(jump):
		operation = jump
	case string(noOperation):
		operation = noOperation
	}
	argument, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatalf("failed to convert argument into int %v", fields[1])
	}
	return instruction{operation: operation, argument: argument}
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
