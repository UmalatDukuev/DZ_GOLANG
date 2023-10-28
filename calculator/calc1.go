package main

import (
	"fmt"
	"strconv"
)

func calculate(expression string) (int, error) {
	precedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	tokens := tokenize(expression)
	operators := []rune{}
	numbers := []int{}

	for _, token := range tokens {
		if isNumber(token) {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("Invalid number: %s", token)
			}
			numbers = append(numbers, num)
		} else {
			operator := []rune(token)[0]
			for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[operator] {
				result, err := applyOperation(numbers, operators)
				if err != nil {
					return 0, err
				}
				numbers = append(numbers, result)
			}
			operators = append(operators, operator)
		}
	}

	for len(operators) > 0 {
		result, err := applyOperation(numbers, operators)
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, result)
	}

	if len(numbers) != 1 || len(operators) != 0 {
		return 0, fmt.Errorf("Invalid expression")
	}

	return numbers[0], nil
}

func tokenize(expression string) []string {
	tokens := []string{}
	currentToken := ""

	for _, char := range expression {
		if isOperator(char) || isParenthesis(char) || isWhitespace(char) {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}

			if !isWhitespace(char) {
				tokens = append(tokens, string(char))
			}
		} else {
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func isOperator(char rune) bool {
	operators := map[rune]bool{
		'+': true,
		'-': true,
		'*': true,
		'/': true,
	}
	return operators[char]
}

func isParenthesis(char rune) bool {
	parenthesis := map[rune]bool{
		'(': true,
		')': true,
	}
	return parenthesis[char]
}

func isWhitespace(char rune) bool {
	whitespace := map[rune]bool{
		' ':  true,
		'\t': true,
		'\n': true,
		'\r': true,
	}
	return whitespace[char]
}

func applyOperation(numbers []int, operators []rune) (int, error) {
	if len(numbers) < 2 || len(operators) < 1 {
		return 0, fmt.Errorf("Invalid expression")
	}

	b := numbers[len(numbers)-1]
	numbers = numbers[:len(numbers)-1]

	a := numbers[len(numbers)-1]
	numbers = numbers[:len(numbers)-1]

	operator := operators[len(operators)-1]
	operators = operators[:len(operators)-1]

	switch operator {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, fmt.Errorf("Divide by zero error")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Invalid operator: %c", operator)
	}
}

func main() {
	expression := "4 + 2 * 3 - 6 / 2"
	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
