package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list := make([]int, 0, 10)
	fmt.Println("Print up to 10 integers, Press X to finish")
	for i := 0; i < cap(list); i++ {
		// scan
		input := ""
		fmt.Scan(&input)

		// finish
		if strings.ToLower(input) == "x" {
			break
		}

		// parse int
		var value int
		if v, err := strconv.Atoi(input); err == nil {
			value = v
		}

		list = append(list, value)
		fmt.Println("Current state of list:")
		fmt.Println(list)
	}

	// sort
	BubbleSort(&list)

	// print
	fmt.Println("Sorted list:")
	fmt.Println(list)

}

func BubbleSort(pointer *[]int) {
	for i := len(*pointer); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			Swap(pointer, j)
		}
	}
}

func Swap(pointer *[]int, i int) {
	var list = *pointer
	if list[i] > list[i+1] {
		var temp = list[i]
		list[i] = list[i+1]
		list[i+1] = temp
	}
}
