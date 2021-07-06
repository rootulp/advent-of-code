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
	partOne := TicketScanningErrorRate("input.txt")
	fmt.Printf("Part one: %v\n", partOne)

	// Part two
	partTwo := ProductOfDepartureValues("example2.txt")
	fmt.Printf("Part two: %v\n", partTwo)
}

func TicketScanningErrorRate(filename string) (errorRate int) {
	lines := readFile(filename)
	unparsedRules, _, nearbyTickets := split(lines)
	rules := parseRules(unparsedRules)

	validNumbers := getValidNumbers(rules)
	for _, ticket := range nearbyTickets {
		errorRate += getErrorRate(ticket, validNumbers)
	}
	return errorRate
}

func ProductOfDepartureValues(filename string) int {
	lines := readFile(filename)
	unparsedRules, _, nearbyTickets := split(lines)
	rules := parseRules(unparsedRules)

	validNumbers := getValidNumbers(rules)
	validTickets := getValidTickets(nearbyTickets, validNumbers)
	rulePositions := getRulePositions(rules, validTickets)
	fmt.Printf("rulePositions %v", rulePositions)
	return 0
}

func getRulePositions(rules []Rule, validTickets []string) (rulePositions map[Rule]int) {
	rulePositions = map[Rule]int{}
	for _, rule := range rules {
		for position := 0; position < len(rules); position++ {
			if isValidRulePosition(rule, position, validTickets) {
				rulePositions[rule] = position
			}
		}
	}
	return rulePositions
}

func isValidRulePosition(rule Rule, position int, validTickets []string) bool {
	startA, endA := parseRange(rule.rangeA)
	startB, endB := parseRange(rule.rangeB)
	for _, ticket := range validTickets {
		ticketNumbers := getNumbers(ticket)
		num := ticketNumbers[position]
		if num >= startA && num <= endA && num >= startB && num <= endB {
			continue
		} else {
			return false
		}
	}
	return true
}

func getValidTickets(nearbyTickets []string, validNumbers map[int]bool) []string {
	validTickets := []string{}
	for _, ticket := range nearbyTickets {
		if isValidTicket(ticket, validNumbers) {
			validTickets = append(validTickets, ticket)
		}
	}
	fmt.Printf("Valid tickets %v\n", validTickets)
	return validTickets
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

func isValidTicket(ticket string, validNumbers map[int]bool) bool {
	return getErrorRate(ticket, validNumbers) == 0
}

func getValidNumbers(rules []Rule) (validNumbers map[int]bool) {
	validNumbers = map[int]bool{}

	for _, rule := range rules {
		startA, endA := parseRange(rule.rangeA)
		startB, endB := parseRange(rule.rangeB)
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

func parseRules(unparsed []string) (rules []Rule) {
	for _, unp := range unparsed {
		rule := parseRule(unp)
		rules = append(rules, rule)
	}
	return rules
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
