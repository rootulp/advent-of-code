package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const YOUR_TICKET = "your ticket:"
const NEARBY_TICKETS = "nearby tickets:"

func main() {
	fmt.Println("Starting day16")

	// Part one
	partOne := TicketScanningErrorRate("example.txt")
	fmt.Printf("Part one: %v\n", partOne)
}

func TicketScanningErrorRate(filename string) (errorRate int) {
	lines := readFile(filename)
	rules, myTicket, nearbyTickets := split(lines)

	return 0
}

// split lines into rules, myTicket, and nearbyTickets
func split(lines []string) (rules []string, myTicket string, nearbyTickets []string) {
	seenYourTicket := false
	seenNearbyTickets := false

	for _, line := range lines {
		if line == YOUR_TICKET {
			seenYourTicket = true
		}
		if line == NEARBY_TICKETS {
			seenNearbyTickets = true
		}
		if line == "" {
			continue
		}
		if seenNearbyTickets && seenYourTicket {
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
