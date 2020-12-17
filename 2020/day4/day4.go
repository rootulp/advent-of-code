package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// A valid passport has the following required fields:
// - byr (Birth Year)
// - iyr (Issue Year)
// - eyr (Expiration Year)
// - hgt (Height)
// - hcl (Hair Color)
// - ecl (Eye Color)
// - pid (Passport ID)
// - [optional] cid (Country ID)
type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
}

func main() {
	fmt.Print("Hello world\n")
	lines := readFile("./input.txt")

	validPassports := 0
	for _, line := range lines {
		if IsValidPassport(line) {
			validPassports++
		}
	}
	fmt.Printf("There are %d valid passports", validPassports)
}

// IsValidPassport returns whether the provided line is a valid passport.
func IsValidPassport(line string) (isValid bool) {
	isValid = true
	fmt.Printf("Line %s, isValid %v", line, isValid)
	return
}

func readFile(filename string) (passportLines []string) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)

	passportLines = []string{}
	for fileScanner.Scan() {
		text := fileScanner.Text()
		fmt.Println(text)
		passportLines = append(passportLines, text)
	}
	return
}
