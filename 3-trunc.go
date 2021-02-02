package main

import (
	"fmt"
)

func main() {
	var f float64
	fmt.Println("Enter a float value : ")
	fmt.Scanf("%f", &f)
	fmt.Println(int(f))
}
