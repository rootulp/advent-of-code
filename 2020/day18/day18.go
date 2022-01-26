package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operatorPrecedence map[string]int = map[string]int{
	"+": 1,
	"*": 1,
}

func main() {
	fmt.Println("Starting day18")

	// Part One
	partOne, err := PartOne("input.txt")
	if err != nil {
		fmt.Printf("Part one encountered error: %v", err)
	}
	fmt.Printf("Part one: %v\n", partOne)
}

func PartOne(filename string) (sum int, err error) {
	expressions, err := readLines(filename)
	if err != nil {
		return sum, err
	}
	for _, expression := range expressions {
		sum += Evaluate(expression)
	}
	return sum, nil
}

// ReversePolishNotation converts expression to Reverse Polish Notation using
// the Shunting-yard algorithm
// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func ReversePolishNotation(expression string) (result string) {
	output := []string{}
	operatorStack := Stack{[]string{}} // operatorStack includes parenthesis
	stripped := strings.ReplaceAll(expression, " ", "")
	tokens := strings.Split(stripped, "")
	for _, token := range tokens {
		if isNumber(token) {
			output = append(output, token)
		}
		if isOperator(token) {
			for operatorStack.Len() > 0 && isLeftParen(operatorStack.Peek()) && operatorPrecedence[operatorStack.Peek()] >= operatorPrecedence[token] {
				output = append(output, operatorStack.Pop())
			}
			operatorStack.Push(token)
		}
		if isLeftParen(token) {
			operatorStack.Push(token)
		}
		if isRightParen(token) {
			for !isLeftParen(operatorStack.Peek()) {
				if operatorStack.Len() == 0 {
					panic("operatorStack empty but expected to encounter a left parenthesis\n")
				}
				output = append(output, operatorStack.Pop())
			}
			// Discard the left parenthesis at the top of the stack
			top := operatorStack.Pop()
			if !isLeftParen(top) {
				panic(fmt.Sprintf("expected %v to be left parenthesis", top))
			}
		}
	}
	for operatorStack.Len() != 0 {
		top := operatorStack.Pop()
		if isLeftParen(top) {
			panic(fmt.Sprintf("expected %v to not be left parenthesis", top))
		}
		output = append(output, top)
	}
	return strings.Join(output, " ")
}

func Evaluate(expression string) (result int) {
	// TODO
	return 0
}

func readLines(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isOperator(s string) bool {
	return s == "+" || s == "*"
}

func isLeftParen(s string) bool {
	return s == "("
}

func isRightParen(s string) bool {
	return s == ")"
}

type Stack struct {
	slice []string
}

func (s *Stack) Push(element string) {
	s.slice = append(s.slice, element)
}
func (s *Stack) Pop() (popped string) {
	popped = s.Peek()
	s.slice = s.slice[:len(s.slice)-1]
	return popped
}
func (s *Stack) Peek() (top string) {
	return s.slice[len(s.slice)-1]
}
func (s *Stack) Len() int {
	return len(s.slice)
}
