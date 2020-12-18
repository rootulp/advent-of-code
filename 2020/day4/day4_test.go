package main

import "testing"

func TestIsValidPassport_ReturnsTrueForValidPassport(t *testing.T) {
	validPassport := "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f"
	isValid := IsValidPassport(validPassport)

	if isValid != true {
		t.Errorf("isValid is %v wanted %v", isValid, true)
	}
}

func TestIsValidPassport_ReturnsFalseForInvalidPassport(t *testing.T) {
	invalidPassport := "eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926"
	isValid := IsValidPassport(invalidPassport)

	if isValid != false {
		t.Errorf("isValid is %v wanted %v", isValid, false)
	}
}
