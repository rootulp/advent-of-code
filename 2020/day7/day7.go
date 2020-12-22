package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// NumberContained represents the
type NumberContained struct {
	Quantity int
	Color    string
}

func main() {
	// Part one
	result1 := GetNumberOfContainers("input.txt", "shiny gold")
	fmt.Printf("%d bag colors can eventually contain at least one shiny gold bag\n", result1)

	// Part two
	result2 := GetNumberOfContainedBags("input.txt", "shiny gold")
	fmt.Printf("A shiny gold bag must contain %d bags\n", result2)

}

// GetNumberOfContainedBags returns the number of total bags that are contained
// inside the target bag.
func GetNumberOfContainedBags(filename string, target string) int {
	rules := ReadFile(filename)
	colorToNumContained := GetColorToNumContained(rules)
	return GetNumContained(colorToNumContained, target) - 1 // the result expects the number of bags contained in the target bag excluding the target bag therefore subtract one
}

// GetColorToNumContained returns a map from color name to a slice of objects
// that represent the number and color of the contained bags.
func GetColorToNumContained(rules []string) map[string][]NumberContained {
	result := make(map[string][]NumberContained)

	for _, rule := range rules {
		color, contained := ParseQuantityAndColorFromRule(rule)
		result[color] = contained
	}
	return result
}

// GetNumContained returns the total number of bags contained inside the target bag.
func GetNumContained(colorToNumContained map[string][]NumberContained, target string) int {
	result := 1
	if len(colorToNumContained[target]) == 0 {
		fmt.Printf("A %v bag does not contain any more bags\n", target)
		return result
	}
	for _, contained := range colorToNumContained[target] {
		result += contained.Quantity * GetNumContained(colorToNumContained, contained.Color)
	}
	fmt.Printf("A %v contains %d total bags\n", target, result)
	return result
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

// ParseQuantityAndColorFromRule gets the color and the number and color of each
// contained bag from a rule.
func ParseQuantityAndColorFromRule(rule string) (color string, contained []NumberContained) {
	fields := strings.Split(rule, "contain")

	// NOTE: It may make more sense to use a RegEx to capture the relevant
	// fields in this string. However, I wanted to try using the methods exposed
	// by strings as a learning exercise.
	color = strings.TrimSpace(strings.TrimSuffix(fields[0], "bags "))
	containedPhrases := strings.Split(strings.Trim(strings.TrimSpace(fields[1]), "."), ", ")
	contained = []NumberContained{}
	for _, phrase := range containedPhrases {
		words := strings.Fields(phrase)
		if len(words) == 4 {
			// Valid contained expressions contain four words:
			// Ex. "5 faded blue bags"
			containedQuantity, err := strconv.Atoi(words[0])
			if err != nil {
				log.Fatal(err)
			}
			containedColor := words[1] + " " + words[2]
			contained = append(contained, NumberContained{Quantity: containedQuantity, Color: containedColor})
		}
	}

	log.Printf("containerColor %#v containedColors %#v\n", color, contained)
	return color, contained
}

// ParseRule gets the color and color of each contained bag from a rule.
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
