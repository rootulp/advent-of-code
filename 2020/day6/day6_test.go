package main

import (
	"reflect"
	"testing"
)

func TestGetNumUnique1(t *testing.T) {
	testGetNumUnique(t, []string{"abc"}, 3)
}
func TestGetNumUnique2(t *testing.T) {
	testGetNumUnique(t, []string{"a", "b", "c"}, 3)
}
func TestGetNumUnique3(t *testing.T) {
	testGetNumUnique(t, []string{"ab", "ac"}, 3)
}
func TestGetNumUnique4(t *testing.T) {
	testGetNumUnique(t, []string{"a", "a", "a", "a"}, 1)
}
func TestGetNumUnique5(t *testing.T) {
	testGetNumUnique(t, []string{"b"}, 1)
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

func testGetNumUnique(t *testing.T, group []string, expected int) {
	result := GetNumUnique(group)

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
