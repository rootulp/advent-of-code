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

	// Part one
	partOne := GetProductOfEarliestBusAndTimeToWait("input.txt")
	log.Printf("Part one %v\n", partOne)

	// Part two
	partTwo := GetEarliestBusWithSubsequentDepartures("input_test.txt")
	log.Printf("Part two %v\n", partTwo)
}

func GetProductOfEarliestBusAndTimeToWait(filename string) int {
	lines := readLines(filename)

	earliestTimestamp := getEarliestTimestamp(lines[0])
	busIds := getBusIds(lines[1])
	log.Printf("earliestTimestamp %v busIds %v\n", earliestTimestamp, busIds)

	busId, timeToWait := getEarliestBus(earliestTimestamp, busIds)
	log.Printf("busId %v timeToWait %v\n", busId, timeToWait)
	return busId * timeToWait
}

func GetEarliestBusWithSubsequentDepartures(filename string) int {
	lines := readLines(filename)

	busIds := getBusIdsWithX(lines[1])
	log.Printf("busIds %v\n", busIds)

	return 0
}

func getEarliestBus(earliestTimestamp int, busIds []int) (busId int, timeToWait int) {
	currentTime := earliestTimestamp
	for {
		for _, busId := range busIds {
			// log.Printf("currentTime %v busId %v", currentTime, busId)
			if currentTime%busId == 0 {
				timeToWait = currentTime - earliestTimestamp
				return busId, timeToWait
			}
		}
		currentTime += 1
	}
}

func getEarliestTimestamp(input string) int {
	earliestTimestamp, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Failed to convert earliest timestamp to int ", err)
	}
	return earliestTimestamp
}

func getBusIds(input string) []int {
	strings := strings.Split(input, ",")
	busIds := []int{}

	for _, s := range strings {
		busId, err := strconv.Atoi(s)
		// Filter out "X" and any other string that is not an int
		if err == nil {
			busIds = append(busIds, busId)
		}
	}
	return busIds
}

func getBusIdsWithX(input string) []int {
	strings := strings.Split(input, ",")
	busIds := []int{}

	for _, s := range strings {
		busId, _ := strconv.Atoi(s)
		// Include "X" in input as 0 in busIds
		busIds = append(busIds, busId)
	}
	return busIds
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
