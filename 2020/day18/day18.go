package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting day18")

	// Part One
	partOne := PartOne("input.txt")
	fmt.Printf("PartOne: %v\n", partOne)

	// Part Two
	partTwo := PartTwo("input.txt")
	fmt.Printf("PartTwo: %v\n", partTwo)
}

func PartOne(filename string) (sum int) {
	var operatorPrecedence map[string]int = map[string]int{
		"+": 1,
		"*": 1,
	}
	expressions := readLines(filename)
	for _, expression := range expressions {
		sum += Evaluate(expression, operatorPrecedence)
	}
	return sum
}

func PartTwo(filename string) (sum int) {
	var operatorPrecedence map[string]int = map[string]int{
		"+": 2,
		"*": 1,
	}
	expressions := readLines(filename)
	for _, expression := range expressions {
		sum += Evaluate(expression, operatorPrecedence)
	}
	return sum
}

func Evaluate(expression string, operatorPrecedence map[string]int) (result int) {
	rpn := ReversePolishNotation(expression, operatorPrecedence)
	evaluated := EvaluateReversePolishNotation(rpn)
	return evaluated
}

// ReversePolishNotation converts expression to Reverse Polish Notation using
// the Shunting-yard algorithm
// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func ReversePolishNotation(expression string, operatorPrecedence map[string]int) (result string) {
	output := []string{}
	operatorStack := OperatorStack{[]string{}} // operatorStack includes parenthesis
	stripped := strings.ReplaceAll(expression, " ", "")
	tokens := strings.Split(stripped, "")
	for _, token := range tokens {
		if isNumber(token) {
			output = append(output, token)
		}
		if isOperator(token) {
			for operatorStack.Len() > 0 && !isLeftParen(operatorStack.Peek()) && operatorPrecedence[operatorStack.Peek()] >= operatorPrecedence[token] {
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

func EvaluateReversePolishNotation(rpn string) (result int) {
	operandStack := OperandStack{[]int{}}
	tokens := strings.Split(rpn, " ")
	for _, token := range tokens {
		if isNumber(token) {
			operand, _ := strconv.Atoi(token)
			operandStack.Push(operand)
		} else if isOperator(token) {
			a := operandStack.Pop()
			b := operandStack.Pop()
			operandStack.Push(evaluate(a, b, token))
		} else {
			panic(fmt.Sprintf("token %v is not a number or operator", token))
		}
	}
	return operandStack.Pop()
}

func evaluate(a int, b int, operator string) int {
	if operator == "+" {
		return a + b
	} else if operator == "*" {
		return a * b
	} else {
		panic(fmt.Sprintf("operator %v is not + or *", operator))
	}
}

func readLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
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
