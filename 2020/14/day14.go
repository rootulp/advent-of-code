package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const mem = "mem"
const mask = "mask"

type instruction struct {
	operation string
	bitmask   string
	address   int
	value     int
}

type memory map[int]int

func main() {
	fmt.Println("Startin day 14")

	// Part one
	SumOfBitmaskedValues("example.txt")
}

func SumOfBitmaskedValues(filename string) (sum int) {
	lines := readLines(filename)
	var memory memory
	var mask string
	for _, line := range lines {
		mask, memory = execute(line, mask, memory)
	}
	return sumOfValues(memory)
}

func execute(line string, mask string, memory memory) (newMask string, newMemory memory) {
	instruction := parse(line)
	fmt.Printf("operation %v address %v value %v bitmask %v\n", instruction.operation, instruction.address, instruction.value, instruction.bitmask)
	return newMask, newMemory
}

func parse(line string) instruction {
	regex := regexp.MustCompile(`(?P<command>\w*)(\[(?P<address>\d*)\])?\s=\s(?P<value>\w*)`)
	match := regex.FindStringSubmatch(line)
	switch operation := match[1]; operation {
	case mem:
		address, _ := strconv.Atoi(match[3])
		value, _ := strconv.Atoi(match[4])
		return instruction{
			operation: operation,
			address:   address,
			value:     value,
		}
	case mask:
		bitmask := match[4]
		return instruction{
			operation: operation,
			bitmask:   bitmask,
		}
	default:
		log.Fatalf("Operation %v is not a valid operation: mem or mask", operation)
		return instruction{}
	}
}

func sumOfValues(state map[int]int) (result int) {
	for _, v := range state {
		result += v
	}
	return result
}

func readLines(filename string) (lines []string) {
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
