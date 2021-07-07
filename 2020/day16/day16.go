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
const DEPARTURE = "departure"

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
	partTwo := ProductOfDepartureValues("input.txt")
	fmt.Printf("Part two: %v\n", partTwo)
}

func TicketScanningErrorRate(filename string) (errorRate int) {
	lines := readFile(filename)
	rules, _, nearbyTickets := split(lines)

	validNumbers := getValidNumbers(rules)
	for _, ticket := range nearbyTickets {
		errorRate += getErrorRate(ticket, validNumbers)
	}
	return errorRate
}

func ProductOfDepartureValues(filename string) int {
	lines := readFile(filename)
	rules, myTicket, nearbyTickets := split(lines)

	validNumbers := getValidNumbers(rules)
	validTickets := getValidTickets(nearbyTickets, validNumbers)
	rulePositions := getRulePositions(rules, validTickets)
	departureRules := getDepartureRules(rules)
	departureRulePositions := getDepartureRulePositions(rulePositions, departureRules)
	values := getValuesAtPosition(myTicket, departureRulePositions)
	product := getProduct(values)
	return product
}

func getProduct(arr []int) (product int) {
	product = 1
	for _, v := range arr {
		product *= v
	}
	return product
}

func getValuesAtPosition(ticket []int, departureRulePositions []int) (values []int) {
	for _, position := range departureRulePositions {
		values = append(values, ticket[position])
	}
	return values
}

func getDepartureRulePositions(rulePositions map[Rule]int, departureRules []Rule) (positions []int) {
	for _, rule := range departureRules {
		positions = append(positions, rulePositions[rule])
	}
	return positions
}

func getDepartureRules(rules []Rule) (departureRules []Rule) {
	for _, rule := range rules {
		if strings.HasPrefix(rule.category, DEPARTURE) {
			departureRules = append(departureRules, rule)
		}
	}
	return departureRules
}

func getRulePositions(rules []Rule, validTickets [][]int) (rulePositions map[Rule]int) {
	possibleRulePositions := map[int][]Rule{}
	for _, rule := range rules {
		for position := 0; position < len(rules); position++ {
			if isValidRulePosition(rule, position, validTickets) {
				possibleRulePositions[position] = append(possibleRulePositions[position], rule)
			}
		}
	}
	return prunePossibleRulePositions(possibleRulePositions, len(rules))
}

func prunePossibleRulePositions(possibleRulePositions map[int][]Rule, numRules int) map[Rule]int {
	rulePositions := map[Rule]int{}
	for len(rulePositions) < numRules {
		for position, v := range possibleRulePositions {
			if len(v) == 1 {
				rule := v[0]
				rulePositions[rule] = position
				delete(possibleRulePositions, position)
				possibleRulePositions = removeRule(possibleRulePositions, rule)
			}
		}

	}
	return rulePositions
}

func removeRule(possibleRulePositions map[int][]Rule, rule Rule) map[int][]Rule {
	result := map[int][]Rule{}
	for k, v := range possibleRulePositions {
		filtered := []Rule{}
		for _, r := range v {
			if r != rule {
				filtered = append(filtered, r)
			}
		}
		result[k] = filtered
	}
	return result
}

func isValidRulePosition(rule Rule, position int, validTickets [][]int) bool {
	startA, endA := parseRange(rule.rangeA)
	startB, endB := parseRange(rule.rangeB)
	for _, ticket := range validTickets {
		num := ticket[position]
		if (num >= startA && num <= endA) ||
			(num >= startB && num <= endB) {
			continue
		} else {
			return false
		}
	}
	return true
}

func getValidTickets(nearbyTickets [][]int, validNumbers map[int]bool) [][]int {
	validTickets := [][]int{}
	for _, ticket := range nearbyTickets {
		if isValidTicket(ticket, validNumbers) {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func getErrorRate(ticket []int, validNumbers map[int]bool) (errorRate int) {
	for _, num := range ticket {
		if validNumbers[num] {
			continue
		} else {
			errorRate += num
		}
	}
	return errorRate
}

func isValidTicket(ticket []int, validNumbers map[int]bool) bool {
	for _, num := range ticket {
		if validNumbers[num] {
			continue
		} else {
			return false
		}
	}
	return true
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

func parseTickets(unparsed []string) (tickets [][]int) {
	for _, unp := range unparsed {
		ticket := parseTicket(unp)
		tickets = append(tickets, ticket)
	}
	return tickets
}

func parseTicket(unparsed string) (ticket []int) {
	strs := strings.Split(unparsed, ",")
	for _, s := range strs {
		number, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ticket = append(ticket, number)
	}
	return ticket
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
func split(lines []string) (rules []Rule, myTicket []int, nearbyTickets [][]int) {
	unparsedNearbyTickets := []string{}
	unparsedRules := []string{}
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
			unparsedNearbyTickets = append(unparsedNearbyTickets, line)
		} else if seenYourTicket {
			myTicket = parseTicket(line)
		} else {
			unparsedRules = append(unparsedRules, line)
		}
	}

	rules = parseRules(unparsedRules)
	nearbyTickets = parseTickets(unparsedNearbyTickets)
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
