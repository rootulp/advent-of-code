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
	number   int
	contents []string
}

func (r Rule) String() string {
	return fmt.Sprintf("rule number %v, contents: [%v]\n", r.number, strings.Join(r.contents, ","))
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
}

func PartOne(filename string) (result int) {
	lines := readLines(filename)
	fmt.Printf("lines: %v\n", lines)

	rawRules, messages := splitRulesAndMessages(lines)
	fmt.Printf("rawRules %v\n", rawRules)
	fmt.Printf("messages %v\n", messages)

	rules := parseRules(rawRules)
	fmt.Printf("rules: %v\n", rules)

	pattern := generateRegex(rules, 0)
	fmt.Printf("pattern: %v\n", pattern)

	result = numMatchingMessages(pattern, messages)
	fmt.Printf("result: %v\n", result)

	return result
}

func numMatchingMessages(pattern string, messages []string) (result int) {
	var matching []string
	for _, message := range messages {
		matched, err := regexp.MatchString(pattern, message)
		if err != nil {
			log.Fatal(err)
		}
		if matched {
			result += 1
			matching = append(matching, message)
		}
	}
	fmt.Printf("matching: %v\n", matching)

	return result
}

func generateRegex(rules map[int]Rule, ruleNumber int) (pattern string) {
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
			subExpression := generateRegex(rules, ruleNum)
			subExpressions = append(subExpressions, subExpression)
		}
		expressions = append(expressions, strings.Join(subExpressions, ""))
	}
	if ruleNumber == 0 {
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

	contents := strings.Split(strings.Trim(parts[1], " "), " | ")
	return Rule{
		number:   number,
		contents: contents,
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
