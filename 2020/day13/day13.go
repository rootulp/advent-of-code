package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Println("Starting day13")

	GetProductOfEarliestBusAndTimeToWait("input_test.txt")
}

func GetProductOfEarliestBusAndTimeToWait(filename string) int {
	lines := readLines(filename)

	earliestTimestamp := getEarliestTimestamp(lines[0])
	busIds := getBusIds(lines[1])
	log.Printf("earliestTimestamp %v busIds %v", earliestTimestamp, busIds)

	getEarliestBus(earliestTimestamp, busIds)
	return 0
}

func getEarliestTimestamp(input string) int {
	earliestTimestamp, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Failed to conver earliest timestamp to int ", err)
	}
	return earliestTimestamp
}

func getBusIds(input string) []int {
	strings := strings.Split(input, ",")
	busIds := make([]int, len(strings))

	for i, s := range strings {
		busIds[i], _ = strconv.Atoi(s)
	}
	return busIds
}

func getEarliestBus(earliestTimestamp int, busIds []int) int {
	return 0
}

func readLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
