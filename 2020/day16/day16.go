package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const YOUR_TICKET = "your ticket:"
const NEARBY_TICKETS = "nearby tickets:"

type Rule struct {
	category string
	rangeA   string
	rangeB   string
}

func main() {
	fmt.Println("Starting day16")

	// Part one
	partOne := TicketScanningErrorRate("example.txt")
	fmt.Printf("Part one: %v\n", partOne)
}

func TicketScanningErrorRate(filename string) (errorRate int) {
	lines := readFile(filename)
	rules, _, nearbyTickets := split(lines)
	validNumbers := getValidNumbers(rules)
	for _, ticket := range nearbyTickets {
		errorRate += getErrorRate(ticket, validNumbers)
	}
	// fmt.Printf("validNumbers %v\n", validNumbers)
	// fmt.Printf("nearbyTickets %v\n", nearbyTickets)
	return errorRate
}

func getErrorRate(ticket string, validNumbers map[int]bool) (errorRate int) {
	ticketNumbers := getNumbers(ticket)
	for _, num := range ticketNumbers {
		if validNumbers[num] {
			continue
		} else {
			errorRate += num
		}
	}
	return errorRate
}

func getValidNumbers(rules []string) (validNumbers map[int]bool) {
	validNumbers = map[int]bool{}

	for _, rule := range rules {
		parsed := parseRule(rule)
		// fmt.Printf("Parsed %+v\n", parsed)
		startA, endA := parseRange(parsed.rangeA)
		startB, endB := parseRange(parsed.rangeB)
		for i := startA; i <= endA; i++ {
			validNumbers[i] = true
		}
		for i := startB; i <= endB; i++ {
			validNumbers[i] = true
		}
	}
	return validNumbers
}

func getNumbers(ticket string) (numbers []int) {
	strs := strings.Split(ticket, ",")
	for _, s := range strs {
		number, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func parseRule(rule string) (parsed Rule) {
	s := strings.Split(rule, ":")
	r := strings.Split(s[1], "or")
	parsed.category = strings.TrimSpace(s[0])
	parsed.rangeA = strings.TrimSpace(r[0])
	parsed.rangeB = strings.TrimSpace(r[1])
	return parsed
}

func parseRange(r string) (start int, end int) {
	s := strings.Split(r, "-")
	start, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err = strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	return start, end
}

// split lines into rules, myTicket, and nearbyTickets
func split(lines []string) (rules []string, myTicket string, nearbyTickets []string) {
	seenYourTicket := false
	seenNearbyTickets := false

	for _, line := range lines {
		if line == YOUR_TICKET {
			seenYourTicket = true
		} else if line == NEARBY_TICKETS {
			seenNearbyTickets = true
		} else if line == "" {
			continue
		} else if seenNearbyTickets && seenYourTicket {
			nearbyTickets = append(nearbyTickets, line)
		} else if seenYourTicket {
			myTicket = line
		} else {
			rules = append(rules, line)
		}
	}
	return rules, myTicket, nearbyTickets
}

// readFile reads filename and returns a list of lines from that file
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
