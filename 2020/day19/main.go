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

func isString(contents string) bool {
	return strings.HasPrefix(contents, "\"") && strings.HasSuffix(contents, "\"")
}

func generateMessages(rules map[int]Rule, ruleContents string) (messages []string) {
	expressions := strings.Split(ruleContents, "|")
	for _, expression := range expressions {
		var message string
		parts := strings.Split(expression, " ")
		for _, part := range parts {
			ruleNumber, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			message += generateMessage(rules, rules[ruleNumber].contents)
		}
		messages = append(messages, message)
	}
	return messages
}

func generateMessage(rules map[int]Rule, contents string) string {
	if isString(contents) {
		return strings.Trim(contents, "\"")
	}
	return ""
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

	firstRule := rules[0]
	validMessages := generateMessages(rules, firstRule.contents)
	for _, message := range messages {
		if contains(validMessages, message) {
			result += 1
		}
	}

	return result
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

func parseRules(rawRules []string) (rules map[int]Rule) {
	for _, raw := range rawRules {
		rule := parseRule(raw)
		rules[rule.number] = rule
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

func contains(validMessages []string, message string) bool {
	for _, valid := range validMessages {
		if valid == message {
			return true
		}
	}
	return false
}
