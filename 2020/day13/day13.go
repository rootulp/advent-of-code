package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Println("Starting day13")

	// Part one
	partOne := GetProductOfEarliestBusAndTimeToWait("input.txt")
	log.Printf("Part one result: %v\n", partOne)

	// Part two
	partTwo := GetEarliestBusWithSubsequentDepartures("input.txt")
	log.Printf("Part two result: %v\n", partTwo)
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

	return getEarliestBusWithSubsequentDepartures(busIds)
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

func getEarliestBusWithSubsequentDepartures(busIds []int) (timestamp int) {
	var a []*big.Int
	var n []*big.Int

	for i, busId := range busIds {
		if busId == 0 {
			continue
		}
		a = append(a, big.NewInt(int64(busId-i)))
		n = append(n, big.NewInt(int64(busId)))
	}

	result, err := chineseRemainderTheorem(a, n)
	if err != nil {
		log.Fatal("Failed to run chinese remainder theorem ", err)
	}

	return int(result.Int64())
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

// Shamelessly copied from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
// See https://brilliant.org/wiki/chinese-remainder-theorem/
var one = big.NewInt(1)

func chineseRemainderTheorem(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
