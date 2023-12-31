package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"
	"uniq"
)

func CheckInput() ([]string, *bufio.Writer) {
	if len(flag.Args()) > 2 {
		fmt.Println("Maximum number of arguments = 2! ")
		os.Exit(0)
	}
	input := os.Stdin
	output := os.Stdout
	var err error
	if flag.Arg(0) != "" {
		input, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(0)
		}
		defer input.Close()
	}
	if flag.Arg(1) != "" {
		output, err = os.Create(flag.Arg(1))
		if err != nil {
			fmt.Println("Error finding output file:", err)
			os.Exit(0)
		}
		defer output.Close()
	}
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, writer
}

func ParseWriter(result []string, writer *bufio.Writer) {
	var err error
	for _, val := range result {
		_, err = writer.WriteString(val)
		if err != nil {
			break
		}
	}
}

func main() {
	strs := []string{
		"111\n",
		"222\n",
		"111\n",
	}
	var opts uniq.Options
	opts = uniq.ParseFlags(opts)
	lines, _ := CheckInput()
	result := uniq.CollapseLines(lines, opts)
	// for _, value := range result {
	// 	fmt.Print(value)
	// }
	// fmt.Println(result)
	// fmt.Println(strs)
	fmt.Println(reflect.DeepEqual(strs, result))
	// 	ParseWriter(result, writer)
	// 	writer.Flush()
}
