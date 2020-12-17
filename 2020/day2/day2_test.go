package main

import "testing"

func TestIsValidPolicyOne_ReturnsTrueForValidPassword(t *testing.T) {
	validPassword := "1-3 a: abcde"
	isValid := IsValidPasswordPolicyOne(validPassword)

	if isValid != true {
		t.Errorf("isValid is %v wanted %v", isValid, true)
	}
}

func TestIsValidPolicyOne_ReturnsFalseForInvalidPassword(t *testing.T) {
	invalidPassword := "1-3 b: cdefg"
	isValid := IsValidPasswordPolicyOne(invalidPassword)

	if isValid != false {
		t.Errorf("isValid is %v wanted %v", isValid, false)
	}
}

func TestParseLine(t *testing.T) {
	input := "1-3 a: abcde"
	expectedPassword := "abcde"
	expectedCharacter := "a"
	expectedMin := 1
	expectedMax := 3

	password, character, min, max := ParseLine(input)

	if password != expectedPassword {
		t.Errorf("ParseLine failed to parse password. Receieved: %v. Wanted: %v", password, expectedPassword)
	}
	if character != expectedCharacter {
		t.Errorf("ParseLine failed to parse character. Receieved: %v. Wanted: %v", character, expectedCharacter)
	}
	if min != expectedMin {
		t.Errorf("ParseLine failed to parse min. Receieved: %v. Wanted: %v", min, expectedMin)
	}
	if max != expectedMax {
		t.Errorf("ParseLine failed to parse max. Receieved: %v. Wanted: %v", max, expectedMax)
	}
}
