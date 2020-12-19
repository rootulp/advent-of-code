package main

import (
	"reflect"
	"testing"
)

func TestGetNumUnique1(t *testing.T) {
	testGetNumUnique(t, "abc", 3)
}
func TestGetNumUnique2(t *testing.T) {
	testGetNumUnique(t, "abac", 3)
}
func TestGetNumUnique3(t *testing.T) {
	testGetNumUnique(t, "aaaa", 1)
}
func TestGetNumUnique4(t *testing.T) {
	testGetNumUnique(t, "b", 1)
}

func TestGetNumCommon1(t *testing.T) {
	testGetNumCommon(t, []string{"abc"}, 3)
}
func TestGetNumCommon2(t *testing.T) {
	testGetNumCommon(t, []string{"a", "b", "c"}, 0)
}
func TestGetNumCommon3(t *testing.T) {
	testGetNumCommon(t, []string{"ab", "ac"}, 1)
}
func TestGetNumCommon4(t *testing.T) {
	testGetNumCommon(t, []string{"a", "a", "a", "a"}, 1)
}
func TestGetNumCommon5(t *testing.T) {
	testGetNumCommon(t, []string{"b"}, 1)
}
func TestReadFileIntoResponses(t *testing.T) {
	expected := []string{
		"abc",
		"abc",
		"abac",
		"aaaa",
		"b"}
	result := ReadFileIntoResponses("./input_test.txt")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ReadFile failed. Received %v expected %v", result, expected)
	}
}
func TestReadFileIntoGroups(t *testing.T) {
	expected := [][]string{
		{"abc"},
		{"a", "b", "c"},
		{"ab", "ac"},
		{"a", "a", "a", "a"},
		{"b"},
	}
	result := ReadFileIntoGroups("./input_test.txt")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ReadFile failed. Received %v expected %v", result, expected)
	}
}

func testGetNumUnique(t *testing.T, response string, expected int) {
	result := GetNumUnique(response)

	if result != expected {
		t.Errorf("GetNumUnique failed. Received %v expected %v", result, expected)
	}
}
func testGetNumCommon(t *testing.T, group []string, expected int) {
	result := GetNumCommon(group)

	if result != expected {
		t.Errorf("GetNumCommon failed. Received %v expected %v", result, expected)
	}
}
