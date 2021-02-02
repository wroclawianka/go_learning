package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter your string : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var firstChar = strings.ToLower(input[:1])
	var middle = strings.ToLower(input[1 : len(input)-1])
	var lastChar = strings.ToLower(input[len(input)-1:])

	var isFound = strings.Compare(firstChar, "i") == 0 &&
		strings.Contains(middle, "a") &&
		strings.Compare(lastChar, "n") == 0

	if isFound {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
