package main

import (
	"fmt"
	"os"
	"strings"
)

func calculate(expression string) {
	elements := strings.Split(expression, "")
	fmt.Println(elements[0])
	priority := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}
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
