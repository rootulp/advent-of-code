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

type Rule struct {
	contents []string
}

func (r Rule) String() string {
	return fmt.Sprintf("rule contents: [%v]\n", strings.Join(r.contents, ","))
}

func (r Rule) isString() bool {
	return len(r.contents) == 1 && strings.HasPrefix(r.contents[0], "\"") && strings.HasSuffix(r.contents[0], "\"")
}

func (r Rule) getString() string {
	return strings.Trim(r.contents[0], "\"")
}

func main() {
	fmt.Printf("Starting day 19\n")

	partOne := PartOne("input.txt")
	fmt.Printf("Part one: %v\n", partOne)

	partTwo := PartTwo("example2.txt")
	fmt.Printf("Part two: %v\n", partTwo)
}

func PartOne(filename string) (result int) {
	lines := readLines(filename)
	rawRules, messages := splitRulesAndMessages(lines)
	rules := parseRules(rawRules)
	pattern := generateRegex(rules, 0, 25)
	return numMatchingMessages(pattern, messages)
}

func PartTwo(filename string) (result int) {
	lines := readLines(filename)
	rawRules, messages := splitRulesAndMessages(lines)

	// Add partTwo specific loop rules
	rawRules = append(rawRules, "8: 42 | 42 8")
	rawRules = append(rawRules, "11: 42 31 | 42 11 31")

	rules := parseRules(rawRules)
	pattern := generateRegex(rules, 0, 25)
	return numMatchingMessages(pattern, messages)
}

func numMatchingMessages(pattern string, messages []string) (result int) {
	for _, message := range messages {
		matched, err := regexp.MatchString(pattern, message)
		if err != nil {
			log.Fatal(err)
		}
		if matched {
			result += 1
		}
	}
	return result
}

func generateRegex(rules map[int]Rule, ruleNumber int, depth int) (pattern string) {
	if depth == 0 {
		return ""
	}
	if rules[ruleNumber].isString() {
		return rules[ruleNumber].getString()
	}
	var expressions []string
	for _, expression := range rules[ruleNumber].contents {
		var subExpressions []string
		for _, rule := range strings.Split(expression, " ") {
			ruleNum, err := strconv.Atoi(string(rule))
			if err != nil {
				log.Fatal(err)
			}
			subExpression := generateRegex(rules, ruleNum, depth-1)
			subExpressions = append(subExpressions, subExpression)
		}
		expressions = append(expressions, strings.Join(subExpressions, ""))
	}
	if ruleNumber == 0 {
		// Anchor expression to start and end of string to ensure entire string matches
		return fmt.Sprintf("^(%v)$", strings.Join(expressions, "|"))
	}
	return fmt.Sprintf("(%v)", strings.Join(expressions, "|"))
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
	rules = map[int]Rule{}
	for _, raw := range rawRules {
		ruleNumber, rule := parseRule(raw)
		rules[ruleNumber] = rule
	}
	return rules
}

func parseRule(rule string) (ruleNumber int, r Rule) {
	parts := strings.Split(rule, ":")
	if len(parts) != 2 {
		log.Fatalf("unexpected parts %v", parts)
	}

	ruleNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(strings.Trim(parts[1], " "), " | ")
	return ruleNumber, Rule{contents}

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
