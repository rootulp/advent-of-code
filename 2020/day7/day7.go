package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ContainedBag represents a bag that is contained inside another bag.
type ContainedBag struct {
	Quantity int
	Color    string
}

// ShinyGold represents the targetBag in this exercise.
const ShinyGold = "shiny gold"

func main() {
	// Part one
	result1 := GetNumberOfContainers("input.txt", ShinyGold)
	fmt.Printf("%d bag colors can eventually contain at least one shiny gold bag\n", result1)

	// Part two
	result2 := GetNumberOfContainedBags("input.txt", ShinyGold)
	fmt.Printf("A shiny gold bag must contain %d bags\n", result2)
}

// GetNumberOfContainedBags returns the number of total bags that are contained
// inside the target bag.
func GetNumberOfContainedBags(filename string, targetBag string) int {
	rules := ReadFile(filename)
	colorToContainedBags := GetColorToContainedBags(rules)
	// the result expects the number of bags contained in the target bag
	// excluding the target bag therefore subtract one
	return GetNumberOfBagsContainedInside(colorToContainedBags, targetBag) - 1
}

// GetColorToContainedBags returns a map from color name to a slice of objects
// that represent the number and color of the contained bags.
func GetColorToContainedBags(rules []string) map[string][]ContainedBag {
	result := make(map[string][]ContainedBag)

	for _, rule := range rules {
		color, contained := ParseRule(rule)
		result[color] = contained
	}
	return result
}

// GetNumberOfBagsContainedInside returns the total number of bags contained inside the target bag.
func GetNumberOfBagsContainedInside(colorToNumContained map[string][]ContainedBag, targetBag string) int {
	result := 1
	for _, contained := range colorToNumContained[targetBag] {
		result += contained.Quantity * GetNumberOfBagsContainedInside(colorToNumContained, contained.Color)
	}
	return result
}

// GetNumberOfContainers gets the number of containers that can possibly hold the target bag.
func GetNumberOfContainers(filename string, targetBag string) int {
	rules := ReadFile(filename)
	colorsToContainers := GetColorToContainers(rules)
	containers := GetContainersOf(colorsToContainers, targetBag)

	return len(containers)
}

// GetContainersOf retuns a list of strings that represent the bags that
// contain the target bag.
func GetContainersOf(colorsToContainers map[string][]string, targetBag string) []string {
	containers := []string{}
	unvisited := []string{}
	unvisited = append(unvisited, colorsToContainers[targetBag]...)
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

// GetColorToContainers returns a map from bag color to a list of bag colors
// that can contain the color.
func GetColorToContainers(rules []string) map[string][]string {
	colorToContainers := make(map[string][]string)
	// Initialize colors
	for _, rule := range rules {
		color, _ := ParseRule(rule)
		colorToContainers[color] = []string{}
	}

	// Initialize containers
	for _, rule := range rules {
		container, colors := ParseRule(rule)
		for _, color := range colors {
			colorToContainers[color.Color] = append(colorToContainers[color.Color], container)
		}
	}
	return colorToContainers
}

// ParseRule gets the color and the number and color of each
// contained bag from a rule.
func ParseRule(rule string) (color string, contained []ContainedBag) {
	fields := strings.Split(rule, "contain")

	// NOTE: It may make more sense to use a RegEx to capture the relevant
	// fields in this string. However, I wanted to try using the methods exposed
	// by strings as a learning exercise.
	color = strings.TrimSpace(strings.TrimSuffix(fields[0], "bags "))
	containedPhrases := strings.Split(strings.Trim(strings.TrimSpace(fields[1]), "."), ", ")
	contained = []ContainedBag{}
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
			contained = append(contained, ContainedBag{Quantity: containedQuantity, Color: containedColor})
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
