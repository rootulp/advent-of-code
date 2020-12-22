package main

import "testing"

func TestGetAccumulatorValuePriorToFirstRepeatedInstruction(t *testing.T) {
	instructions := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
	got := GetAccumulatorValuePriorToFirstRepeatedInstruction(instructions)
	expected := 5

	if got != expected {
		t.Errorf("GetAccumulatorValuePriorToFirstRepeatedInstruction got %d expected %d", got, expected)
	}
}
