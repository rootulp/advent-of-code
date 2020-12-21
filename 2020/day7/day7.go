package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Part one
	result := GetNumberOfContainers("input.txt", "shiny gold")
	fmt.Printf("%d bag colors can eventually contain at least one shiny gold bag\n", result)
}

// GetNumberOfContainers gets the number of containers that can possibly hold the target bag.
func GetNumberOfContainers(filename string, target string) int {
	rules := ReadFile(filename)
	colorsToContainers := GetColorsToContainers(rules)
	containers := GetContainersOf(colorsToContainers, target)

	return len(containers)
}

// GetContainersOf retuns a list of strings that represent that bags that
// contain the target
func GetContainersOf(colorsToContainers map[string][]string, target string) []string {
	containers := []string{}
	unvisited := []string{}
	unvisited = append(unvisited, colorsToContainers[target]...)
	var toVisit string

	for (len(unvisited)) > 0 {
		toVisit, unvisited = unvisited[0], unvisited[1:]
		for _, c := range colorsToContainers[toVisit] {
			if !contains(unvisited, c) && !contains(containers, c) {
				unvisited = append(unvisited, c)
			}
		}
		containers = append(containers, toVisit)
	}

	return containers
}

// GetColorsToContainers returns a map from bag color to a list of bag colors
// that can contain the color.
func GetColorsToContainers(rules []string) map[string][]string {
	colorsToContainers := make(map[string][]string)
	// Initialize colors
	for _, rule := range rules {
		color, _ := ParseRule(rule)
		colorsToContainers[color] = []string{}
	}

	// Initialize containers
	for _, rule := range rules {
		container, colors := ParseRule(rule)
		for _, color := range colors {
			colorsToContainers[color] = append(colorsToContainers[color], container)
		}
	}
	return colorsToContainers
}

// ParseRule gets the color and contained from a rule
func ParseRule(rule string) (color string, contained []string) {
	fields := strings.Split(rule, "contain")

	// NOTE: It may make more sense to use a RegEx to capture the relevant
	// fields in this string. However, I wanted to try using the methods exposed
	// by strings as a learning exercise.
	color = strings.TrimSpace(strings.TrimSuffix(fields[0], "bags "))
	containedPhrases := strings.Split(strings.Trim(strings.TrimSpace(fields[1]), "."), ", ")
	contained = []string{}
	for _, phrase := range containedPhrases {
		words := strings.Fields(phrase)
		if len(words) == 4 {
			// Valid contained expressions contain four words:
			// Ex. "5 faded blue bags"
			containedColor := words[1] + " " + words[2]
			contained = append(contained, containedColor)
		} else {
			// Invalid contained expressions usually contain three words:
			// Ex. "no other bags"
			// log.Printf("Ignoring split %v because it does not contain 4 words", words)
		}
	}

	// log.Printf("containerColor %#v containedColors %#v\n", color, contained)
	return color, contained
}

// ReadFile returns a list of rules from the provided file
func ReadFile(filename string) (rules []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	rules = []string{}
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rules
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
