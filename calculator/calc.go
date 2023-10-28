package main

import (
	"fmt"
	"os"
	"regexp"
)

type numbersStack struct {
	items []int
}

type operationStack struct {
	items []int
}

func tokenize(expression string) []string {
	re := regexp.MustCompile(`\d+|\S`) // Регулярное выражение для чисел или одного символа
	return re.FindAllString(expression, -1)
}

func calculate(expression string) {
	tokens := tokenize(expression)
	for _, token := range tokens {
		//fmt.Println(token)
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
