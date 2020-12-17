package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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
	fields := strings.Split(line, " ")
	keys := make(map[string]bool)
	for _, field := range fields {
		parts := strings.Split(field, ":")
		key := parts[0]
		keys[key] = true
	}
	delete(keys, "cid") // "cid" is optional
	isValid = reflect.DeepEqual(keys, getRequiredPassportFields())
	fmt.Printf("Line %s, isValid %v\n", line, isValid)
	return
}

// A valid passport has the following required fields:
// - byr (Birth Year)
// - iyr (Issue Year)
// - eyr (Expiration Year)
// - hgt (Height)
// - hcl (Hair Color)
// - ecl (Eye Color)
// - pid (Passport ID)
// - [optional] cid (Country ID)
func getRequiredPassportFields() map[string]bool {
	return map[string]bool{"byr": true, "iyr": true, "eyr": true, "hgt": true, "hcl": true, "ecl": true, "pid": true}
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
