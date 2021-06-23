package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const mem = "mem"
const mask = "mask"

type instruction struct {
	operation string
	bitmask   string
	address   int
	value     int
}

func main() {
	fmt.Println("Starting day 14")

	// Part one
	SumOfBitmaskedValues("example.txt")
}

func SumOfBitmaskedValues(filename string) (sum int) {
	lines := readLines(filename)
	memory := make(map[int]int)
	var mask string
	for _, line := range lines {
		mask, memory = execute(line, mask, memory)
	}
	return sumOfValues(memory)
}

func execute(line string, mask string, memory map[int]int) (string, map[int]int) {
	instruction := parse(line)
	switch operation := instruction.operation; operation {
	case "mem":
		value := applyMask(mask, instruction.value)
		fmt.Printf("attempting to set address: %v, value: %v\n", instruction.address, value)
		memory[instruction.address] = value
	case "mask":
		mask = instruction.bitmask
	default:
		log.Fatalf("operation %v is not supported\n", operation)
	}
	fmt.Printf("mask: %v, memory %v\n", mask, memory)
	return mask, memory
}

func applyMask(mask string, value int) int {
	fmt.Printf("applying mask: %v, value %v\n", mask, strconv.FormatInt(int64(value), 2))
	orMask := getOrMask(mask)
	andMask := getAndMask(mask)
	value = int(orMask) | value
	value = int(andMask) & value
	return value
}

func getOrMask(mask string) int64 {
	s := strings.Replace(mask, "X", "0", -1)
	result, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getAndMask(mask string) int64 {
	s := strings.Replace(mask, "X", "1", -1)
	result, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func parse(line string) (i instruction) {
	regex := regexp.MustCompile(`(?P<command>\w*)(\[(?P<address>\d*)\])?\s=\s(?P<value>\w*)`)
	match := regex.FindStringSubmatch(line)
	switch operation := match[1]; operation {
	case mem:
		address, _ := strconv.Atoi(match[3])
		value, _ := strconv.Atoi(match[4])
		i = instruction{
			operation: operation,
			address:   address,
			value:     value,
		}
	case mask:
		bitmask := match[4]
		i = instruction{
			operation: operation,
			bitmask:   bitmask,
		}
	default:
		log.Fatalf("operation %v is not supported\n", operation)
		i = instruction{}
	}
	fmt.Printf("operation: %v, address: %v, value: %v, bitmask: %v\n", i.operation, i.address, i.value, i.bitmask)
	return i
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
