package main

import (
	"fmt"
	"strconv"
)

func calculate(expression string) int {
	operators := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}
	var stack []string

	for _, char := range expression {
		s := string(char)

		if isNumber(s) {
			stack = append(stack, s)
		} else if s == "(" {
			stack = append(stack, s)
		} else if s == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, calculateOperation(top, stack[len(stack)-1], stack[len(stack)-2]))
			}

			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			}
		} else {
			for len(stack) > 0 && operators[s] <= operators[stack[len(stack)-1]] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, calculateOperation(top, stack[len(stack)-1], stack[len(stack)-2]))
			}

			stack = append(stack, s)
		}
	}

	for len(stack) > 1 {
		stack = append(stack[:len(stack)-3], calculateOperation(stack[len(stack)-2], stack[len(stack)-3], stack[len(stack)-4]))
	}

	result, _ := strconv.Atoi(stack[0])
	return result
}

func calculateOperation(operator string, operand1 string, operand2 string) string {
	o1, _ := strconv.Atoi(operand1)
	o2, _ := strconv.Atoi(operand2)

	switch operator {
	case "+":
		return strconv.Itoa(o1 + o2)
	case "-":
		return strconv.Itoa(o1 - o2)
	case "*":
		return strconv.Itoa(o1 * o2)
	case "/":
		return strconv.Itoa(o1 / o2)
	}

	return "0"
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	expression1 := "(1+2)-3"
	expression2 := "(1+2)*3"

	fmt.Println(calculate(expression1))
	fmt.Println(calculate(expression2))
}
