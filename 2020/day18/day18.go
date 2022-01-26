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
	operatorStack := []string{}
	tokens := strings.Split(expression, " ")
	for _, token := range tokens {
		switch {
		case isNumber(token):
			output = append(output, token)
		case isOperator(token):
			for len(operatorStack) > 0 && isLeftParen(operatorStack[len(operatorStack)-1]) && operatorPrecedence[operatorStack[len(operatorStack)-1]] >= operatorPrecedence[token] {
				output = append(output, token)
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		case isLeftParen(token):
			operatorStack = append(operatorStack, token)
		case isRightParen(token):
			for top := operatorStack[len(operatorStack)-1]; !isLeftParen(top); {
				if len(operatorStack) == 0 {
					panic("operatorStack empty but expected more tokens")
				}
				output = append(output, top)
			}
			// Discard the left parenthesis at the top of the stack
			leftParenthesis := operatorStack[len(operatorStack)-1]
			operatorStack = operatorStack[:len(operatorStack)-1]
			if !isLeftParen(leftParenthesis) {
				panic(fmt.Sprintf("expected %v to be left parenthesis", leftParenthesis))
			}
		}
	}
	for len(operatorStack) != 0 {
		top := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]
		if isLeftParen(top) {
			panic(fmt.Sprintf("expected %v to not be left parenthesis", top))
		}
		output = append(output, top)
	}
	fmt.Printf("operatorStack %v\n", operatorStack)
	fmt.Printf("output %v\n", output)
	return strings.Join(output, " ")
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isOperator(s string) bool {
	return s == "+" || s == "*" || s == "(" || s == ")"
}

func isLeftParen(s string) bool {
	return s == "("
}

func isRightParen(s string) bool {
	return s == ")"
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
