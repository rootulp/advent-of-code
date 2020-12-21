package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Started day7")
	rules := ReadFile("input_test.txt")
	colorsToContainers := GetColorsToContainers(rules)
	log.Printf("colorsToContainers with values: %#v\n", colorsToContainers)
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
			log.Printf("Ignoring split %v because it does not contain 4 words", words)
		}
	}

	log.Printf("containerColor %#v containedColors %#v\n", color, contained)
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
