package main

import (
	"reflect"
	"testing"
)

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
