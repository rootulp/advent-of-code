package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid int
}

func (passport Passport) isValid() bool {
	return isValidBirthYear(passport.byr) && isValidIssueYear(passport.iyr) && isValidExpirationYear(passport.eyr) && isValidHeight(passport.hgt) && isValidHairColor(passport.hcl) && isValidEyeColor(passport.ecl) && isValidPassportID(passport.pid)
}

func isValidBirthYear(birthYear int) bool {
	return birthYear >= 1920 && birthYear <= 2002
}

func isValidIssueYear(issueYear int) bool {
	return issueYear >= 2010 && issueYear <= 2020
}

func isValidExpirationYear(expirationYear int) bool {
	return expirationYear >= 2020 && expirationYear <= 2030
}

func isValidHeight(height string) bool {
	if len(height) < 2 {
		// fmt.Printf("Height %v is not long enough to contain a unit", height)
		return false
	}
	unit := string(height[len(height)-2:])
	value, err := strconv.Atoi(string(height[:len(height)-2]))
	if err != nil {
		// fmt.Printf("Failed to parse height %v\n", err)
		return false
	}

	if unit == "cm" {
		return isValidHeightInCentimeters(value)
	} else if unit == "in" {
		return isValidHeightInInches(value)
	} else {
		fmt.Println("Unit is not cm or in")
		return false
	}
}

func isValidHeightInInches(inches int) bool {
	return inches >= 59 && inches <= 76
}

func isValidHeightInCentimeters(centimeters int) bool {
	return centimeters >= 150 && centimeters <= 193
}

func isValidHairColor(hairColor string) bool {
	if len(hairColor) < 7 {
		fmt.Printf("hairColor %v is not long enough to be valid\n", hairColor)
		return false
	}

	if !strings.HasPrefix(hairColor, "#") {
		fmt.Printf("hairColor %v does not start with #\n", hairColor)
		return false
	}
	hexadecimalColor := strings.TrimPrefix(hairColor, "#")
	_, err := hex.DecodeString(hexadecimalColor)
	if err != nil {
		fmt.Printf("Failed to decode hexadecimalColor %v\n", hexadecimalColor)
		return false
	}
	return true

}

func isValidEyeColor(eyeColor string) bool {
	for _, color := range getValidEyeColors() {
		if color == eyeColor {
			return true
		}
	}
	return false
}

func getValidEyeColors() []string {
	return []string{"amb", "blu", "brn", "gry", "hzl", "oth"}
}

func isValidPassportID(passportID int) bool {
	return len(strconv.Itoa(passportID)) == 9
}

func main() {
	lines := readFile("./input_example.txt")

	validPassports := 0
	for _, potentialPassport := range lines {
		if IsValidPassport(potentialPassport) {
			validPassports++
		}
	}
	fmt.Printf("There are %d valid passports\n", validPassports)
}

// IsValidPassport returns whether the provided line is a valid passport.
func IsValidPassport(line string) (isValid bool) {
	passport, err := getPassport(line)
	if err != nil {
		return false
	}
	return passport.isValid()
}

func getProperties(line string) (properties map[string]string) {
	properties = make(map[string]string)
	fields := strings.Split(line, " ")

	for _, property := range fields {
		splitProperty := strings.Split(property, ":")
		key := splitProperty[0]
		val := splitProperty[1]
		properties[key] = val
	}

	return properties
}

// getPassport returns a passport struct for the provided line or returns err if
// the provided line could not be parsed into a passport.
func getPassport(line string) (passport Passport, err error) {
	passport = Passport{}

	properties := getProperties(line)

	// Set all string types of passport
	passport.hgt = properties["hgt"]
	passport.hcl = properties["hcl"]
	passport.ecl = properties["ecl"]

	// Parse strings into ints
	byr, err := strconv.Atoi(properties["byr"])
	if err != nil {
		return passport, err
	}
	passport.byr = byr

	iyr, err := strconv.Atoi(properties["iyr"])
	if err != nil {
		return passport, err
	}
	passport.iyr = iyr

	eyr, err := strconv.Atoi(properties["eyr"])
	if err != nil {
		return passport, err
	}
	passport.eyr = eyr

	pid, err := strconv.Atoi(properties["pid"])
	if err != nil {
		return passport, err
	}
	passport.pid = pid

	fmt.Printf("Passport %#v\n", passport)
	return passport, nil
}

func readFile(filename string) (potentialPassports []string) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)

	potentialPassports = []string{}
	partialPassport := []string{}
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "" { // passports are separated by newlines
			potentialPassports = append(potentialPassports, strings.Join(partialPassport, " "))
			partialPassport = []string{}
		} else {
			partialPassport = append(partialPassport, text)
		}
	}
	potentialPassports = append(potentialPassports, strings.Join(partialPassport, " "))
	// fmt.Printf("passportLines %#v\n", potentialPassports)
	return
}
