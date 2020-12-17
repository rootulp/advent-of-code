package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("./input.txt")

	validPolicyOne := 0
	validPolicyTwo := 0
	for _, line := range lines {
		if IsValidPasswordPolicyOne(line) {
			validPolicyOne++
		}
		if IsValidPasswordPolicyTwo(line) {
			validPolicyTwo++
		}
	}
	fmt.Printf("validPolicyOne: %v\n", validPolicyOne)
	fmt.Printf("validPolicyTwo: %v\n", validPolicyTwo)
}

// IsValidPasswordPolicyOne determines if a line is a valid password based on
// the policy: Each line gives the password policy and then the password. The
// password policy indicates the lowest and highest number of times a given
// letter must appear for the password to be valid. For example, 1-3 a means
// that the password must contain a at least 1 time and at most 3 times.
func IsValidPasswordPolicyOne(line string) bool {
	password, character, min, max := ParseLine(line)
	count := strings.Count(password, character)
	return min <= count && count <= max
}

// IsValidPasswordPolicyTwo determines if a line is a valid password based on
// the policy: Each policy actually describes two positions in the password,
// where 1 means the first character, 2 means the second character, and so on.
// (Be careful; Toboggan Corporate Policies have no concept of "index zero"!)
// Exactly one of these positions must contain the given letter. Other
// occurrences of the letter are irrelevant for the purposes of policy
// enforcement. Given the same example list from above:
// 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
// 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
// 2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
func IsValidPasswordPolicyTwo(line string) bool {
	password, character, pos1, pos2 := ParseLine(line)
	// positions are 1 indexed not 0 indexed. Therefore subtract one from all positions
	posOneMatchAndPosTwoNoMatch := character == string(password[pos1-1]) && character != string(password[pos2-1])
	posTwoMatchAndPosOneNoMatch := character != string(password[pos1-1]) && character == string(password[pos2-1])

	return posOneMatchAndPosTwoNoMatch || posTwoMatchAndPosOneNoMatch
}

// ParseLine extracts the following fields from a line of the format "1-3 a:
// abcde": password, character, min, and max
// ParseLine("1-3 a: abcde") => "abcde", "a", 1, 3
func ParseLine(line string) (password string, character string, min int, max int) {
	words := strings.Fields(line)
	min, max = parseCounts(words[0])
	character = strings.Trim(words[1], ":")
	password = words[2]
	return
}

// parseCounts extracts the min and max from a string that represents a range
// parseCounts("1-3") => 1, 3
func parseCounts(word string) (min int, max int) {
	counts := strings.Split(word, "-")
	min, err := strconv.Atoi(counts[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err = strconv.Atoi(counts[1])
	if err != nil {
		log.Fatal(err)
	}
	return
}

func readFile(filename string) (passwords []string) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" { // Reached end of file
			file.Close()
			return passwords
		}

		passwords = append(passwords, text)
	}

	return passwords
}
