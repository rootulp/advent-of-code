package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

const mem = "mem"
const mask = "mask"
const bitmask_length = 36

type instruction struct {
	operation string
	bitmask   string
	address   int
	value     int
}

func main() {
	fmt.Println("Starting day 14")

	// Part one
	partOne := PartOne("input.txt")
	fmt.Printf("Part one: %v\n", partOne)

	// Part two
	partTwo := PartTwo("input.txt")
	fmt.Printf("Part two: %v\n", partTwo)
}

func PartOne(filename string) (sum int) {
	lines := readLines(filename)
	memory := make(map[int]int)
	var mask string
	for _, line := range lines {
		mask, memory = executePartOne(line, mask, memory)
	}
	return sumOfValues(memory)
}

func PartTwo(filename string) (sum int) {
	lines := readLines(filename)
	memory := make(map[int]int)
	var mask string
	for _, line := range lines {
		mask, memory = executePartTwo(line, mask, memory)
	}
	return sumOfValues(memory)
}

func executePartOne(line string, mask string, memory map[int]int) (string, map[int]int) {
	instruction := parse(line)
	switch operation := instruction.operation; operation {
	case "mem":
		value := applyMask(mask, instruction.value)
		memory[instruction.address] = value
	case "mask":
		mask = instruction.bitmask
	default:
		log.Fatalf("operation %v is not supported\n", operation)
	}
	return mask, memory
}

func executePartTwo(line string, mask string, memory map[int]int) (string, map[int]int) {
	instruction := parse(line)
	switch operation := instruction.operation; operation {
	case "mem":
		memory = applyMemoryAccessDecoder(memory, mask, instruction.address, instruction.value)
	case "mask":
		mask = instruction.bitmask
	default:
		log.Fatalf("operation %v is not supported\n", operation)
	}
	return mask, memory
}

// applyMask for part one
func applyMask(mask string, value int) int {
	orMask := getOrMask(mask)
	andMask := getAndMask(mask)
	value = int(orMask) | value
	value = int(andMask) & value
	return value
}

// applyMemoryAccessDecoder for part two
func applyMemoryAccessDecoder(memory map[int]int, mask string, address int, value int) map[int]int {
	addressStr := leftPad(strconv.FormatInt(int64(address), 2))
	possibleAddresses := getPossibleAddresses(mask, addressStr, []string{""})
	for _, possible := range possibleAddresses {
		possibleInt64, err := strconv.ParseInt(possible, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		memory[int(possibleInt64)] = value
	}
	return memory
}

func getPossibleAddresses(mask string, address string, possibilities []string) []string {
	if len(mask) == 0 || len(address) == 0 {
		return possibilities
	}
	maskBit, mask := shiftRune(mask)
	addressBit, address := shiftRune(address)
	newPossibilities := []string{}
	if maskBit == '0' {
		// If the bitmask bit is 0, the corresponding memory address bit is unchanged.
		for _, possible := range possibilities {
			newPossibilities = append(newPossibilities, possible+string(addressBit))
		}
	} else if maskBit == '1' {
		// If the bitmask bit is 1, the corresponding memory address bit is overwritten with 1.
		for _, possible := range possibilities {
			newPossibilities = append(newPossibilities, possible+"1")
		}
	} else if maskBit == 'X' {
		// If the bitmask bit is X, the corresponding memory address bit is floating.
		for _, possible := range possibilities {
			newPossibilities = append(newPossibilities, possible+"1")
			newPossibilities = append(newPossibilities, possible+"0")
		}
	}
	return getPossibleAddresses(mask, address, newPossibilities)
}

// getRuneFromString returns the rune at index from str
func getRuneFromString(str string, index int) (r rune) {
	for i, c := range str {
		if i == index {
			return c
		}
	}
	log.Fatalf("Couldn't get rune at index %v from string %v", index, str)
	return
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

// leftPad a string with zeros until it has length bitmask_length
func leftPad(str string) string {
	for len(str) < bitmask_length {
		str = "0" + str
	}
	return str
}

// shiftRune returns the first rune of str and the remainder of str
func shiftRune(str string) (r rune, s string) {
	r, size := utf8.DecodeRuneInString(str)
	if len(str) > size {
		s = str[size:]
	} else {
		s = ""
	}
	return r, s
}

// sumOfValues returns the sum of values in memory
func sumOfValues(memory map[int]int) (result int) {
	for _, v := range memory {
		result += v
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
	return i
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
