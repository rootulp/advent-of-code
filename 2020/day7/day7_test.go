package main

import (
	"reflect"
	"testing"
)

// Part one
func TestGetNumberOfContainersWithInputTest(t *testing.T) {
	testGetNumberOfContainers(t, "input_test.txt", ShinyGold, 4)
}

func TestGetNumberOfContainersWithInput(t *testing.T) {
	testGetNumberOfContainers(t, "input.txt", ShinyGold, 144)
}

// Part two
func TestGetNumberOfContainedBagsWithInputTest(t *testing.T) {
	testGetNumberOfContainedBags(t, "input_test.txt", ShinyGold, 32)
}
func TestGetNumberOfContainedBagsWithInputTestTwo(t *testing.T) {
	testGetNumberOfContainedBags(t, "input_test2.txt", ShinyGold, 126)
}
func TestGetNumberOfContainedBagsWith(t *testing.T) {
	testGetNumberOfContainedBags(t, "input.txt", ShinyGold, 5956)
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
	result := GetColorToContainers(rules)

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
