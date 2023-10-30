package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stack struct {
	items []string
}

func tokenize(expression string) []string {
	re := regexp.MustCompile(`\d+|\S`) // Регулярное выражение для чисел или одного символа
	return re.FindAllString(expression, -1)
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func isNumber(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}
func isOperator(token string) bool {
	isOperator := false
	if token == "+" || token == "-" || token == "*" || token == "/" || token == "(" || token == ")" {
		isOperator = true
	}
	return isOperator
}

func calculate(expression string) {
	numberStack := Stack{}
	operatorStack := Stack{}

	tokens := tokenize(expression)
	for _, token := range tokens {

		if isNumber(token) == true {
			numberStack.Push(token)
		}

		if isOperator(token) == true {
			operatorStack.Push(token)
		}

	}

	for _, item := range numberStack.items {

	}

	for i := 0; i < len(numberStack.items); i++ {
		fmt.Println(numberStack.items[i])
	}
	for i := 0; i < len(operatorStack.items); i++ {
		fmt.Println(operatorStack.items[i])
	}
	return
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Empty input. Enter the expression!")
		return
	}
	expression := os.Args[1]
	calculate(expression)

}
