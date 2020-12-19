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
	for _, rule := range rules {
		// fmt.Printf("Rule %v\n", rule)
		ParseRule(rule)
		// fmt.Printf("Container %v contained %v\n", container, contained)
	}
}

type Bag struct {
	Color     string
	Container *Bag
}

// ParseRule gets the containerColor and containedColors from a rule
func ParseRule(rule string) (containerColor string, containedColors []string) {
	fields := strings.Split(rule, "contain")

	containerColor = strings.TrimSpace(strings.TrimSuffix(fields[0], "bags "))
	containedPhrases := strings.Split(strings.Trim(strings.TrimSpace(fields[1]), "."), ", ")
	containedColors = []string{}
	for _, phrase := range containedPhrases {
		words := strings.Fields(phrase)
		if len(words) == 4 {
			// Valid contained expressions contain four words:
			// Ex. "5 faded blue bags"
			containedColor := words[1] + " " + words[2]
			containedColors = append(containedColors, containedColor)
		} else {
			// Invalid contained expressions usually contain three words:
			// Ex. "no other bags"
			log.Printf("Ignoring split %v because it does not contain 4 words", words)
		}
	}

	log.Printf("containerColor %#v containedColors %#v\n", containerColor, containedColors)
	return containerColor, containedColors
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
