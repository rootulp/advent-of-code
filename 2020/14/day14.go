package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type command struct {
	command string
}

type maskCommand struct {
	command
	mask string
}

type writeCommand struct {
	command string
	address int
	value   int
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
	command := parse(line)
	fmt.Printf("command %v", command.command)
	return newMask, newMemory
}

func parse(line string) command {
	r := regexp.MustCompile(`(?P<command>\w*)(\[(?P<address>\d*)\])?\s=\s(?P<value>\w*)`)
	match := r.FindStringSubmatch(line)
	if match[1] == "mask" {
		address, err := strconv.Atoi(match[3])
		if err != nil {
			log.Fatal(err)
		}
		value, err := strconv.Atoi(match[4])
		if err != nil {
			log.Fatal(err)
		}
		result := writeCommand{
			command: match[1],
			address: address,
			value:   value,
		}
		return result
	}
	fmt.Printf(match[0])
	return result
	// _, command, _, address, value :=
	// fmt.Printf("%#v\n", r.FindStringSubmatch(line))
	// fmt.Printf("%#v\n", r.SubexpNames())
	// return command{
	// 	command: command,
	// 	address: address,
	// 	value:   value,
	// }
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
