package main

import (
	"reflect"
	"testing"
)

// Part one
func TestGetNumberOfContainersWithInputTest(t *testing.T) {
	testGetNumberOfContainers(t, "input_test.txt", "shiny gold", 4)
}

func TestGetNumberOfContainersWithInput(t *testing.T) {
	testGetNumberOfContainers(t, "input.txt", "shiny gold", 144)
}

// Part two
func TestGetNumberOfContainedBagsWithInputTest(t *testing.T) {
	testGetNumberOfContainedBags(t, "input_test.txt", "shiny gold", 32)
}
func TestGetNumberOfContainedBagsWithInputTestTwo(t *testing.T) {
	testGetNumberOfContainedBags(t, "input_test2.txt", "shiny gold", 126)
}
func TestGetNumberOfContainedBagsWith(t *testing.T) {
	testGetNumberOfContainedBags(t, "input.txt", "shiny gold", 5956)
}

func TestGetContainersOfWithOneContainer(t *testing.T) {
	colorsToContainers := map[string][]string{
		"bright white": {"light red"},
		"muted yellow": {"light red"},
		"light red":    {},
	}

	testGetContainersOf(t, colorsToContainers, "bright white", []string{"light red"})
}
func TestGetContainersOfWithTwoContainers(t *testing.T) {
	colorsToContainers := map[string][]string{
		"bright white": {"light red", "dark orange"},
		"muted yellow": {"light red", "dark orange"},
		"light red":    {},
		"dark orange":  {},
	}

	testGetContainersOf(t, colorsToContainers, "bright white", []string{"light red", "dark orange"})
}

func TestGetColorsToContainer_WithOneRule(t *testing.T) {
	rules := []string{"light red bags contain 1 bright white bag, 2 muted yellow bags."}
	expected := map[string][]string{
		"bright white": {"light red"},
		"muted yellow": {"light red"},
		"light red":    {},
	}

	testGetColorsToContainer(t, rules, expected)
}

func TestGetColorsToContainer_WithTwoRules(t *testing.T) {
	rules := []string{"light red bags contain 1 bright white bag, 2 muted yellow bags.", "dark orange bags contain 3 bright white bags, 4 muted yellow bags."}
	expected := map[string][]string{
		"bright white": {"light red", "dark orange"},
		"muted yellow": {"light red", "dark orange"},
		"light red":    {},
		"dark orange":  {},
	}

	testGetColorsToContainer(t, rules, expected)
}

func TestParseRuleWithNoContained(t *testing.T) {
	rule := "dotted black bags contain no other bags."
	expectedColor := "dotted black"
	expectedContained := []string{}

	testParseRule(t, rule, expectedColor, expectedContained)
}
func TestParseRuleWithOneContained(t *testing.T) {
	rule := "bright white bags contain 1 shiny gold bag."
	expectedColor := "bright white"
	expectedContained := []string{"shiny gold"}

	testParseRule(t, rule, expectedColor, expectedContained)
}
func TestParseRuleWithMultipleContained1(t *testing.T) {
	rule := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	expectedColor := "light red"
	expectedContained := []string{"bright white", "muted yellow"}

	testParseRule(t, rule, expectedColor, expectedContained)
}
func TestParseRuleWithMultipleContained2(t *testing.T) {
	rule := "dark orange bags contain 3 bright white bags, 4 muted yellow bags."
	expectedColor := "dark orange"
	expectedContained := []string{"bright white", "muted yellow"}

	testParseRule(t, rule, expectedColor, expectedContained)
}

func testParseRule(t *testing.T, rule string, expectedColor string, expectedContained []string) {
	color, contained := ParseRule(rule)

	if color != expectedColor {
		t.Errorf("ParseRule incorrect color. Received %v wanted %v", color, expectedColor)
	}

	if !reflect.DeepEqual(contained, expectedContained) {
		t.Errorf("ParseRule incorrect contained. Received %#v wanted %#v", contained, expectedContained)
	}
}

// Helpers
func testGetNumberOfContainers(t *testing.T, filename string, target string, expected int) {
	result := GetNumberOfContainers(filename, target)

	if result != expected {
		t.Errorf("GetNumberOfContainers failed. Received %v wanted %v", result, expected)
	}
}
func testGetNumberOfContainedBags(t *testing.T, filename string, target string, expected int) {
	result := GetNumberOfContainedBags(filename, target)

	if result != expected {
		t.Errorf("GetNumberOfContainedBags did not match expected. Received %#v expected %#v", result, expected)
	}
}
func testGetColorsToContainer(t *testing.T, rules []string, expected map[string][]string) {
	result := GetColorsToContainers(rules)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetColorsToContainers did not match expected. Received %#v expected %#v", result, expected)
	}
}
func testGetContainersOf(t *testing.T, colorsToContainers map[string][]string, target string, expected []string) {
	result := GetContainersOf(colorsToContainers, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetContainersOf did not match expected. Received %#v expected %#v", result, expected)
	}
}
