package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	input := scanInput(scanner, "Enter values for acceleration:")
	accelerationValues, _ := strconv.ParseFloat(input, 64)

	input = scanInput(scanner, "Enter initial velocity:")
	initVelocity, _ := strconv.ParseFloat(input, 64)

	input = scanInput(scanner, "Enter initial displacement:")
	initDisplacement, _ := strconv.ParseFloat(input, 64)

	input = scanInput(scanner, "Enter time value:")
	time, _ := strconv.ParseFloat(input, 64)

	fn := GenDisplaceFn(accelerationValues, initVelocity, initDisplacement)
	r := fn(time)
	fmt.Println(r)
}

func scanInput(scanner *bufio.Scanner, command string) string {
	fmt.Println(command)
	scanner.Scan()
	return scanner.Text()
}

func GenDisplaceFn(acceleration float64, initVelocity float64, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		r := 0.5 * acceleration
		r *= math.Pow(time, 2)
		r += initVelocity * time
		r += initialDisplacement
		return r
	}
}
