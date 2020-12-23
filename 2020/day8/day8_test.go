package main

import "testing"

func TestGetAccumulatorValuePriorToFirstRepeatedInstruction(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		got := GetAccumulatorValuePriorToFirstRepeatedInstruction("input_test.txt")
		expected := 5

		if got != expected {
			t.Errorf("GetAccumulatorValuePriorToFirstRepeatedInstruction got %d expected %d", got, expected)
		}
	})
	t.Run("input", func(t *testing.T) {
		got := GetAccumulatorValuePriorToFirstRepeatedInstruction("input.txt")
		expected := 1684

		if got != expected {
			t.Errorf("GetAccumulatorValuePriorToFirstRepeatedInstruction got %d expected %d", got, expected)
		}
	})
}

func TestGetAccumulatorValueAfterProgramTerminates(t *testing.T) {
	t.Run("input_test", func(t *testing.T) {
		got := GetAccumulatorValueAfterProgramTerminates("input_test.txt")
		expected := 8

		if got != expected {
			t.Errorf("GetAccumulatorValueAfterProgramTerminates got %d expected %d", got, expected)
		}
	})
}
