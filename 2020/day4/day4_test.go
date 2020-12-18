package main

import "testing"

func TestIsValidPassport_ReturnsTrueForValidPassport1(t *testing.T) {
	validPassport := "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f"
	isValid := IsValidPassport(validPassport)

	if isValid != true {
		t.Errorf("isValid is %v wanted %v", isValid, true)
	}
}

func TestIsValidPassport_ReturnsTrueForValidPassport2(t *testing.T) {
	validPassport := "eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm"
	isValid := IsValidPassport(validPassport)

	if isValid != true {
		t.Errorf("isValid is %v wanted %v", isValid, true)
	}
}

func TestIsValidPassport_ReturnsTrueForValidPassport3(t *testing.T) {
	validPassport := "hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022"
	isValid := IsValidPassport(validPassport)

	if isValid != true {
		t.Errorf("isValid is %v wanted %v", isValid, true)
	}
}

func TestIsValidPassport_ReturnsTrueForValidPassport4(t *testing.T) {
	validPassport := "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
	isValid := IsValidPassport(validPassport)

	if isValid != true {
		t.Errorf("isValid is %v wanted %v", isValid, true)
	}
}

func TestIsValidPassport_ReturnsFalseForInvalidPassport1(t *testing.T) {
	invalidPassport := "eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926"
	isValid := IsValidPassport(invalidPassport)

	if isValid != false {
		t.Errorf("isValid is %v wanted %v", isValid, false)
	}
}

func TestIsValidPassport_ReturnsFalseForInvalidPassport2(t *testing.T) {
	invalidPassport := "iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946"
	isValid := IsValidPassport(invalidPassport)

	if isValid != false {
		t.Errorf("isValid is %v wanted %v", isValid, false)
	}
}

func TestIsValidPassport_ReturnsFalseForInvalidPassport3(t *testing.T) {
	invalidPassport := "hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277"
	isValid := IsValidPassport(invalidPassport)

	if isValid != false {
		t.Errorf("isValid is %v wanted %v", isValid, false)
	}
}

func TestIsValidPassport_ReturnsFalseForInvalidPassport4(t *testing.T) {
	invalidPassport := "hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007"
	isValid := IsValidPassport(invalidPassport)

	if isValid != false {
		t.Errorf("isValid is %v wanted %v", isValid, false)
	}
}
