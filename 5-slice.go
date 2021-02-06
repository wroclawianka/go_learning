package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var slice = make([]int64, 3)
	var i = 0

	for {
		input := ""
		fmt.Println("Enter int value : ")
		fmt.Scan(&input)

		if shouldQuit(input) {
			break
		}

		var value int64
		if v, err := strconv.ParseInt(input, 10, 64); err == nil {
			value = v
		}

		if i < 3 {
			index := indexOf(0, slice)
			if index == -1 {
				fmt.Println("Something went wrong")
			}
			slice[index] = value
		} else {
			slice = append(slice, value)
		}

		sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })

		fmt.Println(slice)
		i++
	}

}

func shouldQuit(input string) bool {
	return strings.ToLower(input) == "x"
}

func indexOf(element int64, data []int64) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
