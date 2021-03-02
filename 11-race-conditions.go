package main

import (
	"fmt"
	"time"
)

// First goroutine is changing value of the i to 1, second is changing to 2, You cannot be sure
// which of the values will be printed

func main() {
	i := 0

	go func(i *int) { *i = 1 }(&i)

	go func(i *int) { *i = 2 }(&i)

	time.Sleep(100 * time.Millisecond) // using times are bad practice, used for demo reasons

	fmt.Println(i)
}
