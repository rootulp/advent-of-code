package main

import (
	"reflect"
	"testing"
)

func TestGetSeatID1(t *testing.T) {
	testGetSeatID(t, "FBFBBFFRLR", 357)
}
func TestGetSeatID2(t *testing.T) {
	testGetSeatID(t, "BFFFBBFRRR", 567)
}
func TestGetSeatID3(t *testing.T) {
	testGetSeatID(t, "FFFBBBFRRR", 119)
}
func TestGetSeatID4(t *testing.T) {
	testGetSeatID(t, "BBFFBBFRLL", 820)
}

func TestGetRow1(t *testing.T) {
	testGetRow(t, "FBFBBFF", 44)
}
func TestGetRow2(t *testing.T) {
	testGetRow(t, "BFFFBBF", 70)
}
func TestGetRow3(t *testing.T) {
	testGetRow(t, "FFFBBBF", 14)
}
func TestGetRow4(t *testing.T) {
	testGetRow(t, "BBFFBBF", 102)
}

func TestGetCol1(t *testing.T) {
	testGetCol(t, "RLR", 5)
}
func TestGetCol2(t *testing.T) {
	testGetCol(t, "RRR", 7)
}
func TestGetCol3(t *testing.T) {
	testGetCol(t, "LLL", 0)
}
func TestGetCol4(t *testing.T) {
	testGetCol(t, "RLL", 4)
}

func TestReadFile(t *testing.T) {
	result := ReadFile("./input_test.txt")
	expectedLines := []string{"FFFBBFBLLR",
		"BFBBBFFRLR",
		"BFBBBBFLRR",
		"BBFBFFFLLR",
		"BBFBFBFLLL"}
	if !reflect.DeepEqual(result, expectedLines) {
		t.Errorf("ReadFiles was incorrect. Got %v wanted %v", result, expectedLines)
	}
}

func testGetSeatID(t *testing.T, boardingPass string, expectedSeatID int) {
	result := GetSeatID(boardingPass)
	if result != expectedSeatID {
		t.Errorf("GetSeatID was incorrect. Got %v wanted %v", result, expectedSeatID)
	}
}
func testGetRow(t *testing.T, boardingRow string, expectedRow int) {
	result := GetRow(boardingRow)
	if result != expectedRow {
		t.Errorf("GetRow was incorrect. Got %v wanted %v", result, expectedRow)
	}
}

func testGetCol(t *testing.T, boardingCol string, expectedCol int) {
	result := GetCol(boardingCol)
	if result != expectedCol {
		t.Errorf("GetCol was incorrect. Got %v wanted %v", result, expectedCol)
	}
}
