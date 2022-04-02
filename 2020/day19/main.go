package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	number   int
	contents string
}

func (r Rule) String() string {
	return fmt.Sprintf("rule number %v, contents %v\n", r.number, r.contents)
}

func main() {
	fmt.Printf("Starting day 19\n")

	partOne := PartOne("example.txt")
	fmt.Printf("Part one: %v\n", partOne)
}

func PartOne(filename string) (result int) {
	lines := readLines(filename)
	fmt.Printf("lines: %v\n", lines)

	rawRules, messages := splitRulesAndMessages(lines)
	fmt.Printf("rawRules %v\n", rawRules)
	fmt.Printf("messages %v\n", messages)

	rules := parseRules(rawRules)
	fmt.Printf("rules: %v\n", rules)

	return 0
}

func splitRulesAndMessages(lines []string) (rawRules []string, messages []string) {
	for _, line := range lines {
		if strings.Contains(line, ":") {
			rawRules = append(rawRules, line)
		} else if len(line) != 0 { // filters out empty line divider between rules and messages
			messages = append(messages, line)
		}
	}
	return rawRules, messages
}

func parseRules(rawRules []string) (rules []Rule) {
	for _, raw := range rawRules {
		rule := parseRule(raw)
		rules = append(rules, rule)
	}
	return rules
}

func parseRule(rule string) Rule {
	parts := strings.Split(rule, ":")
	if len(parts) != 2 {
		log.Fatalf("unexpected parts %v", parts)
	}

	number, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	return Rule{
		number:   number,
		contents: parts[1],
	}
}

func readLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
