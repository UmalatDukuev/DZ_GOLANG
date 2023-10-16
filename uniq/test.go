package main

import (
	//"fmt"
	"reflect"
	"testing"
	"uniq"
)

func TestUniqDefault(t *testing.T) {
	opts := uniq.Options{
		C: false,
		D: false,
		U: false,
		F: 0,
		S: 0,
		I: false,
	}

	input := []string{
		"111\n",
		"222\n",
		"222\n",
		"111\n",
	}

	expectedRes := []string{
		"111\n",
		"222\n",
		"111\n",
	}
	receivedRes := uniq.CollapseLines(input, opts)

	reflect.DeepEqual(expectedRes, receivedRes)

	// fmt.Println(expectedRes)
	// fmt.Println(receivedRes)
	// fmt.Println(reflect.DeepEqual(expectedRes, receivedRes))
}

func main() {

}
