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
	bags := []Bag{}

	for _, rule := range rules {
		// fmt.Printf("Rule %v\n", rule)
		color, _ := ParseRule(rule)
		bag := Bag{
			Color: color,
			// TODO figure out a way to fetch a pointer to the existing
			// contained bags and add them here.
		}
		bags = append(bags, bag)
	}
}

// Bag represents a bag. Color is the bag's color. Contained is a list of
// pointers to bags that can be contained inside this bag.
type Bag struct {
	Color     string
	Contained []*Bag
}

// ParseRule gets the color and contained from a rule
func ParseRule(rule string) (color string, contained []string) {
	fields := strings.Split(rule, "contain")

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
