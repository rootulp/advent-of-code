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

	validPasswords := 0
	for _, line := range lines {
		if IsValidPasswordPolicyOne(line) {
			validPasswords++
		}
	}
	fmt.Printf("The number of validPasswords for policy one: %v\n", validPasswords)
}

func IsValidPasswordPolicyOne(line string) bool {
	password, character, min, max := ParseLine(line)
	count := strings.Count(password, character)
	return min <= count && count <= max
}

// ParseLine extracts the following fields from a string of the format "1-3 a:
// abcde": password, character, min, and max
func ParseLine(line string) (password string, character string, min int, max int) {
	words := strings.Fields(line)
	min, max = parseCounts(words[0])
	character = strings.Trim(words[1], ":")
	password = words[2]
	return
}

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
