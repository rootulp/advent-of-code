package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode/utf8"
)

type Direction string
type Location struct {
	x     float64
	y     float64
	angle int // angle between 0 and 360. East = 90 degrees.
}
type Instruction struct {
	direction Direction
	distance  int
}

const (
	North   = "North"
	East    = "East"
	South   = "South"
	West    = "West"
	Left    = "Left"
	Right   = "Right"
	Forward = "Forward"
)

var runeToDirection = map[rune]Direction{
	'N': North,
	'E': East,
	'S': South,
	'W': West,
	'L': Left,
	'R': Right,
	'F': Forward,
}

func main() {
	fmt.Printf("Starting day 12\n")

	partOne := GetManhattanDistancePartOne("input.txt")
	fmt.Printf("Part one manhattan distance: %v\n", partOne)

	partTwo := GetManhattanDistancePartTwo("input.txt")
	fmt.Printf("Part two manhattan distance: %v\n", partTwo)
}

// Get the manhattan distance of the ship for part one
func GetManhattanDistancePartOne(filename string) int {
	lines := readFile(filename)
	instructions := parseInstructions(lines)
	// log.Print(instructions)

	location := Location{
		x: 0, y: 0, angle: 90,
	}
	// fmt.Printf("Starting location %v\n", location)
	for _, instruction := range instructions {
		location = executeInstruction(instruction, location)
		// fmt.Printf("location %v after %v\n", location, instruction)
	}
	return int(math.Round(manhattanDistance(location.x, location.y)))
}

// Get the manhattan distance of the ship for part two
func GetManhattanDistancePartTwo(filename string) int {
	lines := readFile(filename)
	instructions := parseInstructions(lines)

	location := Location{
		x: 0, y: 0, angle: 90,
	}
	// TODO it is possible to create a new type for waypoint that doesn't include angle
	waypoint := Location{
		x: 10, y: 1,
	}

	// fmt.Printf("Starting location %v waypoint %v\n", location, waypoint)
	for _, instruction := range instructions {
		location, waypoint = executeInstructionWithWaypoint(instruction, location, waypoint)
		// fmt.Printf("Executed instruction %v. New location %v waypoint %v\n", instruction, location, waypoint)
	}
	return int(math.Round(manhattanDistance(location.x, location.y)))
}

func manhattanDistance(x float64, y float64) float64 {
	return math.Abs(x) + math.Abs(y)
}

func executeInstruction(instruction Instruction, location Location) Location {
	switch instruction.direction {
	case North:
		location.y += float64(instruction.distance)
	case East:
		location.x += float64(instruction.distance)
	case South:
		location.y -= float64(instruction.distance)
	case West:
		location.x -= float64(instruction.distance)
	case Right:
		location.angle = location.angle + instruction.distance%360
	case Left:
		diff := location.angle - instruction.distance
		if diff < 0 {
			// diff is negative so we have to wrap around 360
			diff = 360 + diff
		}
		location.angle = diff
	case Forward:
		hypot := instruction.distance
		dy := float64(hypot) * math.Sin(getRadians(90-location.angle))
		dx := float64(hypot) * math.Cos(getRadians(90-location.angle))
		location.x += dx
		location.y += dy
	}

	return location
}

func executeInstructionWithWaypoint(instruction Instruction, location Location, waypoint Location) (Location, Location) {
	switch instruction.direction {
	case North:
		waypoint.y += float64(instruction.distance)
	case East:
		waypoint.x += float64(instruction.distance)
	case South:
		waypoint.y -= float64(instruction.distance)
	case West:
		waypoint.x -= float64(instruction.distance)
	case Right:
		// rotate the waypoint around the ship right (clockwise) the given number of degrees.
		radians := getRadians(instruction.distance)
		waypoint = rotate(location, waypoint, -1.0*radians)
	case Left:
		// rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
		radians := getRadians(instruction.distance)
		waypoint = rotate(location, waypoint, radians)
	case Forward:
		dx := float64(instruction.distance) * (waypoint.x - location.x)
		dy := float64(instruction.distance) * (waypoint.y - location.y)
		location.x += dx
		location.y += dy
		waypoint.x += dx
		waypoint.y += dy
	}

	return location, waypoint
}

func rotate(location Location, waypoint Location, radians float64) Location {
	// Translate waypoint back to origin
	translatedX := waypoint.x - location.x
	translatedY := waypoint.y - location.y

	// Rotate point
	rotatedX := translatedX*math.Cos(radians) - translatedY*math.Sin(radians)
	rotatedY := translatedX*math.Sin(radians) + translatedY*math.Cos(radians)

	// Translate waypoint back relative to ship
	return Location{
		x: rotatedX + location.x,
		y: rotatedY + location.y,
	}

}

// degrees = radians * (180/pi)
// radians = degrees * (pi/180)
func getRadians(degrees int) float64 {
	return float64(degrees) * math.Pi / 180.0
}

func parseInstructions(lines []string) (instructions []Instruction) {
	for _, line := range lines {
		instruction := parseInstruction(line)
		instructions = append(instructions, instruction)
	}
	return instructions
}

func parseInstruction(line string) (instruction Instruction) {
	command, i := utf8.DecodeRuneInString(line)
	distance, err := strconv.Atoi(line[i:])
	if err != nil {
		log.Fatal(err)
	}
	return Instruction{
		direction: runeToDirection[command],
		distance:  distance,
	}
}

func readFile(filename string) (lines []string) {
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
