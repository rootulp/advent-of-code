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

func main() {
	lines := readFile("./input.txt")

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
	return passport.IsValid()
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
	passport.Height = properties["hgt"]
	passport.HairColor = properties["hcl"]
	passport.EyeColor = properties["ecl"]
	passport.PassportID = properties["pid"]

	// Parse strings into ints
	byr, err := strconv.Atoi(properties["byr"])
	if err != nil {
		return passport, err
	}
	passport.BirthYear = byr

	iyr, err := strconv.Atoi(properties["iyr"])
	if err != nil {
		return passport, err
	}
	passport.IssueYear = iyr

	eyr, err := strconv.Atoi(properties["eyr"])
	if err != nil {
		return passport, err
	}
	passport.ExpirationYear = eyr

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

// Passport represents a passport. All fields in this struct are required for
// valid passports.
type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
}

// IsValid returns whethere the passport is a valid passport.
func (passport Passport) IsValid() bool {
	return passport.isValidBirthYear() && passport.isValidIssueYear() && passport.isValidExpirationYear() && passport.isValidHeight() && passport.isValidHairColor() && passport.isValidEyeColor() && passport.isValidPassportID()
}

func (passport Passport) isValidBirthYear() bool {
	return passport.BirthYear >= 1920 && passport.BirthYear <= 2002
}

func (passport Passport) isValidIssueYear() bool {
	return passport.IssueYear >= 2010 && passport.IssueYear <= 2020
}

func (passport Passport) isValidExpirationYear() bool {
	return passport.ExpirationYear >= 2020 && passport.ExpirationYear <= 2030
}

func (passport Passport) isValidHeight() bool {
	if len(passport.Height) < 2 {
		fmt.Printf("Height %v is not long enough to contain a unit", passport.Height)
		return false
	}
	unit := string(passport.Height[len(passport.Height)-2:])
	value, err := strconv.Atoi(string(passport.Height[:len(passport.Height)-2]))
	if err != nil {
		fmt.Printf("Failed to parse height %v\n", err)
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

func (passport Passport) isValidHairColor() bool {
	if len(passport.HairColor) < 7 {
		fmt.Printf("hairColor %v is not long enough to be valid\n", passport.HairColor)
		return false
	}

	if !strings.HasPrefix(passport.HairColor, "#") {
		fmt.Printf("hairColor %v does not start with #\n", passport.HairColor)
		return false
	}
	hexadecimalColor := strings.TrimPrefix(passport.HairColor, "#")
	_, err := hex.DecodeString(hexadecimalColor)
	if err != nil {
		fmt.Printf("Failed to decode hexadecimalColor %v\n", hexadecimalColor)
		return false
	}
	return true

}

func (passport Passport) isValidEyeColor() bool {
	for _, color := range getValidEyeColors() {
		if color == passport.EyeColor {
			return true
		}
	}
	fmt.Printf("Invalid eye color %s\n", passport.EyeColor)
	return false
}

func getValidEyeColors() []string {
	return []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
}

func (passport Passport) isValidPassportID() bool {
	length := len(passport.PassportID)
	fmt.Printf("PassportId length %v", length)
	return length == 9
}
