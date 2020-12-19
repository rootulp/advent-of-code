package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	expected := []string{
		"abc",
		"abc",
		"abac",
		"aaaa",
		"b"}
	result := ReadFile("./input_test.txt")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ReadFile failed. Received %v expected %v", result, expected)
	}
}

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

func testGetNumUnique(t *testing.T, response string, expected int) {
	result := GetNumUnique(response)

	if result != expected {
		t.Errorf("GetNumUnique failed. Received %v expected %v", result, expected)
	}
}
