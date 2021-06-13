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

const (
	North   = "North"
	East    = "East"
	South   = "South"
	West    = "West"
	Left    = "Left"
	Right   = "Right"
	Forward = "Forward"
)

var runeToCommand = map[rune]Command{
	'N': North,
	'E': East,
	'S': South,
	'W': West,
	'L': Left,
	'R': Right,
	'F': Forward,
}

type Location struct {
	x float64
	y float64
}
type LocationWithHeading struct {
	Location
	// Angle between 0 and 360.
	// North = 0 degrees. East = 90 degrees.
	angle int
}
type Command string
type Instruction struct {
	command Command
	val     int
}

func main() {
	fmt.Printf("Starting day 12\n")

	partOne := GetManhattanDistancePartOne("input.txt")
	fmt.Printf("Part one manhattan distance: %v\n", partOne)

	partTwo := GetManhattanDistancePartTwo("input.txt")
	fmt.Printf("Part two manhattan distance: %v\n", partTwo)
}

// Get the Manhattan distance of the ship for part one
func GetManhattanDistancePartOne(filename string) int {
	lines := readFile(filename)
	instructions := parseInstructions(lines)
	// log.Print(instructions)

	location := LocationWithHeading{
		Location: Location{x: 0, y: 0}, angle: 90,
	}

	for _, instruction := range instructions {
		location = executeInstruction(instruction, location)
	}
	return int(math.Round(manhattanDistance(location.x, location.y)))
}

// Get the Manhattan distance of the ship for part two
func GetManhattanDistancePartTwo(filename string) int {
	lines := readFile(filename)
	instructions := parseInstructions(lines)

	location := Location{
		x: 0, y: 0,
	}
	waypoint := Location{
		x: 10, y: 1,
	}

	for _, instruction := range instructions {
		location, waypoint = executeInstructionWithWaypoint(instruction, location, waypoint)
		// fmt.Printf("Executed instruction %v. New location %v waypoint %v\n", instruction, location, waypoint)
	}
	return int(math.Round(manhattanDistance(location.x, location.y)))
}

func manhattanDistance(x float64, y float64) float64 {
	return math.Abs(x) + math.Abs(y)
}

func executeInstruction(instruction Instruction, location LocationWithHeading) LocationWithHeading {
	switch instruction.command {
	case North:
		location.y += float64(instruction.val)
	case East:
		location.x += float64(instruction.val)
	case South:
		location.y -= float64(instruction.val)
	case West:
		location.x -= float64(instruction.val)
	case Right:
		location.angle = location.angle + instruction.val%360
	case Left:
		diff := location.angle - instruction.val
		if diff < 0 {
			// diff is negative so we have to wrap around 360
			diff = 360 + diff
		}
		location.angle = diff
	case Forward:
		hypot := instruction.val
		dy := float64(hypot) * math.Sin(getRadians(90-location.angle))
		dx := float64(hypot) * math.Cos(getRadians(90-location.angle))
		location.x += dx
		location.y += dy
	}

	return location
}

func executeInstructionWithWaypoint(instruction Instruction, location Location, waypoint Location) (Location, Location) {
	switch instruction.command {
	case North:
		waypoint.y += float64(instruction.val)
	case East:
		waypoint.x += float64(instruction.val)
	case South:
		waypoint.y -= float64(instruction.val)
	case West:
		waypoint.x -= float64(instruction.val)
	case Right:
		// rotate the waypoint around the ship right (clockwise) the given number of degrees.
		radians := getRadians(instruction.val)
		waypoint = rotate(location, waypoint, -1.0*radians)
	case Left:
		// rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
		radians := getRadians(instruction.val)
		waypoint = rotate(location, waypoint, radians)
	case Forward:
		dx := float64(instruction.val) * (waypoint.x - location.x)
		dy := float64(instruction.val) * (waypoint.y - location.y)
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
	val, err := strconv.Atoi(line[i:])
	if err != nil {
		log.Fatal(err)
	}
	return Instruction{
		command: runeToCommand[command],
		val:     val,
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
